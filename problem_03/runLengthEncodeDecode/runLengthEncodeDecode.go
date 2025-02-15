package runLengthEncodeDecode

import (
	"strconv"
	"unicode"
)

/*
	===========================================================================================
	Proposed Solution
	===========================================================================================
	Procedure:
		1) runLengthEncode:
			For the encode function here is the way of thinking i followed:
				- Base case to return empty string if the len ==
				- Check for invalid input (numbers in original string) if found return empty string
				- Store first char in "storedChar" and start counter by 1
				- Iterate over the len of string
				- Check whether the following characters matching the stored character
					= if yes, then increment the counter by 1
					= if no, then its a new character, so we concatenate the stored char with
					  the current counter value into the encoded string, reset the counter to 1
					  and store the current char into storedChar.
				- After finishing the iteration, concatenate the last character with its value
				  into the encoded string.
		2) runLengthDecode:
			For the decode function here is the way of thinking i followed:
				- Base case to return empty string if the len == 0
				- Store first char in "storedChar" and initialize an empty string for the counter
				  (counter here is a string to handle multiple digits for chars as [a10, b27, ...])
				- Iterate over the len of string.
				- Check whether the following index is an alphabet
					= if yes, then convert the counter string into integer and expand the character
					  according to its count and concatenate it into the decoded string, reset the
					  counter to empty string and store the current alphabet into storedChar.
					= if no, then it could be the first digit or the second digit of a multiple digit
					  number (> 9), so concatenate it into the counter string.
				- After finishing the iteration, check whether counter still has a value and expand
				  the last alphabet along and concatenate it into the decodedString.
*/

func RunLengthEncode(str string) string {
	// Base case to return empty string if the len == 0
	if len(str) == 0 {
		return ""
	}

	// Check for invalid input (numbers in original string) if found return empty string
	for _, char := range str {
		if unicode.IsDigit(char) {
			return ""
		}
	}

	// Store first char in "storedChar" and start counter by 1
	storedChar := str[0]
	counter := 1
	encodedString := ""

	for i := 1; i < len(str); i++ {
		if str[i] != storedChar {
			// Current character doesn't match the storedChar
			// Concatenate the stored char with the current counter value into the encoded string
			encodedString += (string(storedChar) + strconv.Itoa(counter))

			// Reset the counter to 1 and store the current character
			counter = 1
			storedChar = str[i]
		} else {
			// Current character doesn't match the storedChar
			// Increment the counter by 1
			counter++
		}
	}

	// concatenate the last character with its value into the encoded string.
	encodedString += (string(storedChar) + strconv.Itoa(counter))

	return encodedString
}

func RunLengthDecode(str string) string {
	// Base case to return empty string if the len == 0
	if len(str) == 0 {
		return ""
	}

	// Store first char in "storedChar" and initialize an empty string for the counter
	storedChar := str[0]
	decodedString := ""
	counter := ""

	for i := 1; i < len(str); i++ {
		// Check whether this index is an alphabet
		if unicode.IsLetter(rune(str[i])) {
			// Convert the counter string into integer
			count, err := strconv.Atoi(counter)
			if err != nil || count <= 0 {
				return "" // Invalid encoding format
			}

			//expand the character according to its count and concatenate it into the decoded string
			expandString(&decodedString, storedChar, count)

			// Reset counter to empty string & storedChar with current alphabet
			counter = ""
			storedChar = str[i]
		} else {
			// Handles multiple digits as for (a10, b25, ...)
			counter += string(str[i])
		}
	}

	// Expand the last alphabet along and concatenate it into the decodedString.
	count, err := strconv.Atoi(counter)
	if err != nil || count <= 0 {
		return "" // Invalid encoding format
	}
	expandString(&decodedString, storedChar, count)

	return decodedString
}

// Function to repeat the char accoring to its count and concatenate it into the decoded string
func expandString(decodedString *string, char byte, count int) {
	strToAdd := ""
	for i := 0; i < count; i++ {
		strToAdd += string(char)
	}
	*decodedString += strToAdd
}
