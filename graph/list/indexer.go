package list

import (
	"github.com/liangrog/ds/graph/parts"
)

// Indexer for vertice
type VerticeIndexer struct {
}

// Vertice indexer constructor
func NewVerticeIndexer() *VerticeIndexer {
	return &VerticeIndexer{}
}

// Find a vertice from list by comparing the pointer
func (idxr *VerticeIndexer) Find(l *VerticeStore, v *parts.Vertice) int {
	for idx, val := range l.Items {
		if val.Equal(v) {
			return idx
		}
	}

	return -1
}

// Indexer for edge
type EdgeIndexer struct {
}

// Edge indexer constructor
func NewEdgeIndexer() *EdgeIndexer {
	return &EdgeIndexer{}
}

// Find an edge from list
func (idxr *EdgeIndexer) Find(l *EdgeStore, e *parts.Edge) int {
	for idx, val := range l.Items {
		if val.Equal(e) {
			return idx
		}
	}

	return -1
}
