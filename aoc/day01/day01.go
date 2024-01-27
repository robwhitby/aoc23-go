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
	replacements := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	lines := []string{}
	for _, line := range input {
		for k, v := range replacements {
			line = strings.ReplaceAll(line, k, k+v+k)
		}
		lines = append(lines, line)
	}
	return Part1(lines)
}
