# First Duplicate Index

---

## Problem Statement

Given an integer array of length n, find the index of the **first duplicate** element. in the array and state the runtime and space complexity of your implementation.

### Example

index_of_first_duplicate( [ 5, 17, 3, 17, 4, -1 ] )\
**returns 3**, assuming the index starts with 0

## Proposed Solution

### **Procedure:**

1. **Initialize a Hash Map**

   - Use a **map** to store encountered numbers.

2. **Iterate Through Array**

   - If a number is already in the map, return its index.
   - Otherwise, store the number in the map.

3. **Return -1 If No Duplicates Exist**

## How to Run the Code

### Run the main code with quick test inputs

```sh
go run main.go
```

### Expected Output:

```
Input array: [5 17 3 17 4 -1]
Index of first duplicate: 3

Input array: [1 2 3 4 5]
Index of first duplicate: -1

Input array: [9 7 5 7 8 9]
Index of first duplicate: 3

Input array: [1 1 2 3 4 5]
Index of first duplicate: 1

Input array: []
Index of first duplicate: -1
```

### Run the unit tests

```sh
go test -v ./duplicateIndex/
```

### Expected Output:

```
=== RUN   TestFirstDuplicateIdx
=== RUN   TestFirstDuplicateIdx/Duplicate_at_index_3
=== RUN   TestFirstDuplicateIdx/No_duplicates
=== RUN   TestFirstDuplicateIdx/Duplicate_at_index_3#01
=== RUN   TestFirstDuplicateIdx/First_element_duplicate
=== RUN   TestFirstDuplicateIdx/All_unique_elements
=== RUN   TestFirstDuplicateIdx/Duplicate_at_index_5
=== RUN   TestFirstDuplicateIdx/Single_element
=== RUN   TestFirstDuplicateIdx/Empty_array
--- PASS: TestFirstDuplicateIdx (0.00s)
    --- PASS: TestFirstDuplicateIdx/Duplicate_at_index_3 (0.00s)
    --- PASS: TestFirstDuplicateIdx/No_duplicates (0.00s)
    --- PASS: TestFirstDuplicateIdx/Duplicate_at_index_3#01 (0.00s)
    --- PASS: TestFirstDuplicateIdx/First_element_duplicate (0.00s)
    --- PASS: TestFirstDuplicateIdx/All_unique_elements (0.00s)
    --- PASS: TestFirstDuplicateIdx/Duplicate_at_index_5 (0.00s)
    --- PASS: TestFirstDuplicateIdx/Single_element (0.00s)
    --- PASS: TestFirstDuplicateIdx/Empty_array (0.00s)
PASS
ok      main/duplicateIndex     0.002s
```

---
