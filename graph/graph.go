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
	Indexer *Indexer
}

// Constructor.
// Minimum one indexer is requre.
func NewGraph(idxr *Indexer) *Graph {
	return &Graph{Indexer: idxr}
}

// Remove all edges in the Vertices refering to the given vertice
func (g *Graph) RemoveVeticeEdgeReference(v *Vertice) {
	idx := g.Indexer.Find(v)
	if idx == -1 {
		return
	}

	vs := make(chan *Vertice, 10)

	workerNum = len(g.Vertices)
	if workerNum > workerMax {
		workerNum = workerMax
	}

	var wg sync.WaitGroup
	wg.Add(workerNum)

	worker := func(in <-chan *Vertice) {
		for v := range in {
			var edgeCopy []*Edge
			copy(edgeCopy, v.Edges)

			for eIdx := range edgeCopy {
				if from := v.Edges[eIdx].GetFrom(); from != nil && g.Indexer.Equal(v, from) {
					v.Edges = append(v.Edges[:eIdx], a[eIdx+1:]...)
				}

				if to := v.Edges[ecIdx].GetTo(); to != nil && g.Indexer.Equal(v, to) {
					v.Edges = append(v.Edges[:ecIdx], a[ecIdx+1:]...)
				}
			}
		}

		wg.Done()
	}

	// Spawn works
	for i := 0; i < num; i++ {
		go worker(vs)
	}

	for i, v := range g.Vertices {
		// Don't do anything for the given vertice
		if idx == i {
			continue
		}
		vs <- v
	}

	close(vs)

	wg.Wait()
}

// Add Edge to the vertices that reference the givenn vertice
func (g *Graph) AddVeticeEdgeReference(v *Vertice) *Graph {
	for i, e := range v.Edges {
		ed := &Edge{Weight: e.Weight}

		if from := e.GetFrom(); from != nil {
			from.Edges = append(from.Edges, ed.SetTo(v))
		}

		if to := e.GetTo(); to != nil {
			to.Edges = append(to.Edges, ed.SetFrom(v))
		}
	}
	return g
}

// Add or replace the given vertice
func (g *Graph) UpsertVertice(v *Vertice) *Graph {
	idx := g.Indexer.Find(v)
	// If not found
	if idx == -1 {
		g.Vertices = append(g.Vertices, v)
		idx = len(g.Vertices) - 1
	} else {
		g.RemoveVeticeEdgeReference(v)
		// Replace the old vertice with the new one
		g.Vertices[idx] = v
	}

	g.AddVeticeEdgeReference(v)

	return g
}

// Find a vertice using given indexer.
// The indexer must return -1 if no vertice found.
func (g *Graph) GetVertice(v *Vertice) *Vertice {
	idx := g.Indexer.Find(v)
	if idx != -1 {
		return g.Vertices[idx]
	}

	return nil
}

// Delete a vertice from graph
func (g *Graph) DeleteVertice(v *Vertice) *Graph {
	idx := g.Indexer.Find(v)
	// If not found
	if idx == -1 {
		return g
	}

	g.RemoveVeticeEdgeReference(v)

	g.Vertices = append(g.Vertice[idx:], g.Vertice[idx+1:]...)

	return g
}

type Vertice struct {
	Value *interface{}
	Edges []*Edge
}

func (v *Vertice) AddEdge() {
}

func (v *Vertice) DeleteEdge() {
}

func (v *Vertice) UpdateEdge() {
}

type Edge struct {
	Weight int
	From   *Vertice
	To     *Vertice
}

// Set start vertice
func (e *Edge) SetFrom(v *Vertice) *Edge {
	e.From = v
	return e
}

// Get start vertice
func (e *Edge) GetFrom() *Vertice {
	return e.From
}

// Set to vertice
func (e *Edge) SetTo(v *Vertice) *Edge {
	e.To = v
	return e
}

// Get to vertice
func (e *Edge) GetTo() *Vertice {
	return e.To
}
