package graph

type Indexer interface {
	// Returns -1 if not found
	Find(v *Vertice) int

	// If two vertices are the same
	Equal(a, b *Vertice) bool
}
