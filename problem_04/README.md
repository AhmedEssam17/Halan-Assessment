# Function Composition

Function composition allows combining two functions, `f` and `g`, into a new function `h(x) = f(g(x))`.

---

## Problem Statement

Let `f` and `g` be two one-argument functions. The composition **f after g** is defined as: **h(x) = f(g(x))**\
Define a function compose that implements composition.
For example, if inc is a function that adds 1 to its argument, and square is a function that
squares its argument, then:

- `h = compose(square, inc)` # create a new function h by composing inc and square
- `h(6)` # returns 49

## Proposed Solution

### **Procedure:**

1. **Define a function type** for mathematical operations (`Operation`), which takes an `int` and returns an `int`.
2. **Implement the `square` function** that computes the square of a number (`x²`).
3. **Implement the `inc` function** that increments a number by `1`.
4. **Implement the `compose` function**, which:
   - Takes two functions `f` and `g` as inputs.
   - Returns a new function `h(x) = f(g(x))`, meaning:
     - The inner function `g(x)` is executed first.
     - The outer function `f(g(x))` is executed next, returning the final result.
5. **How it works:**
   - Create a function `h` using `compose(square, inc)`.
   - Calling `h(6)` results in:
     - `inc(6) → 7`
     - `square(7) → 49`
   - The final output is `49`.

## How to Run the Code

### Run the main code with quick test inputs

```sh
go run main.go
```

### Run the unit tests

```sh
go test -v ./composition/
```

---
