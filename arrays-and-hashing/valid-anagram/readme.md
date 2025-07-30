# Valid Anagram

## Description

Given two strings `s` and `t`, return `true` if the two strings are **anagrams** of each other, otherwise return `false`.

An **anagram** is a string formed by rearranging the letters of another string.  
Both strings must contain the **exact same characters**, with the **same frequency**, but the **order may differ**.

---

## Examples

### Example 1

Input: s = "racecar", t = "carrace"
Output: true

Explanation:
Both strings contain the same characters: r, a, c, e (in the same counts).

### Example 2

Input: s = "jar", t = "jam"
Output: false

Explanation:
String t contains an m instead of an r, so it's not an anagram.

---

## Constraints

- `s` and `t` consist of **lowercase English letters only**.
- The strings may be of different or equal length.
- To be an anagram, strings must be the **same length** and contain the **same character counts**.
