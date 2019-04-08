package intrfs

// Value interface for vertice and edge value
type Value interface {
	// Return the value
	Value() interface{}

	// Return the id
	Id() string
}
