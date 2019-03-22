package graph

// Simple indexer for vertice value is a string
type StringValueIndexer string

// Find vertice by string value.
func (si StringValueIndexer) Find(hay []*Vertice, niddle *Vertice) int {
	for i, v := range hay {
		// Ignore non string value
		v1, ok := v.GetValue().(string)
		if !ok {
			continue
		}

		if v1 == niddle.GetValue().(string) {
			return i
		}
	}

	return -1
}

// Compare vertice by string value.
// Note it will only comapre the value, not including the edges.
func (si StringValueIndexer) Equal(a, b *Vertice) bool {

	if a.GetValue().(string) == b.GetValue().(string) {
		return true
	}

	return false
}
