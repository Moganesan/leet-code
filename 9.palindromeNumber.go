package main

import (
	"strconv"
)

func isPalindrome(x int) bool {
	number := strconv.Itoa(x)
	reverse := ""
	for i := len(number) - 1; i >= 0; i-- {
		reverse = reverse + string(number[i])
	}
	if number == reverse {
		return true
	} else {
		return false
	}
}
