package parts

import (
	"fmt"
	"sync"
)

// Graph type can be directed or undirected.
type GraphType string

const (
	// Directed: all edges have direction.
	DIRECTED GraphType = "directed"

	// Undirected: Edge has no direction.
	UNDIRECTED GraphType = "undirected"
)

// The graph that holds all the vertices references.
type Graph struct {
	// Type of Graph.
	Type GraphType

	// Vertices are provided by a type of Store that
	// persists all the vertices.
	Vertices VerticeStore

	// Graph options, provide extensibility and future use.
	Options map[string]interface{}
}

// Graph constructor.
func NewGraph(typ GraphType, store VerticeStore, options ...map[string]interface{}) *Graph {
	var opts map[string]interface{}

	// Merge options
	if len(options) > 0 {
		for _, o := range options {
			for k, v := range o {
				opts[k] = v
			}
		}
	}

	return &Graph{
		Type:     typ,
		Vertices: store,
		Options:  opts,
	}
}

// Add a vertice to graph. This function will trigger an
// update on all neighboring vertices, making sure the new
// edge will be added to all neighboring vertices. The fact
// that both vertices hold the same edge information will
// reduce number of travers/search required.
//
// Caution: This is NOT a commit, so any erros could potentially
// cause edge reference out of sync.
func (g *Graph) AddVertice(v *Vertice) error {
	// Add new vertice to store
	if err := g.Vertices.Add(v, g.Options); err != nil {
		return err
	}

	// Error result for edge add
	edgeRes := make(chan error)

	// Update all neighboring vertices
	go func() {
		// Close result channel
		defer close(edgeRes)

		var wg sync.WaitGroup
		for e := range v.Edges.IterChan() {
			wg.Add(1)
			go func(edge *Edge) {
				// the other related vertice
				var t EdgeType
				switch edge.Type {
				case FROM:
					t = TO
				case TO:
					t = FROM
				case NODIR:
					t = NODIR
				}

				edgeRes <- edge.Neighbor.Edges.Add(NewEdge(edge.Value, v, t))
				wg.Done()
			}(e)
		}

		wg.Wait()
	}()

	// Check edge add result
	for er := range edgeRes {
		if er != nil {
			return er
		}
	}

	return nil
}

// Update a vertice.
func (g *Graph) UpdateVertice(v *Vertice) error {
	return g.AddVertice(v)
}

// Delete a vertice from graph.
// This function is similar to `AddVertice`. The same caution should be noted.
func (g *Graph) DeleteVertice(v *Vertice) error {
	verRes := make(chan error)

	go func() {
		// Close result channel
		defer close(verRes)

		var wg sync.WaitGroup
		for e := range v.Edges.IterChan() {
			wg.Add(1)
			go func(edge *Edge) {
				verRes <- edge.Neighbor.DeleteNeighborRef(v)
				wg.Done()
			}(e)
		}

		wg.Wait()
	}()

	// Check vertice ref delete result
	for vr := range verRes {
		if vr != nil {
			return vr
		}
	}

	if err := g.Vertices.Delete(v, g.Options); err != nil {
		return err
	}

	return nil
}

// String presentation of the graph.
func (g *Graph) String() string {
	var strOut string
	strOut += fmt.Sprintf("Graph Type: %s, Total Vertices(%d)\n\n", g.Type, g.Vertices.Total())
	strOut += fmt.Sprintf("%s\n", g.Vertices)

	return strOut
}

// Perform a deep copy of the graph.
func (g *Graph) DeepCopy() *Graph {
	return &Graph{
		Type:     g.Type,
		Options:  g.Options,
		Vertices: g.Vertices.DeepCopy(),
	}
}

// Find vertice by it's internal ID.
func (g *Graph) GetVerticeById(id string) *Vertice {
	if found := g.Vertices.Get(id); found != nil {
		return found
	}

	return nil
}

// Find vertice by user provided function.
func (g *Graph) GetVerticeByFunc(f VerticeSearchFunc) *Vertice {
	if found := g.Vertices.GetByFunc(f); found != nil {
		return found
	}

	return nil
}
