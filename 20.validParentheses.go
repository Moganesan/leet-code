package main

func isValid(s string) bool {
	parentheses := []string{"()", "[]", "{}"}

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) {
			pair := string(s[i]) + string(s[i+1])

			for j := 0; j < len(parentheses); j++ {
				if parentheses[j] == pair {
					return true
				}
			}
		}
	}
	return false
}
