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

### Run the unit tests

```sh
go test -v ./transpose/
```

---
