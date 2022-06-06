# graph

Graph implements a generic node graph data structure. On top of the data structure, it provides the following methods:

- AddChild(child *Node[id, v])
- WalkDFS(f func(*Node[id, v]))
- WalkBFS(f func(*Node[id, v]))
- Order() map[int][]Node[id, v]
- ID() id
- Value() v
- Weight() int

## Example

```go
package main

import (
    "fmt"

    "github.com/leonsteinhaeuser/graph"
)

type NodePayload struct {
    ID   int
    Name string
}

func main() {
    rootNode := graph.NewNode("A", NodePayload{ID: 1, Name: "root"})
    node1 := graph.NewNode("B", NodePayload{ID: 2, Name: "node1"})
    node2 := graph.NewNode("C", NodePayload{ID: 3, Name: "node2"})

    node3 := graph.NewNode("D", NodePayload{ID: 4, Name: "node3"})
    node4 := graph.NewNode("E", NodePayload{ID: 5, Name: "node4"})
    node5 := graph.NewNode("F", NodePayload{ID: 6, Name: "node5"})

    node1.AddChild(node3)
    node1.AddChild(node4)
    node2.AddChild(node5)
    node3.AddChild(node5)

    rootNode.AddChild(node1)
    rootNode.AddChild(node2)

    rootNode.WalkDFS(func(n *graph.Node[string, NodePayload]) {
        fmt.Printf("%s: %v\n", n.ID(), n.Value())
    })
}
```
