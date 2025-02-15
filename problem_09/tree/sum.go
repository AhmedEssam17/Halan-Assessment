package tree

/*
	===========================================================================================
	Proposed Solution
	===========================================================================================
	Procedure:
		- Create a struct for node with elements
			= value as int
			= children as array of node structs
		- Function sum works in a recursive manner as follows:
			= Base case to return 0 if node == nil
			= Initialize total with value found in the node
			= Iterate over the arr of children and run sum recursively
*/

type Node struct {
	Value    int
	Children []*Node
}

func Sum(root *Node) int {
	// Base case to return 0 if node == nil
	if root == nil {
		return 0
	}

	// Initialize total with value found in the node
	total := root.Value

	// Iterate over the arr of children and run sum recursively
	for _, child := range root.Children {
		total += Sum(child) // Recursively sum all child nodes
	}

	return total
}
