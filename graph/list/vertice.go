package list

import (
	"errors"
	"fmt"
	"sync"

	"github.com/liangrog/ds/graph/parts"
	"github.com/liangrog/ds/graph/utils"
)

var _ parts.VerticeStore = &VerticeStore{}

// List store for vertice
type VerticeStore struct {
	// Thread safe lock
	lock sync.RWMutex

	// Items stored in the list
	Items []*parts.Vertice

	// Indexer for the store
	Indexer *VerticeIndexer
}

// Vertice store constructor
func NewVerticeStore() *VerticeStore {
	return &VerticeStore{
		Indexer: NewVerticeIndexer(),
	}
}

// Iterator. Returns each vertice through a channel.
func (l *VerticeStore) IterChan() chan *parts.Vertice {
	ch := make(chan *parts.Vertice)

	// Thread safe
	all := make([]*parts.Vertice, len(l.Items))

	l.lock.Lock()
	copy(all, l.Items)
	l.lock.Unlock()

	go func() {
		defer close(ch)
		for _, val := range all {
			ch <- val
		}
	}()

	return ch
}

// Deep copy of the vertice store. It replicates each
// vertice in the store including all edges.
func (l *VerticeStore) DeepCopy() parts.VerticeStore {
	cp := NewVerticeStore()

	// Create new vertice store
	// with duplicated vertices.
	for v := range l.IterChan() {
		vcp := &parts.Vertice{
			Value: v.Value.DeepCopy(),
			Edges: NewEdgeStore(),
		}

		cp.Add(vcp)
	}

	// Replicated all edges.
	wg := utils.GetWait(l.Total())
	for v := range l.IterChan() {
		go func(old *parts.Vertice) {
			// Find match new vertice
			vv := cp.Items[cp.Indexer.Find(cp, old)]

			for e := range old.Edges.IterChan() {
				index := cp.Indexer.Find(cp, e.Neighbor)
				ne := &parts.Edge{
					Value:    e.Value.DeepCopy(),
					Neighbor: cp.Items[index],
					Type:     e.Type,
				}

				vv.Edges.Add(ne)
			}

			wg.Done()
		}(v)
	}

	wg.Wait()

	return cp
}

// Iterator. Returns total number of vertice.
func (l *VerticeStore) Total() int {
	return len(l.Items)
}

// Empty the list.
func (l *VerticeStore) Empty() {
	var newList []*parts.Vertice
	l.Items = newList
}

// Returns the last vertice added.
func (l *VerticeStore) Pop() *parts.Vertice {
	x := l.Items[len(l.Items)-1]
	l.Items = l.Items[:len(l.Items)-1]

	return x
}

// String presentation of the store.
func (l *VerticeStore) String() string {
	var strOut string
	for _, v := range l.Items {
		strOut += fmt.Sprintf("== Vertice Id: %s, Total Edges(%d)\n", v.Value.Id(), v.Edges.Total())
		strOut += fmt.Sprintf("%s\n", v.Edges)
	}

	return strOut
}

// Get vertice by ID.
func (l *VerticeStore) Get(id string) *parts.Vertice {
	if idx := l.Indexer.FindById(l, id); idx != -1 {
		return l.Items[idx]
	}

	return nil
}

// Get vertice by custom function.
func (l *VerticeStore) GetByFunc(f parts.VerticeSearchFunc) *parts.Vertice {
	if idx := l.Indexer.FindByFunc(l, f); idx != -1 {
		return l.Items[idx]
	}

	return nil
}

// Add vertice to the store. It's thread safe.
func (l *VerticeStore) Add(obj interface{}, options ...map[string]interface{}) error {
	vertice, ok := obj.(*parts.Vertice)
	if !ok {
		return errors.New("Invalid vertice object given")
	}

	l.lock.Lock()
	defer l.lock.Unlock()

	// Only add if object doesn't exist
	if idx := l.Indexer.Find(l, vertice); idx == -1 {
		l.Items = append(l.Items, vertice)
	}

	return nil
}

// Delete vertice from the store. It's thread safe.
func (l *VerticeStore) Delete(obj interface{}, options ...map[string]interface{}) error {
	vertice, ok := obj.(*parts.Vertice)
	if !ok {
		return errors.New("Invalid vertice object given")
	}

	l.lock.Lock()
	defer l.lock.Unlock()

	if i := l.Indexer.Find(l, vertice); i != -1 {
		copy(l.Items[i:], l.Items[i+1:])
		l.Items[len(l.Items)-1] = nil // make sure garbage collected
		l.Items = l.Items[:len(l.Items)-1]
	}

	return nil
}

// Indexer for vertice.
type VerticeIndexer struct {
}

// Vertice indexer constructor.
func NewVerticeIndexer() *VerticeIndexer {
	return &VerticeIndexer{}
}

// Find the index of a vertice from store. If not found, returns -1.
func (idxr *VerticeIndexer) Find(l *VerticeStore, v *parts.Vertice) int {
	for idx, val := range l.Items {
		if val.Equal(v) {
			return idx
		}
	}

	return -1
}

// Find item by given ID.
func (idxr *VerticeIndexer) FindById(l *VerticeStore, id string) int {
	for idx, val := range l.Items {
		if val.Value.Id() == id {
			return idx
		}
	}

	return -1
}

// Find item by custom function.
func (idxr *VerticeIndexer) FindByFunc(l *VerticeStore, f parts.VerticeSearchFunc) int {
	for idx, val := range l.Items {
		if f(val) {
			return idx
		}
	}

	return -1
}
