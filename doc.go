// Package ds provides data structures that are not implemented by Go's built-in library.
//
// Graph
//
// The graph module provides functions and utilities to construct a graph. It also allows flexible
// extension to the library using which developers can construct custom stores, algorithms or custom
// business logic.
//
// Below is an example of create a three nodes grph using built-in list store.
//
// Notes:
//
// 1. Any changes to the vertices or edges will need to call UpdateVertice to trigger the graph update.
//
// 2. Vertices needed to be exist in the graph if being referennced by another vertice.
//
// 3. Only one edge reference required when adding a new vertice. The AddVertice/UpdateVertice function
// will automatically add the opposite reference to the neighboring vertices. For example, we have
// vertice A and B, only edge A=>B with TO type needed to added to A, after calling the add or update
// function, A=>B with From type will be automatically added to vertice B.
//
//   package main
//
//   import (
//     "fmt"
//
//     "github.com/liangrog/ds/graph/parts"
//     "github.com/liangrog/ds/graph/list"
//     "github.com/liangrog/ds/graph/sort"
//   )
//
//   func main() {
//     // Create a directed graph
//     g := parts.NewGraph(parts.DIRECTED, list.NewVerticeStore())
//
//     // Create three vertices
//     v1 := parts.NewVertice("1", list.NewEdgeStore())
//     v2 := parts.NewVertice("2", list.NewEdgeStore())
//     v3 := parts.NewVertice("3", list.NewEdgeStore())
//
//     g.AddVertice(v1)
//     g.AddVertice(v2)
//     g.AddVertice(v3)
//
//     // Add edges to each vertices.
//
//     // v1 => v2
//     e12 := parts.NewEdge("12", v2, parts.TO)
//     v1.AddEdge(e12)
//     g.UpdateVertice(v1)
//
//     // v2 => v3
//     e23 := parts.NewEdge("23", v3, parts.TO)
//     v2.AddEdge(e23)
//     g.UpdateVertice(v2)
//
//     fmt.Printf("%s", g)
//
//     // Let's Kahn sort it
//     _, sorted, _ = sort.Kahn(g)
//     fmt.Println(sorted)
//   }
//
// How to create your own store for vertices and edges
//
// There are only two interfaces required for your implementation: parts.VerticeStore and
// parts.EdgeStore. You can define any indexers within your store.
//
package ds

// Make godoc list the subdirectories
import (
	_ "github.com/liangrog/ds/graph/intrfs"
	_ "github.com/liangrog/ds/graph/list"
	_ "github.com/liangrog/ds/graph/parts"
	_ "github.com/liangrog/ds/graph/sort"
	_ "github.com/liangrog/ds/graph/utils"
)
