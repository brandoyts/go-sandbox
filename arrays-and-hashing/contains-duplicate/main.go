package main

import "fmt"

func solution(nums []int) bool {
	lookup := make(map[int]struct{})

	for i := 0; i < len(nums); i++ {
		_, exist := lookup[nums[i]]
		if exist {
			return true
		}

		lookup[nums[i]] = struct{}{}
	}

	return false
}

func main() {
	fmt.Println(solution([]int{1, 2, 3, 4}))
	fmt.Println(solution([]int{1, 2, 3, 4, 4}))
	fmt.Println(solution([]int{0, 3, 2, 66, 22, 66}))
}
