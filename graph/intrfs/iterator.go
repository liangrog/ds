package intrfs

// Channel iterator, provide iteration
// for the store contents
type ChannelIterator interface {
	IterChan() chan interface{}

	Total() int
}
