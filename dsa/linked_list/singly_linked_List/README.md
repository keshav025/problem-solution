# Singly Linked List

## Problem Description

**NeetCode URL:** [Singly Linked List Problem](https://neetcode.io/problems/singlyLinkedList/question)

This directory contains the solution and tests for the Singly Linked List problem.

### Problem Statement
Design a Singly Linked List class.

Your `LinkedList` class should support the following operations:

1. **`LinkedList()`**: Initialize an empty linked list.
2. **`int get(int i)`**: Return the value of the `i`-th node (0-indexed). If the index is out of bounds, return `-1`.
3. **`void insertHead(int val)`**: Insert a node with `val` at the head of the list.
4. **`void insertTail(int val)`**: Insert a node with `val` at the tail of the list.
5. **`bool remove(int i)`**: Remove the `i`-th node (0-indexed). If the index is out of bounds, return `false`, otherwise return `true`.
6. **`int[] getValues()`**: Return an array of all the values in the linked list, ordered from head to tail.

---

## Examples

### Example 1
**Input:**
```
["insertHead", 1, "insertTail", 2, "insertHead", 0, "remove", 1, "getValues"]
```
**Output:**
```
[null, null, null, true, [0, 2]]
```

### Example 2
**Input:**
```
["insertHead", 1, "insertHead", 2, "get", 5]
```
**Output:**
```
[null, null, -1]
```

---

## Notes

- The index `int i` provided to `get(int i)` and `remove(int i)` is guaranteed to be greater than or equal to `0`.
