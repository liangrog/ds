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
	lock sync.RWMutex

	// Items stored in the list
	Items []*parts.Vertice

	Indexer *VerticeIndexer
}

// Vertice store constructor
func NewVerticeStore() *VerticeStore {
	return &VerticeStore{
		Indexer: NewVerticeIndexer(),
	}
}

// Override
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

func (l *VerticeStore) DeepCopy() parts.VerticeStore {
	cp := NewVerticeStore()

	// Create new vertice store
	// with duplicated vertices
	for v := range l.IterChan() {
		vcp := &parts.Vertice{
			Value: v.Value.DeepCopy(),
			Edges: NewEdgeStore(),
		}

		cp.Add(vcp)
	}

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

// ChannelIterator implementation
func (l *VerticeStore) Total() int {
	return len(l.Items)
}

func (l *VerticeStore) Empty() {
	var newList []*parts.Vertice
	l.Items = newList
}

func (l *VerticeStore) Pop() *parts.Vertice {
	x := l.Items[len(l.Items)-1]
	l.Items = l.Items[:len(l.Items)-1]

	return x
}

func (l *VerticeStore) String() string {
	var strOut string
	for _, v := range l.Items {
		strOut += fmt.Sprintf("== Vertice Id: %s, Total Edges(%d)\n", v.Value.Id(), v.Edges.Total())
		strOut += fmt.Sprintf("%s\n", v.Edges)
	}

	return strOut
}

// Add object to list
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

// Delete object from list
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
