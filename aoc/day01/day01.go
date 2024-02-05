package day01

import (
	"strconv"
	"strings"
)

func Part1(input []string) int {
	total := 0
	digits := "0123456789"

	for _, line := range input {
		first := strings.IndexAny(line, digits)
		last := strings.LastIndexAny(line, digits)
		num, _ := strconv.Atoi(string(line[first]) + string(line[last]))
		total += num
	}
	return total
}

func Part2(input []string) int {
	nums := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	lines := []string{}
	for _, line := range input {
		for i, n := range nums {
			line = strings.ReplaceAll(line, n, n+strconv.Itoa(i+1)+n)
		}
		lines = append(lines, line)
	}
	return Part1(lines)
}
