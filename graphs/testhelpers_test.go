package graphs

import "testing"

// compareLabels compares the labels of a slice of nodes to a slice of strings.
// It returns true if the labels are equal, false otherwise.
func compareLabels(nodes []*Node, labels []string) bool {
	if len(nodes) != len(labels) {
		return false
	}
	for i, n := range nodes {
		if n.Label != labels[i] {
			return false
		}
	}
	return true
}

// Helper struct for encapsulating test data that is required repeatedly
// when checking properties of a node.
type GraphTester struct {
	n *Node
	t *testing.T
}

// assertReachable_maxdist expects a label and a maximum distance md.
// Checks if a node with the givel label can be reached in at most md steps.
// If not, produces a test failure.
func (gt GraphTester) assertReachable_maxdist(label string, md int) {
	if !gt.n.CanReachLabel_MaxDepth(label, md) {
		gt.t.Errorf("Expected node %s to be reachable with max depth %d.", label, md)
	}
}

// assertReachable_maxdist expects a label and a maximum distance md.
// Checks if a node with the givel label cannot be reached in at most md steps.
// If it can be reached, produces a test failure.
func (gt GraphTester) assertUnreachable_maxdist(label string, md int) {
	if gt.n.CanReachLabel_MaxDepth(label, md) {
		gt.t.Errorf("Expected node %s not to be reachable with max depth %d.", label, md)
	}
}
