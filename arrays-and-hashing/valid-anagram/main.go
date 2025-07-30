package main

import (
	"fmt"
	"strings"
)

func solution(s string, t string) bool {
	sLower := strings.ToLower(s)
	tLower := strings.ToLower(t)

	if len(sLower) != len(tLower) {
		return false
	}

	length := len(sLower)

	// sLookup is the source of truth
	sLookup := map[rune]int{}

	tLookup := map[rune]int{}

	sRune := []rune(sLower)
	tRune := []rune(tLower)

	for i := 0; i < length; i++ {
		key := sRune[i]
		frequency, exist := sLookup[key]
		if exist {
			sLookup[key] = frequency + 1
		} else {
			sLookup[key] = 1
		}
	}

	for i := 0; i < length; i++ {
		key := tRune[i]
		frequency, exist := tLookup[key]
		if exist {
			tLookup[key] = frequency + 1
		} else {
			tLookup[key] = 1
		}
	}

	for i := 0; i < length; i++ {
		key := sRune[i]
		sFrequency, sExist := sLookup[key]
		if !sExist {
			return false
		}

		tFrequency, tExist := tLookup[key]
		if !tExist {
			return false
		}

		if sFrequency != tFrequency {
			return false
		}

	}

	return true
}

func main() {
	fmt.Println(solution("racecar", "carrace"))
	fmt.Println(solution("jam", "jar"))
	fmt.Println(solution("civic", "vicic"))
	fmt.Println(solution("civic", "cicvi"))
}
