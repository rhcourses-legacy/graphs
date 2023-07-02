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
	n.AddNeighbourNode(NewNode(label))
}

// AddNeighbourNode adds node m as a neighbour.
// If a node with that label already exists, it is not added.
// If m is nil, it is not added.
func (n *Node) AddNeighbourNode(m *Node) {
	if m == nil || n.GetNeighbour(m.Label) != nil {
		return
	}
	n.Edges = append(n.Edges, *NewEdgeNode(m))
}

// NeighboursMaxDistance returns all neighbours of a node with a distance of at most maxDistance.
func (n *Node) NeighboursMaxDistance(maxDistance int) []*Node {
	neighbours := []*Node{}
	for _, e := range n.Edges {
		neighbours = append(neighbours, e.To)
		if maxDistance > 1 {
			neighbours = append(neighbours, e.To.NeighboursMaxDistance(maxDistance-1)...)
		}
	}
	return neighbours
}

// CanReachLabel_MaxDepth expects a label and a maximal search depth.
// returns true if a node with that label is reachable.
func (n *Node) CanReachLabel_MaxDepth(label string, maxDepth int) bool {
	if label == n.Label {
		return true
	}
	if maxDepth == 0 {
		return false
	}
	for _, e := range n.Edges {
		if e.To.CanReachLabel_MaxDepth(label, maxDepth-1) {
			return true
		}
	}
	return false
}
