# Matrix Transpose

---

## Problem Statement

In **linear algebra**, the **transpose** of a matrix is a new matrix created by swapping the rows and columns of the original matrix. Given a matrix, we aim to write a function that transposes it.

### Example

For a **2×2 matrix**:\
 | a11 a12 | >> | a11 b21 |\
 | b21 b22 | >> | a12 b22 |

## Proposed Solution

### **Procedure:**

- Base case to return empty matrix if the matrix `len == 0`.
- Retrieve the number of **rows** and **columns** from the input matrix.
- Initialize a new matrix with **swapped dimensions** (columns become rows, and rows become columns).
- Iterate through the input matrix and set the transposed matrix values by swapping indices:
  ```go
  newMatrix[j][i] = matrix[i][j]
  ```

## How to Run the Code

### Run the main code with quick test inputs

```sh
go run main.go
```

### Expected Output:

```
Original matrix: [[1 2] [3 4]]
Transposed matrix: [[1 3] [2 4]]

Original matrix: [[1 2 3 4] [5 6 7 8]]
Transposed matrix: [[1 5] [2 6] [3 7] [4 8]]

Original matrix: [[1 2 3 4]]
Transposed matrix: [[1] [2] [3] [4]]

Original matrix: [[1] [2] [3] [4]]
Transposed matrix: [[1 2 3 4]]

Original matrix: []
Transposed matrix: []
```

### Run the unit tests

```sh
go test -v ./transpose/
```

### Expected Output:

```
=== RUN   TestTransposeMatrix
=== RUN   TestTransposeMatrix/2x2_Matrix
=== RUN   TestTransposeMatrix/2x4_Matrix
=== RUN   TestTransposeMatrix/1x4_Matrix
=== RUN   TestTransposeMatrix/4x1_Matrix
=== RUN   TestTransposeMatrix/Empty_Matrix
--- PASS: TestTransposeMatrix (0.00s)
    --- PASS: TestTransposeMatrix/2x2_Matrix (0.00s)
    --- PASS: TestTransposeMatrix/2x4_Matrix (0.00s)
    --- PASS: TestTransposeMatrix/1x4_Matrix (0.00s)
    --- PASS: TestTransposeMatrix/4x1_Matrix (0.00s)
    --- PASS: TestTransposeMatrix/Empty_Matrix (0.00s)
PASS
ok      main/transpose  0.002s
```

---
