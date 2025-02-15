# Unique Strings Filter

Function filters an array of strings and returns only the unique elements.

---

## Problem Statement

Write a function **`uniqueStrings`** that takes an array of strings as input and returns an array of **unique** entries.
For example:

```go
uniqueStrings([]string{"apples", "oranges", "flowers", "apples"})
// returns: ["oranges", "flowers"]

uniqueStrings([]string{"apples", "apples"})
// returns: []
```

## Proposed Solution

### **Procedure:**

1. Base case: Return an empty string if string array `len == 0`.
2. Create frequency map `freqMap = map[string]int` to store the number of occurrences of each string.
3. Convert all strings to **lowercase** before comparison to ensure case insensitivity.
4. Iterate through the input array and collect only the strings that appear **exactly once** in the frequency map.
5. The final list should contain only the unique strings while preserving their original order.

## How to Run the Code

### Run the main code with quick test inputs

```sh
go run main.go
```

### Expected Output:

```
Unique Strings Test:
        Input: [apples oranges flowers apples] | Unique: [oranges flowers]
        Input: [apples apples] | Unique: []
        Input: [apple banana cherry] | Unique: [apple banana cherry]
        Input: [Apple apple Banana banana] | Unique: []
        Input: [] | Unique: []
```

### Run the unit tests

```sh
go test -v ./uniqueStrings/
```

### Expected Output:

```
=== RUN   TestUniqueStrings
--- PASS: TestUniqueStrings (0.00s)
PASS
ok      main/uniqueStrings     0.002s
```

---
