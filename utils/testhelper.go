package utils

import "testing"

type TestCase struct {
	Name     string
	Input    []string
	Expected int
}

func RunTests(t *testing.T, fn func([]string) int, testCases []TestCase) {
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			got := fn(tc.Input)
			if got != tc.Expected {
				t.Errorf("got %d, want %d", got, tc.Expected)
			}
		})
	}
}
