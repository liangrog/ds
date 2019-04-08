package sort

import (
	"errors"

	"github.com/liangrog/ds/graph/list"
	"github.com/liangrog/ds/graph/parts"
	"github.com/liangrog/ds/graph/utils"
)

func Kahn(gh *parts.Graph) (bool, []*parts.Vertice, error) {
	var sorted []*parts.Vertice

	isCyclic := false

	g := gh.DeepCopy()
	if g.Type == parts.UNDIRECTED {
		return isCyclic, sorted, errors.New("Kahn sort cannot sort an undirected graph")
	}

	// Get initial parents vertices
	pList, err := GetParentOnlyVertices(g)
	if err != nil {
		return isCyclic, sorted, err
	}

	for pList.Total() > 0 {
		v := pList.Pop()
		sorted = append(sorted, v)

		for _, c := range v.GetNeighbors() {
			c.DeleteNeighborRef(v)
			if c.IsParentOnly() {
				pList.Add(c)
			}
		}

		v.Edges.Empty()
	}

	for v := range g.Vertices.IterChan() {
		if v.Edges.Total() > 0 {
			isCyclic = true
			break
		}
	}

	return isCyclic, sorted, nil
}

// Go through the graph looking for
// vertices that only have the source
// edge.
func GetParentOnlyVertices(g *parts.Graph) (*list.VerticeStore, error) {
	s := list.NewVerticeStore()
	res := make(chan error)
	go func() {
		defer close(res)
		wg := utils.GetWait(g.Vertices.Total())
		for v := range g.Vertices.IterChan() {
			go func(vertice *parts.Vertice) {
				if vertice.IsParentOnly() {
					err := s.Add(vertice)
					res <- err
				}
				wg.Done()
			}(v)
		}

		wg.Wait()
	}()

	for re := range res {
		if re != nil {
			return s, re
		}
	}

	return s, nil
}
