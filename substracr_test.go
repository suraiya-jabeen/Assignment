package main

import "testing"

func TestSubstract(t *testing.T) {
	testCases := []struct {
		a, b     int
		expected int
	}{
		{5, 3, 2},
		{100, 39, 61},
		{-2, 50, -52},
		{2, -2, 4},
		{-2, -2, 0},
	}

	for _, tc := range testCases {
		result := Substract(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Lult(%d, %d) = %d; expected %d", tc.a, tc.b, result, tc.expected)
		}
	}

}
