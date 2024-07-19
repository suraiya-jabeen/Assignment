package main

import "testing"

// TestCountSubstring tests the CountSubstring function.
func TestCountSubstring(t *testing.T) {
	tests := []struct {
		str, substr string
		expected    int
	}{
		{"hello, hello, hello", "hello", 3},
		{"hello, world", "world", 1},
		{"abcdabcdabcd", "abc", 3},
		{"aaaaaa", "aa", 3},
		{"no match here", "xyz", 0},
		{"", "empty", 0},
	}

	for _, test := range tests {
		result := CountSubstring(test.str, test.substr)
		if result != test.expected {
			t.Errorf("CountSubstring(%q, %q) = %d; want %d", test.str, test.substr, result, test.expected)
		}
	}
}
