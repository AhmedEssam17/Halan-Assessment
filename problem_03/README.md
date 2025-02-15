# Run-Length Encoding and Decoding

**Run-Length Encoding (RLE)**, a simple lossless data compression technique.

---

## Problem Statement

Given a string containing characters (a-z), implement a function **runLengthEncode** that compresses repeated ‘runs’ of the same character by storing the length of that run. Additionally, implement a function **runLengthDecode** to reverse the compression.\
The output can be anything, as long as you can recreate the input with it.\

For example:

- `runLengthEncode("aaaaaaaaaabbbaxxxxyyyzyx")` → returns `"a10b3a1x4y3z1y1x1"`
- `runLengthDecode("a10b3a1x4y3z1y1x1")` → returns `"aaaaaaaaaabbbaxxxxyyyzyx"`

## Proposed Solution

### **Procedure:**

1. **runLengthEncode (Encoding)**

   - Base case: Return an empty string if `len == 0`.
   - Check for invalid input (e.g., numbers in the original string). If found, return an empty string.
   - Store the first character in `storedChar` and start a counter at `1`.
   - Iterate through the string:
     - If the next character matches `storedChar`, increment the counter.
     - If the next character is different, concatenate `storedChar` with the counter into the encoded string, reset the counter to `1`, and update `storedChar`.
   - After the loop, append the last character and its count to the encoded string.

2. **runLengthDecode (Decoding)**
   - Base case: Return an empty string if `len == 0`.
   - Store the first character in `storedChar` and initialize an empty string for the counter.
   - Iterate through the string:
     - If the character is a letter, convert the counter string into an integer and expand the character according to its count.
     - If the character is a digit, add it to the counter string (handles multi-digit counts like `a10, b27`).
   - After the iteration, expand the last character based on its count and concatenate it into the decoded string.

## How to Run the Code

### Run the main code with quick test inputs

```sh
go run main.go
```

### Expected Output:

```
Testing runLengthEncode/Decode Functions:
        Original: aaaaaaaaaabbbaxxxxyyyzyx | Encoded: a10b3a1x4y3z1y1x1    | Decoded: aaaaaaaaaabbbaxxxxyyyzyx | Match: true
        Original: abcd                 | Encoded: a1b1c1d1             | Decoded: abcd                 | Match: true
        Original: aaa                  | Encoded: a3                   | Decoded: aaa                  | Match: true
        Original:                      | Encoded:                      | Decoded:                      | Match: true
        Original: x                    | Encoded: x1                   | Decoded: x                    | Match: true
        Original: zzzzzzzzzzz          | Encoded: z11                  | Decoded: zzzzzzzzzzz          | Match: true
```

### Run the unit tests

```sh
go test -v ./runLengthEncodeDecode/
```

### Expected Output:

```
=== RUN   TestRunLengthEncode
--- PASS: TestRunLengthEncode (0.00s)
=== RUN   TestRunLengthDecode
--- PASS: TestRunLengthDecode (0.00s)
=== RUN   TestRunLengthEncode_InvalidInput
--- PASS: TestRunLengthEncode_InvalidInput (0.00s)
=== RUN   TestRunLengthDecode_InvalidCases
--- PASS: TestRunLengthDecode_InvalidCases (0.00s)
PASS
ok      main/runLengthEncodeDecode      0.002s
```

---
