package graphs

// Edge is a struct containing a pointer to a node.
// It is used to represent an outgoing edge from a node.
type Edge struct {
	To *Node
}

// NewEdge creates a new Edge.
// It expects a label and creates a new node with that label.
func NewEdge(label string) *Edge {
	return &Edge{To: NewNode(label)}
}

// NewEdgeNode expects a node and creates a new Edge with that node.
func NewEdgeNode(n *Node) *Edge {
	return &Edge{To: n}
}

// IsValid checks if the internal Node pointer is set.
func (e *Edge) IsValid() bool {
	return e.To != nil
}
