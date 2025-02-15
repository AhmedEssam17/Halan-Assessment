package main

import (
	"fmt"
	"main/composition"
)

func main() {
	testCases := []int{6, 0, -3, 10, 100}

	fmt.Println("Function Composition Test:")
	for _, test := range testCases {
		composedFunction := composition.Compose(composition.Square, composition.Inc)
		result := composedFunction(test)

		fmt.Printf("\tInput: %-3d | Composed: %-5d\n",
			test, result)
	}
}
