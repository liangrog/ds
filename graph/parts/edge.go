package parts

type EdgeType string

const (
	// Vertice is origin
	FROM EdgeType = "from"

	// Vertice is reciever
	TO EdgeType = "to"

	// No direction
	NODIR EdgeType = "nodir"
)

// Graph edge
type Edge struct {
	// Edge value
	Value *Value

	// The vertice that the
	// current associated vertice
	// connects to
	Neighbor *Vertice

	// Same as graph type
	Type EdgeType
}

// Edge constructor
func NewEdge(value interface{}, v *Vertice, typ EdgeType) *Edge {
	return &Edge{
		Value:    NewValue(value),
		Neighbor: v,
		Type:     typ,
	}
}

// If given edge has the same
// information as current edge.
// NOTE: It ignores the edge value.
func (e *Edge) Equal(double *Edge) bool {
	return e.Neighbor.Value.Id() == double.Neighbor.Value.Id() &&
		e.Type == double.Type
}
