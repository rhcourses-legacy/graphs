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

	gt := GraphTester{n, t}

	gt.assertReachable_maxdist("A", 0)
	gt.assertReachable_maxdist("B", 1)
	gt.assertReachable_maxdist("C", 1)
	gt.assertReachable_maxdist("D", 2)
	gt.assertReachable_maxdist("E", 2)
	gt.assertReachable_maxdist("F", 2)
	gt.assertReachable_maxdist("G", 2)

	gt.assertUnreachable_maxdist("G", 1)
	gt.assertUnreachable_maxdist("H", 2)
}
