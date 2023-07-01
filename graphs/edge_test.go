package graphs

import "testing"

// TestEdge_New tests the NewEdge function.
// It creates an edge pointing to to a label.
// Checks that the edge is valid and points to a node with that label.
func TestEdge_New(t *testing.T) {
	e := NewEdge("A")
	if !e.IsValid() {
		t.Errorf("Expected edge to be valid, got invalid edge")
	}
	if e.To.Label != "A" {
		t.Errorf("Expected edge to point to node A, got %v", e.To)
	}
}
