package graphs

import "testing"

// TestNode_NeighboursMaxDistance tests the NeighboursMaxDistance method.
func TestNode_NeighboursMaxDistance(t *testing.T) {
	gt := GraphTester{t}
	n := NewNode("A")
	n.AddNeighbour("B")
	n.AddNeighbour("C")

	B := n.GetNeighbour("B")
	B.AddNeighbour("D")
	B.AddNeighbour("E")

	C := n.GetNeighbour("C")
	C.AddNeighbour("F")
	C.AddNeighbour("G")

	gt.assertNodesHaveLabels(n.NeighboursMaxDistance(1), "B", "C")
	gt.assertNodesHaveLabels(n.NeighboursMaxDistance(2), "B", "D", "E", "C", "F", "G")

	// Anmerkung: Die Reihenfolge kann je nach Implementierung auch anders sein.
	//            Wie wäre es lösbar, dass dies im Test keine Rolle spielt?
}

// TestNode_CanReachLabel_MaxDepth tests the CanReachLabel_MaxDepth method.
func TestNode_CanReachLabel_MaxDepth(t *testing.T) {
	gt := GraphTester{t}

	n := NewNode("A")
	n.AddNeighbour("B")
	n.AddNeighbour("C")

	B := n.GetNeighbour("B")
	B.AddNeighbour("D")
	B.AddNeighbour("E")

	C := n.GetNeighbour("C")
	C.AddNeighbour("F")
	C.AddNeighbour("G")

	gt.assertReachable_maxdist(n, "A", 0)
	gt.assertReachable_maxdist(n, "B", 1)
	gt.assertReachable_maxdist(n, "C", 1)
	gt.assertReachable_maxdist(n, "D", 2)
	gt.assertReachable_maxdist(n, "E", 2)
	gt.assertReachable_maxdist(n, "F", 2)
	gt.assertReachable_maxdist(n, "G", 2)

	gt.assertUnreachable_maxdist(n, "G", 1)
	gt.assertUnreachable_maxdist(n, "H", 2)
}
