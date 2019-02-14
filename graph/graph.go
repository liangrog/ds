package graph

import (
	"sync"
)

const (
	workerMax = 100
)

// The graph that holds all the vertices.
// The vertice can have identities, it depends
// How indexer is implemented.
type Graph struct {
	// All vertices belong to the graph.
	Vertices []*Vertice

	// Graph indices.
	Indexer Indexer
}

// Constructor.
// Minimum one indexer is requre.
func NewGraph(idxr Indexer) *Graph {
	return &Graph{Indexer: idxr}
}

// Remove all edges in the Vertices refering to the given vertice
func (g *Graph) RemoveVeticeEdgeReference(v *Vertice) {
	idx := g.Indexer.Find(g.Vertices, v)
	if idx == -1 {
		return
	}

	vs := make(chan *Vertice, 10)

	workerNum := len(g.Vertices)
	if workerNum > workerMax {
		workerNum = workerMax
	}

	var wg sync.WaitGroup
	wg.Add(workerNum)

	// func that removes the edges have reference to the vertice
	worker := func(in <-chan *Vertice) {
		for v := range in {
			var edgeCopy []*Edge
			copy(edgeCopy, v.Edges)

			for eIdx := range edgeCopy {
				if g.Indexer.Equal(v, v.Edges[eIdx].GetAttach()) {
					v.Edges = append(v.Edges[:eIdx], v.Edges[eIdx+1:]...)
				}
			}
		}

		wg.Done()
	}

	// Spawn works
	for i := 0; i < workerNum; i++ {
		go worker(vs)
	}

	for i, v := range g.Vertices {
		// Don't do anything for the given vertice
		if idx == i {
			continue
		}
		vs <- v
	}

	// Signal job complete
	close(vs)

	wg.Wait()
}

// Add Edge to the vertices that reference the givenn vertice
func (g *Graph) AddVeticeEdgeReference(v *Vertice) *Graph {
	for _, e := range v.Edges {
		ed := new(Edge).
			SetWeight(e.Weight).
			SetAttach(v)

		ed.SetDirection(EDGE_FROM)
		if e.GetDirection() == EDGE_FROM {
			ed.SetDirection(EDGE_TO)
		}

		ref := e.GetAttach()
		exist := false
		for _, ee := range ref.Edges {
			if ee.GetDirection() == e.GetDirection() && g.Indexer.Equal(ee.GetAttach(), e.GetAttach()) {
				exist = true
			}
		}

		if !exist {
			ref.Edges = append(ref.Edges, ed)
		}
	}

	return g
}

// Add or replace the given vertice
func (g *Graph) UpsertVertice(v *Vertice) *Graph {
	idx := g.Indexer.Find(g.Vertices, v)
	// If not found
	if idx == -1 {
		g.Vertices = append(g.Vertices, v)
		idx = len(g.Vertices) - 1
	} else {
		g.RemoveVeticeEdgeReference(v)
		// Replace the old vertice with the new one
		g.Vertices[idx] = v

		g.AddVeticeEdgeReference(v)
	}

	return g
}

// Find a vertice using given indexer.
// The indexer must return -1 if no vertice found.
func (g *Graph) GetVertice(v *Vertice) *Vertice {
	idx := g.Indexer.Find(g.Vertices, v)
	if idx != -1 {
		return g.Vertices[idx]
	}

	return nil
}

// Delete a vertice from graph
func (g *Graph) DeleteVertice(v *Vertice) *Graph {
	idx := g.Indexer.Find(g.Vertices, v)
	// If not found
	if idx == -1 {
		return g
	}

	// Remove vertice referennces
	g.RemoveVeticeEdgeReference(v)

	// Remove vertice
	g.Vertices = append(g.Vertices[idx:], g.Vertices[idx+1:]...)

	return g
}

// Vertice for the graph
type Vertice struct {
	// Value can be anything user defined
	Value interface{}

	// All edges that come from or
	// go out to from this vertice.
	// Note each edge information exists
	// on both from vertice and to vertice
	// so there is no traverse required for
	// looking up for parent or child.
	Edges []*Edge
}

// Set vertice value
func (v *Vertice) SetValue(val interface{}) *Vertice {
	v.Value = val
	return v
}

// Get vertice value
func (v *Vertice) GetValue() interface{} {
	return v.Value
}

// Search if an directional edge already exists, if non-exist, return -1
func (v *Vertice) SearchEdgeDirectional(direction EdgeDirection, whom *Vertice, indexer Indexer) int {
	for i, e := range v.Edges {
		if e.GetDirection() == direction && indexer.Equal(e.GetAttach(), whom) {
			return i
		}
	}

	return -1
}

// Search all edges that have an relationship to the given vertice
func (v *Vertice) SearchEdgeAll(whom *Vertice, indexer Indexer) []int {
	var found []int

	for i, e := range v.Edges {
		if a := e.GetAttach(); indexer.Equal(a, whom) {
			found = append(found, i)
		}
	}

	return found
}

// Add an edge to vertice
func (v *Vertice) AddEdge(direction EdgeDirection, whom *Vertice, indexer Indexer) *Vertice {
	// Don't add the edge if it's already exist
	if v.SearchEdgeDirectional(direction, whom, indexer) != -1 {
		return v
	}

	v.Edges = append(v.Edges, new(Edge).SetAttach(whom).SetDirection(direction))

	return v
}

// Delete an edge from vertice
func (v *Vertice) DeleteEdge(direction EdgeDirection, whom *Vertice, indexer Indexer) *Vertice {
	// Don't delete the edge if it doesn't exist
	i := v.SearchEdgeDirectional(direction, whom, indexer)
	if i != -1 {
		return v
	}

	v.Edges = append(v.Edges[:i], v.Edges[i+1:]...)

	return v
}

// Update an edge's weight
func (v *Vertice) UpdateEdgeWeight(w int, direction EdgeDirection, whom *Vertice, indexer Indexer) *Vertice {
	// Don't do anything if it doesn't exist
	i := v.SearchEdgeDirectional(direction, whom, indexer)
	if i != -1 {
		return v
	}

	v.Edges[i].SetWeight(w)

	return v
}

// Swapping/transfering relationship from one vertice to another.
func (v *Vertice) SwapEdgeVertice(whom, toWhom *Vertice, indexer Indexer) *Vertice {
	found := v.SearchEdgeAll(whom, indexer)
	if len(found) > 0 {
		for _, i := range found {
			v.Edges[i].SetAttach(toWhom)
		}
	}

	return v
}

// Edge direction if it's to or from other vertice
type EdgeDirection string

const (
	EDGE_TO   EdgeDirection = "to"
	EDGE_FROM EdgeDirection = "from"
)

// Graph edge
type Edge struct {
	// Edge weight
	Weight int

	//Vertice attached to
	Attach *Vertice

	// To or from vertice
	Direction EdgeDirection
}

// Set vertice attached to
func (e *Edge) SetAttach(v *Vertice) *Edge {
	e.Attach = v
	return e
}

// Get vertice attached to
func (e *Edge) GetAttach() *Vertice {
	return e.Attach
}

// Set edge direction
func (e *Edge) SetDirection(ed EdgeDirection) *Edge {
	e.Direction = ed
	return e
}

// Get edge direction
func (e *Edge) GetDirection() EdgeDirection {
	return e.Direction
}

// Set weight
func (e *Edge) SetWeight(w int) *Edge {
	e.Weight = w
	return e
}

// Get weight
func (e *Edge) GetWeight() int {
	return e.Weight
}
