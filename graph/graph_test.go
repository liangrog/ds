package graph

import (
	"fmt"
	"github.com/liangrog/ds/graph/list"
	"github.com/liangrog/ds/graph/parts"
	"github.com/liangrog/ds/graph/sort"
	_ "github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateGraph(t *testing.T) {
	g := parts.NewGraph(parts.DIRECTED, list.NewVerticeStore())
	v1 := parts.NewVertice("v1", list.NewEdgeStore())
	v2 := parts.NewVertice("v2", list.NewEdgeStore())

	// First node has no edge
	g.AddVertice(v1)

	e2 := parts.NewEdge("e2", v1, parts.TO)
	v2.AddEdge(e2)
	if err := g.AddVertice(v2); err != nil {
		fmt.Println(err)
	}

	v3 := parts.NewVertice("v3", list.NewEdgeStore())
	e1 := parts.NewEdge("e1", v1, parts.FROM)
	v3.AddEdge(e1)
	if err := g.AddVertice(v3); err != nil {
		fmt.Println(err)
	}

	v4 := parts.NewVertice("v4", list.NewEdgeStore())
	e4 := parts.NewEdge("e4", v3, parts.TO)
	v4.AddEdge(e4)
	if err := g.AddVertice(v4); err != nil {
		fmt.Println(err)
	}

	v5 := parts.NewVertice("v5", list.NewEdgeStore())
	e52 := parts.NewEdge("e52", v2, parts.FROM)
	v5.AddEdge(e52)
	e53 := parts.NewEdge("e53", v3, parts.FROM)
	v5.AddEdge(e53)
	if err := g.AddVertice(v5); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s", g)

	cyclic, res, err := sort.Kahn(g)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(cyclic)

	fmt.Printf("%s\n", res)

	/*
		if err := g.DeleteVertice(v3); err != nil {
			fmt.Println(err)
		}
	*/
	fmt.Printf("%s", g)

}
