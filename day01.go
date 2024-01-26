package main

import (
	"os"
	"strconv"
	"strings"
)

func read(name string) []string {
	bs, err := os.ReadFile("inputs/" + name + ".txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(bs), "\n")
}

func main() {
	println(part1(read("day01ex")) == 142)
	println(part1(read("day01")))
	println(part2(read("day01ex2")) == 281)
	println(part2(read("day01")))
}

func part1(input []string) int {
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

func part2(input []string) int {
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
	return part1(lines)
}