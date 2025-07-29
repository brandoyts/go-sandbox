package main

import (
	"fmt"
	"slices"
)

func solution(nums []int, target int) []int {
	result := []int{}

	if len(nums) < 2 {
		return result
	}

	slices.Sort(nums)

	left := 0
	right := len(nums) - 1

	for left < right {
		sum := nums[left] + nums[right]

		if sum > target {
			right -= 1
		} else if sum < target {
			left += 1
		}

		sum = nums[left] + nums[right]

		if sum == target {

			result = append(result, left+1, right+1)
			break
		}
	}

	return result
}

func main() {
	fmt.Println(solution([]int{1, 2, 3, 4}, 3))
	fmt.Println(solution([]int{1, 1, 3, 5, 7, 8}, 12))
	fmt.Println(solution([]int{1, 1, 3, 5, 7, 8}, 8))
}
