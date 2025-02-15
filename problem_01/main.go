package main

import (
	"fmt"
	"main/palindrome"
)

/*
	Example Inputs with Expected Outputs:
		Testing Palindrome Function:
			isPalindrome("rotator") 	=> true
			isPalindrome("Level") 		=> true
			isPalindrome("race car")	=> true
			isPalindrome("apple") 		=> false
			isPalindrome("Essam") 		=> false
			isPalindrome("Anna") 		=> true
			isPalindrome("a") 			=> true
			isPalindrome("") 			=> true
*/

func main() {
	words := []string{
		"rotator",
		"Level",
		"race car",
		"apple",
		"Essam",
		"Anna",
		"a",
		"",
	}

	fmt.Println("Testing Palindrome Function:")
	for _, word := range words {
		fmt.Printf("\tisPalindrome(%q) => %t\n", word, palindrome.IsPalindrome(word))
	}
}
