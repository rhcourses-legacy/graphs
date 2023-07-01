package graphs

import "testing"

// TestEdge_New tests the NewEdge function.
// It creates a node and an edge pointing to it.
// Checks that the edge is valid and points to the node.
func TestEdge_New(t *testing.T) {
	n := NewNode("A")
	e := NewEdge(n)
	if !e.IsValid() {
		t.Errorf("Expected edge to be valid, got invalid edge")
	}
	if e.To != n {
		t.Errorf("Expected edge to point to node A, got %v", e.To)
	}
}
