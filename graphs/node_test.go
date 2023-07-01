package graphs

import "testing"

// TestNode_New tests the NewNode function.
// It creates a new node and checks that the label is correct
// and that there are no neighbours.
func TestNode_New(t *testing.T) {
	n := NewNode("A")
	if n.Label != "A" {
		t.Errorf("Expected label A, got %s", n.Label)
	}
	if n.NeighbourCount() != 0 {
		t.Errorf("Expected 0 neighbours, got %d", n.NeighbourCount())
	}
}

// TestNode_String tests the String method.
// It creates a new node and checks that the label is correct.
func TestNode_String(t *testing.T) {
	n := NewNode("A")
	if n.String() != "A" {
		t.Errorf("Expected label A, got %s", n.Label)
	}
}
