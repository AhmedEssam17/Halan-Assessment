# File Renaming in Linux

---

## Problem Statement

Write the Linux command needed to change a file name from original.txt to changed.txt

## Proposed Solution

### **Procedure:**

So the Linux command to rename a file called original.txt to changed.txt is as follows:

```sh
mv original.txt changed.txt
```

This simple Go code that runs as follows:

1. Delete changed.txt if exists from previous runs.
2. Create the "original.txt" file.
3. Wait 3 seconds to see the file clearly in the file explorer.
4. Proceed to run the "mv" Linux command to rename it to changed.txt.

## How to Run the Code

### Run the main code to test file renaming

```sh
go run main.go
```

### Expected Output:

Observe the file explorer to see the original.txt file being renamed to changed.txt

```
File renamed successfully using Linux command!
```

---
