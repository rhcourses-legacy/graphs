package graphs

// Node is a struct containing a label and a slice of edges.
// It is used to represent a node in a graph.
type Node struct {
	Label string
	Edges []Edge
}

// New creates a new Node.
func NewNode(label string) *Node {
	return &Node{Label: label}
}

// String returns the label of the node.
func (n *Node) String() string {
	return n.Label
}

// NeighbourCount returns the number of neighbours of a node.
func (n *Node) NeighbourCount() int {
	return len(n.Edges)

	// Die Anzahl der Nachbarn ist die Anzahl der Kanten, die von diesem Knoten ausgehen.
}

// GetNeighbour returns the neighbour with the given label.
// If no neighbour with the given label exists, nil is returned.
func (n *Node) GetNeighbour(label string) *Node {
	for _, e := range n.Edges {
		if e.To.Label == label {
			return e.To
		}
	}
	return nil

	// Laufe durch alle Kanten und prüfe, ob der Zielknoten der Kante das gesuchten Label hat.
}

// AddNeighbour adds a neighbour with a label to a node.
// If a Neighbour with the same label already exists, it is not added.
func (n *Node) AddNeighbour(label string) {
	n.AddNeighbourNode(NewNode(label))

	// Erstelle einen neuen Knoten mit dem Label und füge ihn als Nachbarn hinzu.
	// Nutze dafür die Methode AddNeighbourNode.
}

// AddNeighbourNode adds node m as a neighbour.
// If a node with that label already exists, it is not added.
// If m is nil, it is not added.
func (n *Node) AddNeighbourNode(m *Node) {
	if m == nil || n.GetNeighbour(m.Label) != nil {
		return
	}
	n.Edges = append(n.Edges, *NewEdgeNode(m))

	// Füge eine neue Kante mittels NewEdgeNode von diesem Knoten zu m hinzu.
	// Prüfe vorher, ob der neue Knoten überhaupt existiert und ob noch kein
	// Knoten mit diesem Label vorhanden ist.
	// Die letztere Prüfung kann mittels GetNeighbour erfolgen.
}

// NeighboursMaxDistance returns all neighbours of a node with a distance of at most maxDistance.
func (n *Node) NeighboursMaxDistance(maxDistance int) []*Node {
	neighbours := []*Node{}
	for _, e := range n.Edges {
		neighbours = append(neighbours, e.To)
		if maxDistance > 1 {
			neighbours = append(neighbours, e.To.NeighboursMaxDistance(maxDistance-1)...)
		}
	}

	/* Zusatzanforderung aus den Tests mit Graphen, die Kreise enthalten:
	 * Duplikate entfernen!
	 */
	result := []*Node{}
	for _, neighbour := range neighbours {
		duplicate := false
		for _, node := range result {
			if neighbour == node {
				duplicate = true
			}
		}
		if !duplicate {
			result = append(result, neighbour)
		}
	}

	return result

	// Erstelle eine Liste für die Nachbarn und füge alle direkten Nachbarn (also Zielknoten der Kanten) hinzu.
	// Falls maxDistance > 1 ist, rufe diese Funktion rekursiv mit jedem der direkten Nachbarn und maxDistance-1 auf.
	// Für die Zusatzanforderung aus den Tests mit Graphen, die Kreise enthalten: Entferne Duplikate aus der Liste.
}

// CanReachLabel_MaxDepth expects a label and a maximal search depth.
// returns true if a node with that label is reachable.
func (n *Node) CanReachLabel_MaxDepth(label string, maxDepth int) bool {
	if label == n.Label {
		return true
	}
	if maxDepth == 0 {
		return false
	}
	for _, e := range n.Edges {
		if e.To.CanReachLabel_MaxDepth(label, maxDepth-1) {
			return true
		}
	}
	return false

	// Rekursive Tiefensuche
	// Prüfe, ob der aktuelle Knoten das gesuchte Label hat oder ob maxDepth == 0 ist.
	// Dies sind die Basisfälle, in denen die Suche abgebrochen wird.
	// Falls nicht, rufe diese Funktion rekursiv mit jedem der direkten Nachbarn und maxDepth-1 auf.
}

// CanReachLabel expects a label.
// returns true if a node with that label is reachable in any number of steps.
func (n *Node) CanReachLabel(label string) bool {
	lastneighbours := []*Node{}
	currentneighbours := n.NeighboursMaxDistance(1)

	for depth := 2; len(lastneighbours) != len(currentneighbours); depth++ {
		lastneighbours = currentneighbours
		currentneighbours = n.NeighboursMaxDistance(depth)
	}

	for _, node := range currentneighbours {
		if node.Label == label {
			return true
		}
	}
	return false

	// Im Gegensatz zu CanReachLabel_MaxDepth wird hier nicht rekursiv vorgegangen.
	// Für die Knoten, die tatsächlich erreichbar sind, würde ein rekursiver Ansatz
	// ähnlich wie bei CanReachLabel_MaxDepth funktionieren.
	// Für Knoten, die nicht erreichbar sind, bricht dann aber die Rekursion nicht ab.

	// Daher wird hier ein iterativer Ansatz gewählt:
	// Erstelle zwei Listen:
	// * lastneighbours enthält alle Knoten, die in der letzten Iteration gefunden wurden.
	// * currentneighbours enthält alle Knoten, die in der aktuellen Iteration gefunden wurden.
	// Die Suche berechnet nun immer weiter die Nachbarn der Nachbarn, bis sich die beiden Listen nicht mehr unterscheiden.
	// Dazu kann die Funktion NeighboursMaxDistance in einer Schleife verwendet werden.
	// Wenn die beiden Listen gleich sind, ist die Suche abgeschlossen,
	// da dann keine neuen Knoten mehr gefunden werden können.
	// Prüfe nun, ob der gesuchte Knoten in der Liste der aktuellen Nachbarn enthalten ist.
}
