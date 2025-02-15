package main

import (
	"fmt"
	"main/transpose"
)

func main() {
	testCases := [][][]int{
		{{1, 2}, {3, 4}},
		{{1, 2, 3, 4}, {5, 6, 7, 8}},
		{{1, 2, 3, 4}},
		{{1}, {2}, {3}, {4}},
		{},
	}

	for _, matrix := range testCases {
		fmt.Println("Original matrix:", matrix)
		fmt.Println("Transposed matrix:", transpose.TransposeMatrix(matrix))
		fmt.Println()
	}
}
