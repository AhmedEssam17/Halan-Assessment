# First Duplicate Index

---

## Problem Statement

Given an integer array of length n, find the index of the **first duplicate** element.

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

### Run the unit tests

```sh
go test -v ./duplicateIndex/
```

---
