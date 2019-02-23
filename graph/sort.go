package graph

import (
	"errors"
	"fmt"
)

func Kahn(g *Graph) ([]*Vertice, error) {
	var sorted []*Vertice
	var p []*Vertice
	for _, v := range g.Vertices {
		parentNodeOnly := true
		for _, e := range v.Edges {
			if e.Direction == EDGE_UNDIR {
				return nil, errors.New("Cannot perform Kahn topological sort on undirectional graph")
			}

			// If there is edge from other nodes
			if e.Direction == EDGE_FROM {
				parentNodeOnly = false
				break
			}
		}

		if parentNodeOnly {
			p = append(p, v)
		}
	}

	for len(p) > 0 {
		//pop
		x := p[len(p)-1]
		fmt.Println("------ pop", x.GetValue())
		// update p
		p = p[:len(p)-1]

		sorted = append(sorted, x)

		fmt.Println("----- p length", len(p))

		for _, v := range g.Vertices {
			if g.Indexer.Equal(v, x) {
				continue
			}

			// Don't example parent nodes
			ap := false
			for _, s := range sorted {
				if g.Indexer.Equal(s, v) {
					ap = true
					break
				}
			}

			if ap {
				continue
			}

			ec1 := make([]*Edge, len(v.Edges))
			copy(ec1, v.Edges)

			fmt.Println("##### examing", v.GetValue())

			child := false
			for i, e := range ec1 {
				if e.Direction == EDGE_FROM && g.Indexer.Equal(e.Attach, x) {
					v.Edges = append(v.Edges[:i], v.Edges[i+1:]...)

					fmt.Println(x.GetValue(), "%%%% before")
					for _, tt := range x.Edges {
						fmt.Println(tt.Direction, tt.Attach.GetValue())
					}
					// remove same edge from x
					ec2 := make([]*Edge, len(x.Edges))
					copy(ec2, x.Edges)
					for ii, ee := range ec2 {
						if ee.Direction == EDGE_TO && g.Indexer.Equal(ee.Attach, v) {
							x.Edges = append(x.Edges[:ii], x.Edges[ii+1:]...)
						}
					}

					fmt.Println(x.GetValue(), "%%%% after")
					for _, tt := range x.Edges {
						fmt.Println(tt.Direction, tt.Attach.GetValue())
					}

					continue
				}

				// If still having parent node
				if e.Direction == EDGE_FROM {
					child = true
					break
				}
			}

			// If pure parent node
			if !child {
				exit := false
				for _, c := range p {
					if g.Indexer.Equal(c, v) {
						exit = true
						break
					}
				}

				if !exit {
					fmt.Println("Parent only node found, add", v.GetValue())
					p = append(p, v)
					fmt.Println("Now p length changed to", len(p))
				}
			}

		}
	}

	for _, v := range g.Vertices {
		if len(v.Edges) > 0 {
			return []*Vertice{v}, errors.New("Found circular dependency at node")
		}
	}

	return sorted, nil
}
