# Group Anagrams

## Description

Given an array of strings `strs`, group all **anagrams** together into sublists.

An **anagram** is a word formed by rearranging the letters of another word.  
Two strings are anagrams if they contain the **exact same characters**, regardless of the order.

You may return the output in **any order**.

---

## Examples

### Example 1

Input: strs = ["act", "pots", "tops", "cat", "stop", "hat"]
Output: [["hat"], ["act", "cat"], ["stop", "pots", "tops"]]

### Example 2

Input: strs = ["x"]
Output: [["x"]]

### Example 3

Input: strs = [""]
Output: [[""]]

---

## Constraints

- `1 <= strs.length <= 1000`
- `0 <= strs[i].length <= 100`
- `strs[i]` consists of **lowercase English letters** only
