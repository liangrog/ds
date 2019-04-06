package intrfs

// Indexer interface
type Index interface {
	// Find an object in store
	Find(s Store, obj interface{}) interface{}
}
