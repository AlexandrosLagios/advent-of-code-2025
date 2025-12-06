package main

import (
	"advent-of-code-2025/utils"
	"sort"
	"strings"
)

func Part1(lines []string) int {
	var availableIds []int
	var ranges [][2]int

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			start := utils.ToInt(parts[0])
			end := utils.ToInt(parts[1])
			ranges = append(ranges, [2]int{start, end})
		} else {
			id := utils.ToInt(line)
			availableIds = append(availableIds, id)
		}
	}
	freshIds := 0
	mergedRanges := mergeRanges(ranges)
	for _, id := range availableIds {
		for _, mergedRange := range mergedRanges {
			if id >= mergedRange[0] && id <= mergedRange[1] {
				freshIds++
				break
			}
		}
	}
	return freshIds
}

func mergeRanges(ranges [][2]int) [][2]int {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})
	mergedRanges := make([][2]int, 0)
	for _, r := range ranges {
		if len(mergedRanges) == 0 || mergedRanges[len(mergedRanges)-1][1] < r[0]-1 {
			mergedRanges = append(mergedRanges, r)
		} else {
			mergedRanges[len(mergedRanges)-1][1] = max(mergedRanges[len(mergedRanges)-1][1], r[1])
		}
	}
	return mergedRanges
}

func Part2(lines []string) int {
	var ranges [][2]int

	for _, line := range lines {
		if len(line) == 0 {
			break
		}
		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			start := utils.ToInt(parts[0])
			end := utils.ToInt(parts[1])
			ranges = append(ranges, [2]int{start, end})
		}
	}
	freshIds := 0
	mergedRanges := mergeRanges(ranges)
	for _, mergedRange := range mergedRanges {
		freshIds += mergedRange[1] - mergedRange[0] + 1
	}
	return freshIds
}

func main() {
	utils.Solution{
		Day:   5,
		Part1: Part1,
		Part2: Part2,
	}.Run()
}
