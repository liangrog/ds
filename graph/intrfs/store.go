package intrfs

// Underlyinng vertices and edges store interface
type Store interface {
	// Add objct to store
	Add(obj interface{}, options ...map[string]interface{}) error

	// Delete an object for store
	Delete(obj interface{}, options ...map[string]interface{}) error

	// String presentation of the store
	String() string
}
