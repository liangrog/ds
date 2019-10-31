package sort

import (
	"errors"
	"sync"

	"github.com/liangrog/ds/graph/list"
	"github.com/liangrog/ds/graph/parts"
)

// Kahn sorting algorithm for graph. It uses list store for storing vertices.
// This function will make a deep copy of the graph. So be cautious if your
// graph is big as it will consume double the memory.
//
// It returs three values. The `bool` value is used for if the graph is cyclic.
// The second value is the sorted list of the vertices. The error returns as the
// third value.
func Kahn(gh *parts.Graph) (bool, []*parts.Vertice, error) {
	var sorted []*parts.Vertice

	isCyclic := false

	// Make a copy of graph so we won't change the original.
	g := gh.DeepCopy()
	if g.Type == parts.UNDIRECTED {
		return isCyclic, sorted, errors.New("Kahn sort cannot sort an undirected graph")
	}

	// Get initial parents vertices
	pList, err := GetParentOnlyVertices(g)
	if err != nil {
		return isCyclic, sorted, err
	}

	// Perform Kahn
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

	// Check if graph cyclic
	for v := range g.Vertices.IterChan() {
		if v.Edges.Total() > 0 {
			isCyclic = true
			break
		}
	}

	return isCyclic, sorted, nil
}

// Go through the graph looking for vertices that has no incoming edges.
func GetParentOnlyVertices(g *parts.Graph) (*list.VerticeStore, error) {
	s := list.NewVerticeStore()
	res := make(chan error)
	go func() {
		defer close(res)

		var wg sync.WaitGroup
		for v := range g.Vertices.IterChan() {
			wg.Add(1)
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
