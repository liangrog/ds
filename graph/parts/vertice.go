package parts

// Vertice for the graph
type Vertice struct {
	// Value can be anything user defined
	Value *Value

	// All edges that come from or
	// go out to from this vertice.
	// Note each edge is referenced by
	// both from vertice and to vertice
	// so there is no traverse required for
	// looking up for parent or child.
	Edges EdgeStore
}

// Vertice constructor
func NewVertice(value interface{}, edgeStore EdgeStore) *Vertice {
	return &Vertice{
		Value: NewValue(value),
		Edges: edgeStore,
	}
}

// Delete all edge reference to a neighbor vertice
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

func (v *Vertice) GetNeighbors() []*Vertice {
	var res []*Vertice
	for e := range v.Edges.IterChan() {
		res = append(res, e.Neighbor)
	}

	return res
}

func (v *Vertice) AddEdge(e *Edge) error {
	return v.Edges.Add(e)
}

func (v *Vertice) DeleteEdge(e *Edge) error {
	return v.Edges.Delete(e)
}

func (v *Vertice) IsParentOnly() bool {
	for e := range v.Edges.IterChan() {
		if e.Type == FROM {
			return false
		}
	}

	return true
}

func (v *Vertice) IsChildOnly() bool {
	for e := range v.Edges.IterChan() {
		if e.Type == TO {
			return false
		}
	}

	return true
}

func (v *Vertice) Equal(double *Vertice) bool {
	return v.Value.Id() == double.Value.Id()
}
