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
Procedure:
1. Convert the whole string to lower case to handle Case Senstivity "Not mentioned but suggested"
2. Remove all spaces from string "Not mentioned but to handle strings like (race car) and similar ones"
3. Use 2 pointers to compare first and last chars till the pointers meet or cross return true or false accordingly.

## How to Run the Code

### Run the main code with quick test inputs
```sh
go run main.go
```

### Run the unit tests
```sh
go test ./palindrome/
```
---

