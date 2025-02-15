package main

import (
	"fmt"
	"main/uniqueStrings"
)

func main() {
	testCases := [][]string{
		{"apples", "oranges", "flowers", "apples"},
		{"apples", "apples"},
		{"apple", "banana", "cherry"},
		{"Apple", "apple", "Banana", "banana"},
		{},
	}

	fmt.Println("Unique Strings Test:")
	for _, test := range testCases {
		result := uniqueStrings.UniqueStrings(test)
		fmt.Printf("\tInput: %v | Unique: %v\n", test, result)
	}
}
