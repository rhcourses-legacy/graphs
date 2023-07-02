package graphs

import "testing"

// Helper struct encapsulating a *testing.T object.
// Provides methods for checking graph properties.
type GraphTester struct {
	t *testing.T
}

// assertNodesHaveLabels expects a list of nodes and some strings.
// Checks that the nodes have the given strings as labels (in that order).
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

// assertReachable_maxdist expects a Node, a label and a maximum distance md.
// Checks if a node with the givel label can be reached in at most md steps.
// If not, produces a test failure.
func (gt *GraphTester) assertReachable_maxdist(n *Node, label string, md int) {
	if !n.CanReachLabel_MaxDepth(label, md) {
		gt.t.Errorf("Expected node %s to be reachable with max depth %d.", label, md)
	}
}

// assertReachable_maxdist expects a Node, a label and a maximum distance md.
// Checks if a node with the givel label cannot be reached in at most md steps.
// If it can be reached, produces a test failure.
func (gt *GraphTester) assertUnreachable_maxdist(n *Node, label string, md int) {
	if n.CanReachLabel_MaxDepth(label, md) {
		gt.t.Errorf("Expected node %s not to be reachable with max depth %d.", label, md)
	}
}
