package graphs

import "testing"

// TestNode_New tests the NewNode function.
// It creates a new node and checks that the label is correct
// and that there are no neighbours.
func TestNode_New(t *testing.T) {
	gt := GraphTester{t}

	n := NewNode("A")

	gt.assertNodeLabelEquals(n, "A")
	gt.assertNeighbourCountEquals(n, 0)
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
	gt := GraphTester{t}

	n := NewNode("A")
	n.AddNeighbour("B")
	n.AddNeighbour("C")
	gt.assertNeighbourCountEquals(n, 2)
}

// TestNode_GetNeighbour tests the GetNeighbour method.
// It creates a new node, adds neighbours manually, and checks the following:
// - the neighbours are correct
// - for a non-existent neighbour, nil is returned
func TestNode_GetNeighbour(t *testing.T) {
	gt := GraphTester{t}
	n := NewNode("A")
	n.Edges = append(n.Edges, *NewEdge("B"))
	n.Edges = append(n.Edges, *NewEdge("C"))

	gt.assertNeighbourHasCorrectLabel(n, "B")
	gt.assertNeighbourHasCorrectLabel(n, "C")
	gt.assertNeighbourIsNode(n, nil, "D")
}

// TestNode_AddNeighbour tests the AddNeighbour method.
// It performs the same test as TestNode_GetNeighbour, but uses AddNeighbour to create the nodes.
// Additionally checks that an existing label isn't added twice.
func TestNode_AddNeighbour(t *testing.T) {
	gt := GraphTester{t}

	n := NewNode("A")
	n.AddNeighbour("B")
	n.AddNeighbour("C")

	gt.assertNeighbourHasCorrectLabel(n, "B")
	gt.assertNeighbourHasCorrectLabel(n, "C")

	n.AddNeighbour("C")
	gt.assertNeighbourHasCorrectLabel(n, "C")
	gt.assertNeighbourCountEquals(n, 2)
}

// TestNode_AddNeighbourNode tests the AddNeighbourNode method.
// It performs the same test as TestNode_AddNeighbour, but uses nodes instead of labels.
// Additionally checks that an existing node isn't added twice.
func TestNode_AddNeighbourNode(t *testing.T) {
	gt := GraphTester{t}

	a := NewNode("A")
	b := NewNode("B")
	c := NewNode("C")

	a.AddNeighbourNode(b)
	a.AddNeighbourNode(c)

	gt.assertNeighbourIsNode(a, b, "B")
	gt.assertNeighbourIsNode(a, c, "C")

	gt.assertNeighbourCountEquals(a, 2)
	gt.assertNeighbourCountEquals(b, 0)
	gt.assertNeighbourCountEquals(c, 0)

	a.AddNeighbourNode(c)
	gt.assertNeighbourIsNode(a, c, "C")
}
