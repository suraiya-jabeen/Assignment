package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// CountSubstring counts the number of occurrences of substring `sub` in string `s`.
func CountSubstring(s, sub string) int {
	count := 0
	for i := strings.Index(s, sub); i != -1; i = strings.Index(s, sub) {
		count++
		s = s[i+len(sub):]
	}
	return count
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the main string: ")
	mainStr, _ := reader.ReadString('\n')
	mainStr = strings.TrimSpace(mainStr)

	fmt.Print("Enter the substring: ")
	subStr, _ := reader.ReadString('\n')
	subStr = strings.TrimSpace(subStr)

	count := CountSubstring(mainStr, subStr)
	fmt.Printf("The substring '%s' appears %d times in the string '%s'.\n", subStr, count, mainStr)
}
