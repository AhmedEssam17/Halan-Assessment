# Palindrome Checker

A **palindrome** is a word or phrase that reads the same forward and backward.

---

## Problem Statement

A palindrome is a word or sequence of characters which reads the same backward or forward.\
For example the words: level, anna, noon, rotator are all palindromes.\
Write a function palindrom that accepts a string as an argument and returns a boolean
indicating whether the input is a palindrome or not.\
For Example:

- palindrome("anna") # returns True
- palindrome("apple") # returns False

## Proposed Solution

### **Procedure:**

1. Convert the whole string to lower case to handle Case Senstivity "Not mentioned but suggested"
2. Remove all spaces from string "Not mentioned but to handle strings like (race car) and similar ones"
3. Use 2 pointers to compare first and last chars till the pointers meet or cross return true or false accordingly.

## How to Run the Code

### Run the main code with quick test inputs

```sh
go run main.go
```

### Expected Output:

```
Testing Palindrome Function:
        isPalindrome("rotator") => true
        isPalindrome("Level") => true
        isPalindrome("race car") => true
        isPalindrome("apple") => false
        isPalindrome("Essam") => false
        isPalindrome("Anna") => true
        isPalindrome("a") => true
        isPalindrome("") => true
```

### Run the unit tests

```sh
go test ./palindrome/
```

### Expected Output:

```
=== RUN   TestIsPalindrome
=== RUN   TestIsPalindrome/rotator
=== RUN   TestIsPalindrome/Level
=== RUN   TestIsPalindrome/race_car
=== RUN   TestIsPalindrome/apple
=== RUN   TestIsPalindrome/Essam
=== RUN   TestIsPalindrome/Anna
=== RUN   TestIsPalindrome/a
=== RUN   TestIsPalindrome/#00
--- PASS: TestIsPalindrome (0.00s)
    --- PASS: TestIsPalindrome/rotator (0.00s)
    --- PASS: TestIsPalindrome/Level (0.00s)
    --- PASS: TestIsPalindrome/race_car (0.00s)
    --- PASS: TestIsPalindrome/apple (0.00s)
    --- PASS: TestIsPalindrome/Essam (0.00s)
    --- PASS: TestIsPalindrome/Anna (0.00s)
    --- PASS: TestIsPalindrome/a (0.00s)
    --- PASS: TestIsPalindrome/#00 (0.00s)
PASS
ok      main/palindrome 0.002s
```

---
