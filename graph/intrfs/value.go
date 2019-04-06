package intrfs

type Value interface {
	Value() interface{}

	Id() string
}
