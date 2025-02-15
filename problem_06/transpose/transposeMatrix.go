package transpose

/*
	===========================================================================================
	Proposed Solution
	===========================================================================================
	Procedure:
		- Base case to return empty matrix if the len of matrix == 0
		- Retreive value of rows and cols from the input matrix.
		- Swap the values of rows and cols when creating the transposed matrix.
		- Iterate over the input matrix values and set the new matrix values with swapped indicies
			>> newMatrix[j][i] = matrix[i][j]
*/

func TransposeMatrix(matrix [][]int) [][]int {
	// Base case to return empty matrix if the len of matrix == 0
	if len(matrix) == 0 {
		return [][]int{}
	}

	// Retreive value of rows and cols from the input matrix.
	rows := len(matrix)
	cols := len(matrix[0])

	// Swap the values of rows and cols when creating the transposed matrix.
	newMatrix := make([][]int, cols)
	for i := range newMatrix {
		newMatrix[i] = make([]int, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// Set the new matrix values with swapped indicies
			newMatrix[j][i] = matrix[i][j]
		}
	}

	return newMatrix
}
