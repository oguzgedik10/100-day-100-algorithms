# Day 04: Merge Two Sorted Lists

## Problem Description
Merge two sorted linked lists and return it as a sorted list. The list should be made by splicing together the nodes of the first two lists.

## Technical Analysis
The problem requires an in-place merging strategy. By utilizing a dummy head node, we can simplify the edge cases and maintain a pointer to the tail of the newly merged list. Each comparison step attaches the smaller node to the tail, ensuring the final list remains sorted.

- **Time Complexity:** O(n + m)
Where n and m are the lengths of the two linked lists. We traverse each list exactly once.
- **Space Complexity:** O(1)
The merge is performed in-place by changing the next pointers of existing nodes. No additional nodes are created (excluding the dummy head).

## Implementation Details
- **Python:** Leverages object references for clean and readable pointer-like logic.
- **Go:** Uses pointers and struct types to perform efficient in-place merging.
- **C:** Implements direct pointer manipulation and struct member access, demonstrating manual memory address handling.