# Halan-Assessment

## Overview

This repository contains solutions for multiple programming problems, each implemented with a structured approach. Every problem is self-contained in its respective directory and follows a consistent structure to ensure clarity, testability, and maintainability.

---

## Repository Structure

Each problem is stored in its own directory, following this pattern:

```
problem_0X
    ├── README.md
    ├── go.mod
    ├── main.go
    └── package
        ├── function.go
        └── function_test.go
```

Where:

- **`function.go`** contains the implementation of the required function with detailed comments explaining the chosen approach.
- **`function_test.go`** includes unit tests to verify that the function covers all possible cases.
- **`main.go`** provides a quick test execution of the function.
- **`README.md`** explains the problem, the approach taken, expected output, and test cases.

---

## Features

- Each problem has a standalone solution with no dependencies between problems.
- Functions are well-commented to explain the logic.
- Unit tests are provided to validate correctness across multiple test cases.
- A main function is included for quick testing.
- Each problem directory has its own README detailing:
  - The problem statement
  - The proposed solution approach
  - Expected outputs
  - Edge cases covered

---

## How to Run

To test any problem:

1. Navigate to the problem directory:

```sh
cd problem_XX
```

2. Run the main function to see a quick test:

```sh
go run main.go
```

2. Execute unit tests to check correctness:

```sh
go test -v ./package/
```
