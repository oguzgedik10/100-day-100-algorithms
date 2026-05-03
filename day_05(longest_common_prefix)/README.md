# Day 05: Longest Common Prefix

## Problem Description
Write a function to find the longest common prefix string amongst an array of strings. If there is no common prefix, return an empty string "".

## Technical Analysis
The vertical scanning approach is used across all implementations. We compare the characters of all strings at the same index simultaneously. This allows us to exit early as soon as a mismatch is found or the shortest string is exhausted.

- **Time Complexity:** $O(S)$
Where $S$ is the sum of all characters in all strings. In the worst case, we compare every character.
- **Space Complexity:** $O(1)$
Ignoring the space required for the output string, we use constant extra space for pointers and counters.

## Implementation Details
- **Python:** Utilizes efficient slicing and range-based iteration.
- **Go:** Performs byte-level comparisons for high efficiency.
- **C:** Direct manipulation of null-terminated character arrays with manual memory allocation for the result string.