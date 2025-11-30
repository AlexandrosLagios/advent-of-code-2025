package utils

import "math"

// Abs returns the absolute value of an integer
// Note: math.Abs only works with float64, so we need this for int
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Sum returns the sum of all integers in a slice
func Sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// Product returns the product of all integers in a slice
func Product(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result := 1
	for _, n := range nums {
		result *= n
	}
	return result
}

// GCD returns the greatest common divisor using Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM returns the least common multiple
func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

// LCMSlice returns the LCM of multiple numbers
func LCMSlice(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		result = LCM(result, nums[i])
	}
	return result
}

// ManhattanDistance calculates Manhattan distance between two points
func ManhattanDistance(x1, y1, x2, y2 int) int {
	return Abs(x1-x2) + Abs(y1-y2)
}

// EuclideanDistance calculates Euclidean distance between two points
func EuclideanDistance(x1, y1, x2, y2 int) float64 {
	dx := float64(x1 - x2)
	dy := float64(y1 - y2)
	return math.Sqrt(dx*dx + dy*dy)
}

// Pow returns base^exp for integers using fast exponentiation
func Pow(base, exp int) int {
	result := 1
	for exp > 0 {
		if exp%2 == 1 {
			result *= base
		}
		base *= base
		exp /= 2
	}
	return result
}
