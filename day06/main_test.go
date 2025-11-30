package main

import (
	"testing"

	"advent-of-code-2025/utils"
)

const expected = 0

func TestPart1(t *testing.T) {
	input := utils.MustReadLines("input_test.txt")
	got := part1(input)
	if got != expected {
		t.Errorf("part1() = %d, want %d", got, expected)
	}
}
