package main

import (
	"fmt"
	"main/runLengthEncodeDecode"
)

func main() {
	testCases := []string{
		"aaaaaaaaaabbbaxxxxyyyzyx",
		"abcd",
		"aaa",
		"",
		"x",
		"zzzzzzzzzzz",
	}

	fmt.Println("Testing runLengthEncode/Decode Functions:")
	for _, test := range testCases {
		encoded := runLengthEncodeDecode.RunLengthEncode(test)
		decoded := runLengthEncodeDecode.RunLengthDecode(encoded)

		fmt.Printf("\tOriginal: %-20s | Encoded: %-20s | Decoded: %-20s | Match: %v\n",
			test, encoded, decoded, test == decoded)
	}
}
