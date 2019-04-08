package parts

type EdgeType string

const (
	// Vertice is origin.
	FROM EdgeType = "from"

	// Vertice is reciever.
	TO EdgeType = "to"

	// No direction.
	NODIR EdgeType = "nodir"
)

// Graph edge struct.
type Edge struct {
	// Edge value.
	Value *Value

	// The neighboring vertice connected to.
	Neighbor *Vertice

	// Edge type.
	Type EdgeType
}

// Edge constructor.
func NewEdge(value interface{}, v *Vertice, typ EdgeType) *Edge {
	return &Edge{
		Value:    NewValue(value),
		Neighbor: v,
		Type:     typ,
	}
}

// Compare two edges based on value internal ID and edge type.
func (e *Edge) Equal(double *Edge) bool {
	return e.Neighbor.Value.Id() == double.Neighbor.Value.Id() &&
		e.Type == double.Type
}
