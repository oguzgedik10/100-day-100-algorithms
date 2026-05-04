# Day 06: Binary Search

## Problem Description
Given an array of integers `nums` which is sorted in ascending order, and an integer `target`, write a function to search `target` in `nums`. If `target` exists, then return its index. Otherwise, return -1.

## Technical Analysis
Binary Search is a divide-and-conquer algorithm that finds the position of a target value within a sorted array. It compares the target value to the middle element of the array; if they are not equal, the half in which the target cannot lie is eliminated, and the search continues on the remaining half.

- **Time Complexity:** $O(\log n)$
The search space is halved in each iteration, leading to logarithmic time complexity.
- **Space Complexity:** $O(1)$
The iterative approach uses a constant amount of extra space.

## Implementation Details
- **Python:** Iterative implementation using dynamic range adjustment.
- **Go:** Optimized midpoint calculation to prevent potential integer overflow in large datasets.
- **C:** Efficient index manipulation focusing on low-level array access and minimal memory overhead.