package graphs

// Node is a struct containing a label and a slice of edges.
// It is used to represent a node in a graph.
type Node struct {
	Label string
	Edges []Edge
}

// New creates a new Node.
func NewNode(label string) *Node {
	return &Node{Label: label}
}

// String returns the label of the node.
func (n *Node) String() string {
	return n.Label
}

// NeighbourCount returns the number of neighbours of a node.
func (n *Node) NeighbourCount() int {
	return len(n.Edges)
}
