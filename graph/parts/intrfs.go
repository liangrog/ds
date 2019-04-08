package parts

import (
	"github.com/liangrog/ds/graph/intrfs"
)

// Channel iteration for vertice store
type VerticeChanIterator interface {
	// Returns vertice pointer.
	IterChan() chan *Vertice

	// Returns total vertice number.
	Total() int
}

// Vertice store interface.
type VerticeStore interface {
	// Pop the last vertice stored in the store.
	Pop() *Vertice

	// Remove all stored vertices.
	Empty()

	// Deep copy the store.
	DeepCopy() VerticeStore

	// Store interface.
	intrfs.Store

	// Vertice iterator interface.
	VerticeChanIterator
}

// Channel iteration for edge store
type EdgeChanIterator interface {
	// Returns edge pointer.
	IterChan() chan *Edge

	// Returns total edge number.
	Total() int
}

// Edge store interface.
type EdgeStore interface {
	// Pop the last edge stored in the store.
	Pop() *Edge

	// Remove all stored edges.
	Empty()

	// Store interface.
	intrfs.Store

	// Edge iterator interface.
	EdgeChanIterator
}
