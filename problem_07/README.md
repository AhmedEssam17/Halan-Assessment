# Water Jug Problem

---

## Problem Statement

You are given 2 containers: A and B. **Container A** can hold **5 litres** of water, while **Container B** can hold **3 litres**.
You are also given a source of water that you can use as you wish.
Show how you can use the containers and the water source to put exactly **4 litres** of water in **Container A**. No coding required, just write down the steps.

## Proposed Solution

### Step-by-Step Procedure

1. **Start with both containers empty**

   - A(5) = **0L**, B(3) = **0L**

2. **Fill Container B completely**

   - A(5) = **0L**, B(3) = **3L**

3. **Pour all water from Container B into Container A**

   - A(5) = **3L**, B(3) = **0L**

4. **Fill Container B completely again**

   - A(5) = **3L**, B(3) = **3L**

5. **Pour water from Container B into Container A until A is full**

   - A(5) = **5L**, B(3) = **1L**

6. **Empty Container A completely**

   - A(5) = **0L**, B(3) = **1L**

7. **Pour remaining water from Container B into Container A**

   - A(5) = **1L**, B(3) = **0L**

8. **Fill Container B completely again**

   - A(5) = **1L**, B(3) = **3L**

9. **Pour all water from Container B into Container A**
   - A(5) = **4L**, B(3) = **0L**

Now, **Container A contains exactly 4 liters**, achieving the goal.

## Summary

By strategically filling, pouring, and emptying the containers, we successfully measure exactly **4 liters** using only the 5L and 3L containers.

---
