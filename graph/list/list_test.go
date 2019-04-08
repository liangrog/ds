package list

import (
	"fmt"
	"testing"

	"github.com/liangrog/ds/graph/parts"
	"github.com/stretchr/testify/assert"
)

func TestCreateGraph(t *testing.T) {
	g := parts.NewGraph(parts.DIRECTED, NewVerticeStore())
	v1 := parts.NewVertice("v1", NewEdgeStore())
	v2 := parts.NewVertice("v2", NewEdgeStore())

	// First node has no edge
	err := g.AddVertice(v1)
	assert.NoError(t, err)

	e2 := parts.NewEdge("e2", v1, parts.TO)
	err = v2.AddEdge(e2)
	assert.NoError(t, err)
	err = g.AddVertice(v2)
	assert.NoError(t, err)

	v3 := parts.NewVertice("v3", NewEdgeStore())
	e1 := parts.NewEdge("e1", v1, parts.FROM)
	err = v3.AddEdge(e1)
	assert.NoError(t, err)
	err = g.AddVertice(v3)
	assert.NoError(t, err)

	v4 := parts.NewVertice("v4", NewEdgeStore())
	e4 := parts.NewEdge("e4", v3, parts.TO)
	err = v4.AddEdge(e4)
	assert.NoError(t, err)
	err = g.AddVertice(v4)
	assert.NoError(t, err)

	v5 := parts.NewVertice("v5", NewEdgeStore())
	e52 := parts.NewEdge("e52", v2, parts.FROM)
	err = v5.AddEdge(e52)
	assert.NoError(t, err)
	e53 := parts.NewEdge("e53", v3, parts.FROM)
	err = v5.AddEdge(e53)
	assert.NoError(t, err)
	err = g.AddVertice(v5)
	assert.NoError(t, err)

	fmt.Printf("%s", g)
}
