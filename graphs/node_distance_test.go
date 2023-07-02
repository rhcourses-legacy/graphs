package graphs

import "testing"

// TestNode_NeighboursMaxDistance tests the NeighboursMaxDistance method.
func TestNode_NeighboursMaxDistance(t *testing.T) {
	n := NewNode("A")
	n.AddNeighbour("B")
	n.AddNeighbour("C")

	B := n.GetNeighbour("B")
	B.AddNeighbour("D")
	B.AddNeighbour("E")

	C := n.GetNeighbour("C")
	C.AddNeighbour("F")
	C.AddNeighbour("G")

	result1 := n.NeighboursMaxDistance(1)
	expectedLabels1 := []string{"B", "C"}
	if !compareLabels(result1, expectedLabels1) {
		t.Errorf("Expected %v, got %v", expectedLabels1, result1)
	}

	result2 := n.NeighboursMaxDistance(2)
	// Anmerkung: Die Reihenfolge kann je nach Implementierung auch anders sein.
	//            Wie wäre es lösbar, dass dies dem Test egal ist?
	expectedLabels2 := []string{"B", "D", "E", "C", "F", "G"}
	if !compareLabels(result2, expectedLabels2) {
		t.Errorf("Expected %v, got %v", expectedLabels2, result2)
	}
}

// TestNode_CanReachLabel_MaxDepth tests the CanReachLabel_MaxDepth method.
func TestNode_CanReachLabel(t *testing.T) {
	n := NewNode("A")
	n.AddNeighbour("B")
	n.AddNeighbour("C")

	B := n.GetNeighbour("B")
	B.AddNeighbour("D")
	B.AddNeighbour("E")

	C := n.GetNeighbour("C")
	C.AddNeighbour("F")
	C.AddNeighbour("G")

	if !n.CanReachLabel_MaxDepth("A", 0) {
		t.Errorf("Expected node A to be reachable with max depth 0.")
	}
	if !n.CanReachLabel_MaxDepth("B", 1) {
		t.Errorf("Expected node B to be reachable with max depth 1.")
	}
	if !n.CanReachLabel_MaxDepth("C", 1) {
		t.Errorf("Expected node C to be reachable with max depth 1.")
	}
	if !n.CanReachLabel_MaxDepth("D", 2) {
		t.Errorf("Expected node E to be reachable with max depth 2.")
	}
	if !n.CanReachLabel_MaxDepth("E", 2) {
		t.Errorf("Expected node E to be reachable with max depth 2.")
	}
	if !n.CanReachLabel_MaxDepth("F", 2) {
		t.Errorf("Expected node F to be reachable with max depth 2.")
	}
	if !n.CanReachLabel_MaxDepth("G", 2) {
		t.Errorf("Expected node G to be reachable with max depth 2.")
	}
	if n.CanReachLabel_MaxDepth("G", 1) {
		t.Errorf("Expected node G not to be reachable with max depth 1.")
	}
	if n.CanReachLabel_MaxDepth("H", 2) {
		t.Errorf("Expected node H not to be reachable with max depth 2.")
	}
}
