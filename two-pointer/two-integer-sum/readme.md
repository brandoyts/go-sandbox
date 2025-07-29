# Two Integer Sum II

## Description

Given an array of integers `numbers` that is **sorted in non-decreasing order**,  
return the **1-indexed** positions `[index1, index2]` of the two numbers such that:

- `numbers[index1 - 1] + numbers[index2 - 1] == target`
- `index1 < index2`
- You **may not use the same element twice**

There will always be **exactly one valid solution**.

Your solution must use **O(1)** additional space.

---

## Example

### Example 1

Input: numbers = [1, 2, 3, 4], target = 3

Output: [1, 2]

Explanation:
The sum of numbers[0] + numbers[1] = 1 + 2 = 3.
Since the array is 1-indexed in the result, we return [1, 2].

---

## Constraints

- `2 <= numbers.length <= 1000`
- `-1000 <= numbers[i] <= 1000`
- `-1000 <= target <= 1000`
- Exactly one solution exists
