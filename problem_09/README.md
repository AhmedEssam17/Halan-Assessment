# Tree Sum Calculation

---

## Problem Statement

Given a **tree structure** where each node contains an integer value and a list of child nodes, implement a function that calculates the **sum of all values** in the tree.

### Example

Consider the following tree:

```
    10
   /  \
  5    3
 / \    \
2   1    4
```

## Proposed Solution

### **Procedure:**

1. **Define a Node Structure**

   - Each node contains:
     - An **integer value**.
     - A **list of child nodes**.

2. **Recursive Sum Function**
   - Base case: If the node is `nil`, return `0`.
   - Initialize `total` with the value of the current node.
   - Recursively call `Sum` for each child node and add their values.

## How to Run the Code

### Run the main code with quick test inputs

```sh
go run main.go
```

### Run the unit tests

```sh
go test -v ./tree/
```

---
