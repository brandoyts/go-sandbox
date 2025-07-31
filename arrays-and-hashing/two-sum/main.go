package main

import "fmt"

func solution(nums []int, target int) []int {
	result := []int{}
	lookup := map[int]int{}

	if len(nums) < 2 {
		return result
	}

	for i, val := range nums {
		dif := target - val

		lookupVal, exist := lookup[dif]
		if !exist {
			lookup[val] = i
			continue
		}

		result = append(result, lookupVal, i)

	}

	return result
}

func main() {
	fmt.Println(solution([]int{3, 4, 5, 6}, 7))
	fmt.Println(solution([]int{4, 5, 6}, 10))
	fmt.Println(solution([]int{5, 5}, 10))
	fmt.Println(solution([]int{7, 2, 8, 3, 5, 12}, 8))
}
