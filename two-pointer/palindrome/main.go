package main

import (
	"fmt"
	"regexp"
	"strings"
)

func solution(s string) bool {

	// convert string to lowercase
	lowercase := strings.ToLower(s)

	// remove whitespaces
	removedWhitespace := strings.ReplaceAll(lowercase, " ", "")

	// remove non alphanumeric characters
	reg := regexp.MustCompile("[^a-zA-Z0-9 ]+")
	cleanedString := reg.ReplaceAllString(removedWhitespace, "")

	characters := []rune(cleanedString)

	for index, char := range characters {

		// char = left pointer (move forward)
		// r = right pointer (move backward)
		r := (len(characters) - 1) - index

		if char != characters[r] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(solution("Was it a car or a cat I saw??"))
	fmt.Println(solution("tab a cat"))
}
