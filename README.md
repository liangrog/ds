# ds
Package ds provides data structures that are not implemented by Go's built-in library.

## Documentations
Please see the package [godoc](https://godoc.org/github.com/liangrog/ds)

## Example
Graph
```go
package main

import (
    "fmt"

    "github.com/liangrog/ds/graph/parts"
    "github.com/liangrog/ds/graph/list"
    "github.com/liangrog/ds/graph/sort"
)

func main() {
    // Create a directed graph
    g := parts.NewGraph(parts.DIRECTED, list.NewVerticeStore())

    // Create three vertices
    v1 := parts.NewVertice("1", list.NewEdgeStore())
    v2 := parts.NewVertice("2", list.NewEdgeStore())
    v3 := parts.NewVertice("3", list.NewEdgeStore())

    g.AddVertice(v1)
    g.AddVertice(v2)
    g.AddVertice(v3)

    // Add edges to each vertices.

    // v1 => v2
    e12 := parts.NewEdge("12", v2, parts.TO)
    v1.AddEdge(e12)
    g.UpdateVertice(v1)

    // v2 => v3
    e23 := parts.NewEdge("23", v3, parts.TO)
    v2.AddEdge(e23)
    g.UpdateVertice(v2)

    fmt.Printf("%s", g)

    // Let's Kahn sort it
    _, sorted, _ = sort.Kahn(g)
    fmt.Println(sorted)
}
```
