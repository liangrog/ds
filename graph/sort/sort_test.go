package sort

import (
	"github.com/liangrog/ds/graph/list"
	"github.com/liangrog/ds/graph/parts"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKahn(t *testing.T) {
	g := parts.NewGraph(parts.DIRECTED, list.NewVerticeStore())
	v1 := parts.NewVertice("v1", list.NewEdgeStore())
	v2 := parts.NewVertice("v2", list.NewEdgeStore())

	// First node has no edge
	g.AddVertice(v1)

	e2 := parts.NewEdge("e2", v1, parts.TO)
	v2.AddEdge(e2)
	g.AddVertice(v2)

	v3 := parts.NewVertice("v3", list.NewEdgeStore())
	e1 := parts.NewEdge("e1", v1, parts.FROM)
	v3.AddEdge(e1)
	g.AddVertice(v3)

	v4 := parts.NewVertice("v4", list.NewEdgeStore())
	e4 := parts.NewEdge("e4", v3, parts.TO)
	v4.AddEdge(e4)
	g.AddVertice(v4)

	v5 := parts.NewVertice("v5", list.NewEdgeStore())
	e25 := parts.NewEdge("e25", v2, parts.FROM)
	v5.AddEdge(e25)
	e53 := parts.NewEdge("e53", v3, parts.FROM)
	v5.AddEdge(e53)
	g.AddVertice(v5)

	cyclic, res, err := Kahn(g)

	assert.False(t, cyclic)
	assert.NoError(t, err)
	assert.Equal(t, 5, len(res))

	e52 := parts.NewEdge("e52", v2, parts.TO)
	v5.AddEdge(e52)
	g.UpdateVertice(v5)

	cyclic, _, err = Kahn(g)
	assert.True(t, cyclic)
	assert.NoError(t, err)
}
