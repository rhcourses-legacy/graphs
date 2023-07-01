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

// TestNode_NeighbourCount tests the NeighbourCount method.
// It creates a new node, adds neighbours, and checks that the number of neighbours is correct.
func TestNode_NeighbourCount(t *testing.T) {
	n := NewNode("A")
	n.AddNeighbour("B")
	n.AddNeighbour("C")
	if n.NeighbourCount() != 2 {
		t.Errorf("Expected 2 neighbours, got %d", n.NeighbourCount())
	}
}

// TestNode_GetNeighbour tests the GetNeighbour method.
// It creates a new node, adds neighbours manually, and checks the following:
// - the neighbours are correct
// - for a non-existent neighbour, nil is returned
func TestNode_GetNeighbour(t *testing.T) {
	n := NewNode("A")
	n.Edges = append(n.Edges, *NewEdge("B"))
	n.Edges = append(n.Edges, *NewEdge("C"))
	if n.GetNeighbour("B").Label != "B" {
		t.Errorf("Expected neighbour B, got %v", n.GetNeighbour("B"))
	}
	if n.GetNeighbour("C").Label != "C" {
		t.Errorf("Expected neighbour C, got %v", n.GetNeighbour("C"))
	}
	if n.GetNeighbour("D") != nil {
		t.Errorf("Expected nil, got %v", n.GetNeighbour("D"))
	}
}

// TestNode_AddNeighbour tests the AddNeighbour method.
// It performs the same test as TestNode_GetNeighbour, but uses AddNeighbour to create the nodes.
// Additionally checks that an existing label isn't added twice.
func TestNode_AddNeighbour(t *testing.T) {
	n := NewNode("A")
	n.AddNeighbour("B")
	n.AddNeighbour("C")
	if n.GetNeighbour("B").Label != "B" {
		t.Errorf("Expected neighbour B, got %v", n.GetNeighbour("B"))
	}
	C := n.GetNeighbour("C")
	if C.Label != "C" {
		t.Errorf("Expected neighbour C, got %v", n.GetNeighbour("C"))
	}
	n.AddNeighbour("C")
	if C.Label != "C" {
		t.Errorf("Expected neighbour C, got %v", n.GetNeighbour("C"))
	}
	if n.NeighbourCount() != 2 {
		t.Errorf("Expected 2 neighbours, got %d", n.NeighbourCount())
	}
}

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
