package graph

import (
	"fmt"
	"testing"

	_ "github.com/stretchr/testify/assert"
)

func SimpleGraph() *Graph {
	indexer := StringValueIndexer("name")
	a := new(Vertice).SetValue("a")
	b := new(Vertice).SetValue("b")
	c := new(Vertice).SetValue("c")
	d := new(Vertice).SetValue("d")
	e := new(Vertice).SetValue("e")

	a.AddEdge(EDGE_TO, b, indexer)
	a.AddEdge(EDGE_TO, c, indexer)

	b.AddEdge(EDGE_TO, d, indexer)
	b.AddEdge(EDGE_FROM, a, indexer)

	d.AddEdge(EDGE_FROM, b, indexer)
	e.AddEdge(EDGE_FROM, b, indexer)

	graph := NewGraph(indexer).
		UpsertVertice(a).
		UpsertVertice(b).
		UpsertVertice(c).
		UpsertVertice(d).
		UpsertVertice(e)

	//
	d.AddEdge(EDGE_FROM, c, indexer)
	graph.UpsertVertice(d)
	return graph
}

func SimpleNoDirGraph() *Graph {
	indexer := StringValueIndexer("name")
	a := new(Vertice).SetValue("a")
	b := new(Vertice).SetValue("b")
	c := new(Vertice).SetValue("c")
	d := new(Vertice).SetValue("d")
	e := new(Vertice).SetValue("e")

	a.AddEdge(EDGE_UNDIR, b, indexer)
	a.AddEdge(EDGE_UNDIR, c, indexer)

	b.AddEdge(EDGE_UNDIR, d, indexer)
	b.AddEdge(EDGE_UNDIR, a, indexer)

	d.AddEdge(EDGE_UNDIR, b, indexer)
	e.AddEdge(EDGE_UNDIR, b, indexer)

	graph := NewGraph(indexer).
		UpsertVertice(a).
		UpsertVertice(b).
		UpsertVertice(c).
		UpsertVertice(d).
		UpsertVertice(e)

	//
	d.AddEdge(EDGE_UNDIR, c, indexer)
	graph.UpsertVertice(d)
	return graph
}

func TestSearchVertice(t *testing.T) {
	g := SimpleGraph()
	for i, v := range g.Vertices {
		fmt.Println("== ", i, v.GetValue())
		for _, e := range v.Edges {
			fmt.Println(e.GetDirection(), e.GetAttach().GetValue())
		}
	}

	fmt.Println("+++++++++++++++")

	v, err := Kahn(g)

	if err != nil {
		fmt.Println(err)
	}

	if v != nil {
		for _, vv := range v {
			fmt.Println(vv.GetValue())
		}
	}
	/*
		for i, v := range SimpleNoDirGraph().Vertices {
			fmt.Println("++ ", i, v.GetValue())
			for _, e := range v.Edges {
				fmt.Println(e.GetDirection(), e.GetAttach().GetValue())
			}
		}
	*/
}
