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

// GetNeighbour returns the neighbour with the given label.
// If no neighbour with the given label exists, nil is returned.
func (n *Node) GetNeighbour(label string) *Node {
	for _, e := range n.Edges {
		if e.To.Label == label {
			return e.To
		}
	}
	return nil
}

// AddNeighbour adds a neighbour with a label to a node.
// If a Neighbour with the same label already exists, it is not added.
func (n *Node) AddNeighbour(label string) {
	if n.GetNeighbour(label) != nil {
		return
	}
	n.Edges = append(n.Edges, *NewEdge(label))
}
