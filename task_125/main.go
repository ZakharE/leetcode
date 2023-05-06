package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(isPalindrome("race a car"))
}

func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1
	for l <= r {
		for !(unicode.IsLetter(rune(s[l])) || unicode.IsDigit(rune(s[l]))) && l < r {
			l++
		}
		for !(unicode.IsLetter(rune(s[r])) || unicode.IsDigit(rune(s[r]))) && l < r {
			r--
		}
		if unicode.ToLower(rune(s[r])) != unicode.ToLower(rune(s[l])) {
			return false
		}
		l++
		r--
	}
	return true
}
