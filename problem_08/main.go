package main

import (
	"fmt"
	"main/duplicateIndex"
)

func main() {
	testCases := [][]int{
		{5, 17, 3, 17, 4, -1},
		{1, 2, 3, 4, 5},
		{9, 7, 5, 7, 8, 9},
		{1, 1, 2, 3, 4, 5},
		{},
	}

	for _, arr := range testCases {
		fmt.Println("Input array:", arr)
		fmt.Println("Index of first duplicate:", duplicateIndex.FirstDuplicateIdx(arr))
		fmt.Println()
	}
}
