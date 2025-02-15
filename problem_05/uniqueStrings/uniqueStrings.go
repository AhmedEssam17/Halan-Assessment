package uniqueStrings

import "strings"

/*
	===========================================================================================
	Proposed Solution
	===========================================================================================
	Procedure:
		- Base case to return empty string array if the len of string array == 0
		- Create a frequency map to store occurences of strings in the array.
		- Iterate through the map, convert to lowercase to make comparison case insensitive,
		  store strings with frequence = 1 in the result array.
*/

func UniqueStrings(strArr []string) []string {
	// Base case to return empty string array if the len of string array == 0
	if len(strArr) == 0 {
		return []string{}
	}

	// Create a frequency map to store occurences of strings in the array.
	freqMap := make(map[string]int)
	for _, word := range strArr {
		// Convert to lowercase to make comparison case insensitive
		lowerCaseWord := strings.ToLower(word)
		freqMap[lowerCaseWord]++
	}

	// Iterate through the map and store strings with frequence = 1 in the result array.
	var result []string
	for _, word := range strArr {
		if freqMap[strings.ToLower(word)] == 1 {
			result = append(result, word)
		}
	}

	return result
}
