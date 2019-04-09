package parts

// Vertice for the graph.
type Vertice struct {
	// Value can be anything user defined by using Value struct.
	Value *Value

	// All edges that come from or
	// go out to from this vertice.
	// Note each edge is referenced by
	// both from vertice and to vertice.
	Edges EdgeStore
}

// Vertice constructor.
func NewVertice(value interface{}, edgeStore EdgeStore) *Vertice {
	return &Vertice{
		Value: NewValue(value),
		Edges: edgeStore,
	}
}

// Delete edge reference to a neighboring vertice.
func (v *Vertice) DeleteNeighborRef(neighbor *Vertice) error {
	for e := range v.Edges.IterChan() {
		if e.Neighbor.Value.Id() == neighbor.Value.Id() {
			if err := v.Edges.Delete(e); err != nil {
				return err
			}
		}
	}

	return nil
}

// Get all neighboring vertices.
func (v *Vertice) GetNeighbors() []*Vertice {
	var res []*Vertice
	for e := range v.Edges.IterChan() {
		res = append(res, e.Neighbor)
	}

	return res
}

// Add an edge to vertice.
func (v *Vertice) AddEdge(e *Edge) error {
	return v.Edges.Add(e)
}

// Update an edge to vertice.
func (v *Vertice) UpdateEdge(e *Edge) error {
	return v.Edges.Add(e)
}

// Delete an edge to vertice.
func (v *Vertice) DeleteEdge(e *Edge) error {
	return v.Edges.Delete(e)
}

// If vertice only has type TO edges, which means it's a parent to other vertices only.
func (v *Vertice) IsParentOnly() bool {
	for e := range v.Edges.IterChan() {
		if e.Type == FROM {
			return false
		}
	}

	return true
}

// If vertice only has type FROM edges, which means it's a child of other vertices only.
func (v *Vertice) IsChildOnly() bool {
	for e := range v.Edges.IterChan() {
		if e.Type == TO {
			return false
		}
	}

	return true
}

// Compare two vertices based on value internal ID.
func (v *Vertice) Equal(double *Vertice) bool {
	return v.Value.Id() == double.Value.Id()
}
