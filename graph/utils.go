package graph

import (
	"sort"
)

// Search vertice in a vertice slice
func SearchVertice(vs []*Vertice, key string) *Vertice {
	sort.Sort(vs)

	idx := sort.Search(len(vs), func(i int) bool {
		return vs[i].GetKey() == key
	})

	// If found index
	if idx < len(vs) {
		return vs[idx]
	}

	return nil
}
