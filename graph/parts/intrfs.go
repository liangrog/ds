package parts

import (
	"github.com/liangrog/ds/graph/intrfs"
)

type VerticeChanIterator interface {
	IterChan() chan *Vertice
	Total() int
}

type VerticeStore interface {
	Pop() *Vertice
	Empty()
	DeepCopy() VerticeStore

	intrfs.Store
	VerticeChanIterator
}

type EdgeChanIterator interface {
	IterChan() chan *Edge
	Total() int
}

type EdgeStore interface {
	Pop() *Edge
	Empty()

	intrfs.Store
	EdgeChanIterator
}
