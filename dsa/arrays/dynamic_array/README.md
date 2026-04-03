# Dynamic Array

## Problem Description

**NeetCode URL:** [Dynamic Array Problem](https://neetcode.io/problems/dynamicArray/question)

This directory contains the solution and tests for the Dynamic Array problem.

### Problem Statement
Design a Dynamic Array (Resizable Array), similar to an `ArrayList` in Java or a `vector` in C++.

Your `DynamicArray` class should support the following operations:

1. **`DynamicArray(int capacity)`**: Initialize an empty array with a capacity of `capacity`, where `capacity > 0`.
2. **`int get(int i)`**: Return the element at index `i`. Assume that index `i` is valid.
3. **`void set(int i, int n)`**: Set the element at index `i` to `n`. Assume that index `i` is valid.
4. **`void pushback(int n)`**: Push the element `n` to the end of the array.
5. **`int popback()`**: Pop and return the element at the end of the array. Assume that the array is non-empty.
6. **`void resize()`**: Double the capacity of the array.
7. **`int getSize()`**: Return the number of elements in the array.
8. **`int getCapacity()`**: Return the capacity of the array.

If the array is full and `pushback(int n)` is called, the array should be resized first.

---

## Examples

### Example 1
**Input:**
```
["Array", 1, "getSize", "getCapacity"]
```
**Output:**
```
[null, 0, 1]
```

### Example 2
**Input:**
```
["Array", 1, "pushback", 1, "getCapacity", "pushback", 2, "getCapacity"]
```
**Output:**
```
[null, null, 1, null, 2]
```

### Example 3
**Input:**
```
["Array", 1, "getSize", "getCapacity", "pushback", 1, "getSize", "getCapacity", "pushback", 2, "getSize", "getCapacity", "get", 1, "set", 1, 3, "get", 1, "popback", "getSize", "getCapacity"]
```
**Output:**
```
[null, 0, 1, null, 1, 1, null, 2, 2, 2, null, 3, 3, 1, 2]
```

---

## Notes

- The index `i` provided to `get(int i)` and `set(int i)` is guaranteed to be greater than or equal to `0` and less than the number of elements in the array.