package graphs

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
