package main

import "fmt"

func solution(nums []int, k int) []int {
	max := 0
	maxArr := make([]int, 0)
	maxPosition := 0

	if k > len(nums) {
		return maxArr
	}

	for i := 0; i < k; i++ {
		if nums[i] > max {
			max = nums[i]
			maxPosition = i
		}
	}

	maxArr = append(maxArr, max)

	fmt.Println(maxPosition)

	for i := k; i < len(nums); i++ {

		if nums[i] >= nums[maxPosition] {
			maxPosition = i
			maxArr = append(maxArr, max)
			fmt.Println("right", nums[i])
		}

	}
	fmt.Println("end of loop")

	return maxArr

}

func main() {
	fmt.Println(solution([]int{1, 3, 4, 8, 5}, 3))
}
