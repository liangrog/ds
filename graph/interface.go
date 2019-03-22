package graph

// Graph interface
type Graphiface interface {
	// Add object to graph
	AddVertice(obj interface{}) error

	// Update an object in graph
	Update(obj interface{}) error

	// Replace an object with a new one
	Replace(oldKey, objNew interface{}) error

	// Delete an object in graph
	Delete(obj interface{}) error

	// Queries
	Queryiface
}

// Store interface
type Storeiface interface {
	// Add objct to store
	Add(obj interface{}, options map[string]interface{}) error

	// Update an object for store
	Update(obj interface{}, options map[string]interface{}) error

	// Delete an object for store
	Delete(obj interface{}, options map[string]interface{}) error

	// Replace an object for store
	Replace(objSrc, objDest interface{}, options map[string]interface{}) error

	// Queries
	Queryiface

	ChannelIteratoriface
}

// Channel iterator, provide iteration
// for the store contents
type ChannelIteratoriface interface {
	Iter() chan interface{}
}

// Query interface
type Queryiface interface {
	// Query by difference fitlers
	Query(filters ...interface{}) error
}

// Indexer interface
type Indexiface interface {
	// Find a object in store
	Find(hay Storeiface, niddle interface{}) (interface{}, error)
}
