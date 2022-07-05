package graph

// Node defines a node in a graph.
// Each node has a reference to its parent and children nodes.
// The ID is used to identify the node and the value can be used to store any data associated with the node.
// The weight of a node is used to determine the order in which nodes are traversed.
type Node[id comparable, v any] struct {
	parents  []*Node[id, v]
	children []*Node[id, v]
	weight   int
	id       id
	value    v
}

// NewNode creates a new node with the given ID and value.
func NewNode[id comparable, v any](name id, value v) *Node[id, v] {
	return &Node[id, v]{
		parents:  []*Node[id, v]{},
		children: []*Node[id, v]{},
		weight:   0,
		id:       name,
		value:    value,
	}
}

// AddChild adds a child node to the current node.
func (n *Node[id, v]) AddChild(child *Node[id, v]) {
	n.children = append(n.children, child)
	child.parents = append(child.parents, n)
	// set the weight of the child node
	if child.weight < n.weight+1 {
		child.weight = n.weight + 1
	}
}

// WalkDFS traverses the graph in depth-first order.
func (n *Node[id, v]) WalkDFS(f func(*Node[id, v])) {
	f(n)
	for _, child := range n.children {
		child.WalkDFS(f)
	}
}

// WalkBFS traverses the graph in breadth-first order.
func (n *Node[id, v]) WalkBFS(f func(*Node[id, v])) {
	queue := []*Node[id, v]{n}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		f(node)
		queue = append(queue, node.children...)
	}
}

// Order returns the order of the node in the graph.
func (n *Node[id, v]) Order() map[int][]Node[id, v] {
	ordered := map[int][]Node[id, v]{}
	for _, child := range n.children {
		for key, value := range child.Order() {
			ordered[key] = append(ordered[key], value...)
		}
	}
	ordered[n.weight] = append(ordered[n.weight], *n)
	return ordered
}

// ID returns the ID of the node.
func (n *Node[id, v]) ID() id {
	return n.id
}

// Value returns the value of the node.
func (n *Node[id, v]) Value() v {
	return n.value
}

// Weight returns the weight of the node.
func (n *Node[id, v]) Weight() int {
	return n.weight
}

// Parents returns the parents of the node.
func (n *Node[id, v]) Parents() []*Node[id, v] {
	return n.parents
}

// Children returns the children of the node.
func (n *Node[id, v]) Children() []*Node[id, v] {
	return n.children
}
