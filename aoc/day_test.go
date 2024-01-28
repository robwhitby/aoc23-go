package aoc_test

import (
	"github.com/robwhitby/aocgo/aoc/day01"
	"github.com/robwhitby/aocgo/aoc/day02"
	"github.com/robwhitby/aocgo/aoc/day03"
	"os"
	"strings"
	"testing"
)

func input(name string) []string {
	bs, err := os.ReadFile("../inputs/" + name + ".txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(bs), "\n")
}

func assert(t *testing.T, actual int, expected int) {
	t.Helper()
	if actual != expected {
		t.Errorf("got %v but expected %v", actual, expected)
	}
}

func Test_Day01(t *testing.T) {
	assert(t, day01.Part1(input("day01ex")), 142)
	println(day01.Part1(input("day01")))
	assert(t, day01.Part2(input("day01ex2")), 281)
	println(day01.Part2(input("day01")))
}

func Test_Day02(t *testing.T) {
	assert(t, day02.Part1(input("day02ex")), 8)
	println(day02.Part1(input("day02")))
	assert(t, day02.Part2(input("day02ex")), 2286)
	println(day02.Part2(input("day02")))
}

func Test_Day03(t *testing.T) {
	assert(t, day03.Part1(input("day03ex")), 4361)
	println(day03.Part1(input("day03")))
	assert(t, day03.Part2(input("day03ex")), 467835)
	println(day03.Part2(input("day03")))
}
