package graphs

import (
	"reflect"
	"testing"
)

// Helper struct encapsulating a *testing.T object.
// Provides methods for checking graph properties.
type GraphTester struct {
	t *testing.T
}

// assertNodeLabelEquals
func (gt *GraphTester) assertNodeLabelEquals(n *Node, expectedlabel string) {
	if n == nil {
		gt.t.Errorf("Expected label %s, but node is nil", expectedlabel)
		return
	}
	if n.Label != expectedlabel {
		gt.t.Errorf("Expected label %s, but got %s", expectedlabel, n.Label)
	}
}

// assertNeighbourCountEquals expects a node and an int.
// Checks if the node has that number of neighbours.
// Produces a test failure if not.
func (gt *GraphTester) assertNeighbourCountEquals(n *Node, expectedCount int) {
	actualCount := n.NeighbourCount()
	if actualCount != expectedCount {
		gt.t.Errorf("Expected %d neighbours, got %d", expectedCount, actualCount)
	}
}

// assertNeighbourHasCorrectLabel expects a node and a label string.
// Checks if the node's GetNeighbour method returns a node for that label
// and that the node's label is correct.
// Produces a test failure if not.
func (gt *GraphTester) assertNeighbourHasCorrectLabel(n *Node, label string) {
	neighbour := n.GetNeighbour(label)
	if neighbour == nil {
		gt.t.Errorf("Expected neighbour %s, got nil", label)
		return
	}
	if neighbour.Label != label {
		gt.t.Errorf("Expected neighbour %s, got %s", label, neighbour.Label)
	}
}

// assertNeighbourIsNode expects two nodes n, m, and a label string.
// Checks if the node returned by n.GetNeighbour for that label is n.
// Produces a test failure if not.
func (gt *GraphTester) assertNeighbourIsNode(n, m *Node, label string) {
	neighbour := n.GetNeighbour(label)
	if neighbour != m {
		gt.t.Errorf("Expected neighbour %v, got %v", m, neighbour)
	}
}

// assertNodesHaveLabels expects a list of nodes and some strings.
// Checks that the nodes have the given strings as labels (in that order).
// None of the nodes may be nil!
// Produces a test failure if not.
func (gt *GraphTester) assertNodesHaveLabels(nodes []*Node, expectedlabels ...string) {
	actuallabels := []string{}
	for _, node := range nodes {
		actuallabels = append(actuallabels, node.Label)
	}
	if len(actuallabels) != len(expectedlabels) {
		gt.t.Errorf("Expected labels %v, got %v", expectedlabels, actuallabels)
		return
	}
	for i, expectedlabel := range expectedlabels {
		if expectedlabel != actuallabels[i] {
			gt.t.Errorf("Expected labels %v, got %v", expectedlabels, actuallabels)
			return
		}
	}
}

// assertNodesEqual expects a list of nodes and some more node parameters.
// Checks that the list matches the given nodes.
// Produces a test failure if not.
func (gt *GraphTester) assertNodesEqual(actualnodes []*Node, expectednodes ...*Node) {
	if !reflect.DeepEqual(actualnodes, expectednodes) {
		gt.t.Errorf("Expected nodes %v, got %v", expectednodes, actualnodes)
	}
}

// assertReachable_maxdist expects a node, a label and a maximum distance md.
// Checks if a node with the givel label can be reached in at most md steps.
// If not, produces a test failure.
func (gt *GraphTester) assertReachable_maxdist(n *Node, label string, md int) {
	if !n.CanReachLabel_MaxDepth(label, md) {
		gt.t.Errorf("Expected node %s to be reachable with max depth %d.", label, md)
	}
}

// assertReachable_maxdist expects a node, a label and a maximum distance md.
// Checks if a node with the givel label cannot be reached in at most md steps.
// If it can be reached, produces a test failure.
func (gt *GraphTester) assertUnreachable_maxdist(n *Node, label string, md int) {
	if n.CanReachLabel_MaxDepth(label, md) {
		gt.t.Errorf("Expected node %s not to be reachable with max depth %d.", label, md)
	}
}

// assertReachable expects a node and a label.
// Checks if a node with the givel label can be reached in any number of steps.
// If not, produces a test failure.
func (gt *GraphTester) assertReachable(n *Node, label string) {
	if !n.CanReachLabel(label) {
		gt.t.Errorf("Expected node %s to be reachable.", label)
	}
}

// assertUnreachable expects a node and a label.
// Checks if a node with the givel label can be reached in any number of steps.
// If not, produces a test failure.
func (gt *GraphTester) assertUnreachable(n *Node, label string) {
	if n.CanReachLabel(label) {
		gt.t.Errorf("Expected node %s to be unreachable.", label)
	}
}
