package main

import "fmt"

func romanToInt(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		s1 := value(string(s[i]))
		if i+1 < len(s) {
			s2 := value(string(s[i+1]))

			if s1 >= s2 {
				sum = sum + s1
			} else {
				sum = sum - s1
			}
		} else {
			sum = sum + s1
		}
		fmt.Println(value(string(s[i])))
	}
	return sum
}

func value(c string) int {
	switch c {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	default:
		fmt.Println("Invlid Romen")
		return 0
	}
}
