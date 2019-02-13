package graph

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchVertice(t *testing.T) {
	var vs []*Vertice

	for i := 0; i < 10; i++ {
		vs := append(vs, &Vertice{Key: strconv.Itoa(i)})
	}

	toFind := "7"
	idx := SearchVertice(vs, toFind)
	assert.Equal(t, toFind, idx)
}
