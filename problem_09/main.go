package main

import (
	"fmt"
	"main/tree"
)

func main() {
	root := &tree.Node{Value: 10, Children: []*tree.Node{
		{Value: 5, Children: []*tree.Node{
			{Value: 2, Children: []*tree.Node{}},
			{Value: 1, Children: []*tree.Node{}},
		}},
		{Value: 3, Children: []*tree.Node{
			{Value: 4, Children: []*tree.Node{}},
		}},
	}}

	fmt.Println("Sum of tree values:", tree.Sum(root)) // Expected output: 25
}
