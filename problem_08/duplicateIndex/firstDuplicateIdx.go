package duplicateIndex

/*
	===========================================================================================
	Proposed Solution
	===========================================================================================
	Procedure:
		- Create a flag map that stores int value with bool of true or false values
		- Loop over the arr and check if num is already stored in the flagMap or not
			= if yes, this will be the first duplicate number, so return its index
			= if no, this num occured for the first time, store it in the flagMap with true value
		- Return -1 if no duplicates found
*/

func FirstDuplicateIdx(arr []int) int {
	// Create a flag map that stores int value with bool of true or false values
	flagMap := make(map[int]bool)

	for idx, num := range arr {
		// Check if num is already stored in the flagMap or not
		if flagMap[num] {
			// First duplicate number, so return its index
			return idx
		} else {
			// Num occured for the first time, store it in the flagMap with true value
			flagMap[num] = true
		}
	}

	// Return -1 if no duplicates found
	return -1
}
