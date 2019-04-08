// Graph holds all vertices references
package parts

import (
	"fmt"

	"github.com/liangrog/ds/graph/utils"
)

// Indicate if graph is directed or undirected
type GraphType string

const (
	// Directed: all edges have direction
	DIRECTED GraphType = "directed"

	// Undirected: Edge has no direction.
	// But here we use edge has both directions
	// to simulate it
	UNDIRECTED GraphType = "undirected"
)

// The graph that holds all the vertices references.
// The vertice can have identities, it depends
// How Store indexer is implemented.
type Graph struct {
	// Type of Graph
	Type GraphType

	// Vertices are provided by a type of Store that
	// persists all the vertices
	Vertices VerticeStore

	// Graph options, provide extensibility
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

// Add a vertice to graph
func (g *Graph) AddVertice(v *Vertice) error {
	// Add new vertice to store
	if err := g.Vertices.Add(v, g.Options); err != nil {
		return err
	}

	// Error result for edge add
	edgeRes := make(chan error)

	// Update all assocaited vertices
	go func() {
		// Close result channel
		defer close(edgeRes)

		wg := utils.GetWait(v.Edges.Total())
		for e := range v.Edges.IterChan() {
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

// Delete vertice from graph
// Delete perform the oppsite order to add:
// Delete edge reference first then the vertice.
func (g *Graph) DeleteVertice(v *Vertice) error {
	verRes := make(chan error)

	go func() {
		// Close result channel
		defer close(verRes)

		wg := utils.GetWait(v.Edges.Total())
		for e := range v.Edges.IterChan() {
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

func (g *Graph) String() string {
	var strOut string
	strOut += fmt.Sprintf("Graph Type: %s, Total Vertices(%d)\n\n", g.Type, g.Vertices.Total())
	strOut += fmt.Sprintf("%s\n", g.Vertices)

	return strOut
}

func (g *Graph) DeepCopy() *Graph {
	return &Graph{
		Type:     g.Type,
		Options:  g.Options,
		Vertices: g.Vertices.DeepCopy(),
	}
}
