package main

import "fmt"

func solution(strs []string) [][]string {
	hs := make(map[[26]int][]string)
	result := [][]string{}

	for _, s := range strs {
		var count [26]int

		for _, c := range s {
			count[c-'a']++
		}

		hs[count] = append(hs[count], s)
	}

	for _, group := range hs {
		result = append(result, group)
	}

	return result
}

func main() {
	fmt.Println(solution([]string{"act", "pots", "tops", "cat", "stop", "hat", "civic", "vicic", "civics"}))
}
