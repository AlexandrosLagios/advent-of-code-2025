package main

import (
	"advent-of-code-2025/utils"
	"strconv"
	"strings"
)

func Part1(lines []string) int {
	resultChan := make(chan int, len(lines))

	for _, lineStr := range lines {
		go func(r string) {
			digits := strings.Split(r, "")

			maxCharge := findMaxCharge(digits)

			resultChan <- maxCharge
			//log.Println(digits, maxCharge)
		}(lineStr)
	}
	total := 0
	for i := 0; i < len(lines); i++ {
		total += <-resultChan
	}
	return total
}

func findMaxCharge(digits []string) int {
	maxCharge := 0
	maxDigit, _ := strconv.Atoi(digits[0])

	for i := 1; i < len(digits); i++ {
		current, _ := strconv.Atoi(digits[i])

		charge, _ := strconv.Atoi(strconv.Itoa(maxDigit) + digits[i])
		if charge > maxCharge {
			maxCharge = charge
		}

		if current > maxDigit {
			maxDigit = current
		}
	}

	return maxCharge
}

func Part2(lines []string) int {
	// TODO: Implement part 2
	return 0
}

func main() {
	utils.Solution{
		Day:   3,
		Part1: Part1,
		Part2: Part2,
	}.Run()
}
