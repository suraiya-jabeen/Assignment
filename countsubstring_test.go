// string_functions_test.go
package main

import "testing"

func TestCountSubstring(t *testing.T) {
	testCases := []struct {
		input    string
		sub      string
		expected int
	}{
		{"Hello, hello, hello, how low?", "hello", 3}, // Case-sensitive count
		{"aaa", "aa", 2},       // Overlapping substrings
		{"", "foo", 0},         // Empty input string
		{"hello", "world", 0},  // Substring not found
		{"aaaaa", "aa", 2},     // Non-overlapping substrings
		{"banana", "ana", 2},   // Edge case: overlapping start and end
		{"banana", "anana", 1}, // Edge case: whole string as substring
	}

	for _, tc := range testCases {
		actual := CountSubstring(tc.input, tc.sub)
		if actual != tc.expected {
			t.Errorf("CountSubstring(%q, %q) = %d; want %d", tc.input, tc.sub, actual, tc.expected)
		}
	}
}
