package palindrome

import "strings"

/*
	===========================================================================================
										Problem 01
	===========================================================================================
	A palindrome is a word or sequence of characters which reads the same backward or forward.
	For example the words: level, anna, noon, rotator are all palindromes.
	Write a function palindrom that accepts a string as an argument and returns a boolean
	indicating whether the input is a palindrome or not, for example:
		palindrome("anna") # returns True
		palindrome("apple") # returns False
*/

/*
	===========================================================================================
										Proposed Solution
	===========================================================================================
	Procedure:
		1) Convert the whole string to lower case to handle Case Senstivity "Not mentioned but suggested"
		2) Remove all spaces from string "Not mentioned but to handle strings like (race car) and similar ones"
		3) Use 2 pointers to compare first and last chars till the pointers meet or cross
			return true or false accordingly.
*/

func IsPalindrome(word string) bool {
	// Convert the whole string to lowercase alphabets (Ignore Cases)
	word = strings.ToLower(word)
	// Remove all spaces from string (Ignore Spaces)
	word = strings.ReplaceAll(word, " ", "")

	leftPtr := 0
	rightPtr := len(word) - 1

	// loop until pointers meet or cross
	for leftPtr < rightPtr {
		if word[leftPtr] != word[rightPtr] {
			// return false when they don't match
			return false
		}
		leftPtr++
		rightPtr--
	}

	// Return true if the previous loop didn't break == Palindrom
	return true
}
