package day03

import (
	"github.com/robwhitby/aocgo/aoc"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type number struct {
	Value int
	Start aoc.Point
}

func (n number) points() []aoc.Point {
	ps := []aoc.Point{}
	for x := n.Start.X; x < n.Start.X+len(strconv.Itoa(n.Value)); x++ {
		ps = append(ps, aoc.Point{x, n.Start.Y})
	}
	return ps
}

func (n number) isPartNumber(grid aoc.Grid) bool {
	for _, p := range n.points() {
		for _, q := range grid.Neighbours(p) {
			if !strings.ContainsAny(grid.Get(q), "0123456789.") {
				return true
			}
		}
	}
	return false
}

func findNumbers(grid aoc.Grid) []number {
	numbers := []number{}
	re := regexp.MustCompile(`\d+`)
	for y := 0; y < len(grid); y++ {
		for _, match := range re.FindAllStringIndex(grid[y], -1) {
			v, _ := strconv.Atoi(grid[y][match[0]:match[1]])
			number := number{
				Value: v,
				Start: aoc.Point{match[0], y},
			}
			if number.isPartNumber(grid) {
				numbers = append(numbers, number)
			}
		}
	}
	return numbers
}

func Part1(input []string) int {
	grid := aoc.Grid(input)
	numbers := findNumbers(grid)
	return aoc.Reduce(numbers, func(acc int, n number) int { return acc + n.Value })
}

func Part2(input []string) int {
	grid := aoc.Grid(input)
	numbers := findNumbers(grid)
	stars := grid.Points(func(p aoc.Point) bool { return grid.Get(p) == "*" })

	answer := 0
	for _, star := range stars {
		neighbours := grid.Neighbours(star)
		nums := []number{}
		for _, n := range numbers {
			if slices.ContainsFunc(n.points(), func(p aoc.Point) bool {
				return slices.Contains(neighbours, p)
			}) {
				nums = append(nums, n)
			}
		}
		if len(nums) == 2 {
			answer += nums[0].Value * nums[1].Value
		}
	}
	return answer
}
