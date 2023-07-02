package graphs

import "testing"

// TestNode_NeighboursMaxDistance_cycles tests the NeighboursMaxDistance method.
// For testing, it uses a graph that contains cycles.
//
// We expect the method to return a list without duplicates!
// Graphs that are not trees (e.g. with cycles) may have duplicate neighbour entries.
func TestNode_NeighboursMaxDistance_cycles(t *testing.T) {
	gt := GraphTester{t}
	a := NewNode("A")
	b := NewNode("B")
	c := NewNode("C")

	a.AddNeighbourNode(b)
	b.AddNeighbourNode(c)
	c.AddNeighbourNode(a)

	gt.assertNodesEqual(a.NeighboursMaxDistance(1), b)
	gt.assertNodesEqual(a.NeighboursMaxDistance(2), b, c)
	gt.assertNodesEqual(a.NeighboursMaxDistance(3), b, c, a)
	gt.assertNodesEqual(a.NeighboursMaxDistance(4), b, c, a)
}

// TestNode_CanReachLabel_MaxDepth tests the CanReachLabel_MaxDepth method.
// For testing, it uses a graph that contains cycles.
func TestNode_CanReachLabel_MaxDepth_cycles(t *testing.T) {
	gt := GraphTester{t}

	a := NewNode("A")
	b := NewNode("B")
	c := NewNode("C")

	a.AddNeighbourNode(b)
	b.AddNeighbourNode(c)
	c.AddNeighbourNode(a)

	gt.assertReachable_maxdist(a, "A", 0)
	gt.assertReachable_maxdist(a, "B", 1)
	gt.assertReachable_maxdist(a, "C", 2)
}

// TestNode_CanReachLabel_MaxDepth tests the CanReachLabel_MaxDepth method.
// For testing, it uses a graph that contains cycles.
func TestNode_CanReachLabel_cycles(t *testing.T) {
	gt := GraphTester{t}

	a := NewNode("A")
	b := NewNode("B")
	c := NewNode("C")

	a.AddNeighbourNode(b)
	b.AddNeighbourNode(c)
	c.AddNeighbourNode(a)

	gt.assertReachable(a, "A")
	gt.assertReachable(a, "B")
	gt.assertReachable(a, "C")
	gt.assertUnreachable(a, "D")

}
