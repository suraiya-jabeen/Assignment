// string_functions.go
package main

import "strings"

// CountSubstring counts the number of occurrences of substring `sub` in string `s`.
func CountSubstring(s, sub string) int {
	count := 0
	for strings.Contains(s, sub) {
		count++
		s = strings.Replace(s, sub, "", 1)
	}
	return count
}
