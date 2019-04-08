package list

import (
	"errors"
	"fmt"
	"sync"

	"github.com/liangrog/ds/graph/parts"
)

var _ parts.EdgeStore = &EdgeStore{}

// Edge store for vertice
type EdgeStore struct {
	lock sync.RWMutex

	// Items stored in the list
	Items []*parts.Edge

	Indexer *EdgeIndexer
}

// Vertice store constructor
func NewEdgeStore() *EdgeStore {
	return &EdgeStore{
		Indexer: NewEdgeIndexer(),
	}
}

// Override
func (l *EdgeStore) IterChan() chan *parts.Edge {
	ch := make(chan *parts.Edge)

	// Thread safe
	all := make([]*parts.Edge, len(l.Items))

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

// ChannelIterator implementation
func (l *EdgeStore) Total() int {
	return len(l.Items)
}

func (l *EdgeStore) Empty() {
	var newList []*parts.Edge
	l.Items = newList
}

func (l *EdgeStore) Pop() *parts.Edge {
	x := l.Items[len(l.Items)-1]
	l.Items = l.Items[:len(l.Items)-1]
	return x
}

func (l *EdgeStore) String() string {
	var strOut string
	for _, e := range l.Items {
		strOut += fmt.Sprintf("\t-- Edge Id: %s, Type: %s, Neightbor: %s\n", e.Value.Id(), e.Type, e.Neighbor.Value.Id())
	}

	return strOut
}

// Add object to list
func (l *EdgeStore) Add(obj interface{}, options ...map[string]interface{}) error {
	edge, ok := obj.(*parts.Edge)
	if !ok {
		return errors.New("Invalid edge object given")
	}

	l.lock.Lock()
	defer l.lock.Unlock()

	// Only add if object doesn't exist
	if idx := l.Indexer.Find(l, edge); idx == -1 {
		l.Items = append(l.Items, edge)
	}

	return nil
}

// Delete object from list
func (l *EdgeStore) Delete(obj interface{}, options ...map[string]interface{}) error {
	edge, ok := obj.(*parts.Edge)
	if !ok {
		return errors.New("Invalid edge object given")
	}

	l.lock.Lock()
	defer l.lock.Unlock()

	if i := l.Indexer.Find(l, edge); i != -1 {

		copy(l.Items[i:], l.Items[i+1:])
		l.Items[len(l.Items)-1] = nil // make sure garbage collected
		l.Items = l.Items[:len(l.Items)-1]

	}
	return nil
}
