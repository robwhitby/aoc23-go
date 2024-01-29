package day04

import (
	"math"
	"slices"
	"strings"
)

func matches(card string) int {
	fields := strings.Split(strings.Replace(card, ":", "|", 1), "|")
	numbers := strings.Fields(fields[1])
	winners := strings.Fields(fields[2])
	count := 0
	for _, n := range numbers {
		if slices.Contains(winners, n) {
			count += 1
		}
	}
	return count
}

func Part1(input []string) int {
	answer := 0
	for _, card := range input {
		m := matches(card)
		if m > 0 {
			answer += int(math.Pow(2, float64(m-1)))
		}
	}
	return answer
}

func Part2(input []string) int {
	wins := map[int]int{}
	for n := 0; n < len(input); n++ {
		m := matches(input[n])
		for i := n + 1; i < min(n+m+1, len(input)); i++ {
			wins[i] += wins[n] + 1
		}
	}
	answer := len(input)
	for _, v := range wins {
		answer += v
	}
	return answer
}
