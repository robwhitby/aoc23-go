package day04

import (
	"math"
	"slices"
	"strings"
)

func Part1(input []string) int {
	answer := 0
	for _, card := range input {
		fields := strings.Split(strings.Replace(card, ":", "|", 1), "|")
		numbers := strings.Fields(fields[1])
		winners := strings.Fields(fields[2])
		count := 0
		for _, n := range numbers {
			if slices.Contains(winners, n) {
				count += 1
			}
		}
		if count > 0 {
			answer += int(math.Pow(2, float64(count-1)))
		}
	}
	return answer
}

func Part2(input []string) int {
	return 0
}
