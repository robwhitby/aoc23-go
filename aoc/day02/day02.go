package day02

import (
	"github.com/robwhitby/aocgo/aoc"
	"regexp"
	"strconv"
)

type game struct {
	Id, R, G, B int
}

func parse(input []string) []game {
	games := []game{}
	for idx, line := range input {
		games = append(games, parseLine(idx, line))
	}
	return games
}

func parseLine(idx int, line string) game {
	maxColour := func(colour string) int {
		gs := regexp.MustCompile(`(\d+) `+colour).FindAllStringSubmatch(line, -1)
		return aoc.Reduce(gs, func(acc int, match []string) int {
			i, _ := strconv.Atoi(match[1])
			return max(acc, i)
		})
	}
	return game{
		Id: idx + 1,
		R:  maxColour("red"),
		G:  maxColour("green"),
		B:  maxColour("blue"),
	}
}

func Part1(input []string) int {
	games := parse(input)
	validGames := aoc.Filter(games, func(g game) bool {
		return g.R <= 12 && g.G <= 13 && g.B <= 14
	})
	return aoc.Reduce(validGames, func(acc int, g game) int { return acc + g.Id })
}

func Part2(input []string) int {
	games := parse(input)
	return aoc.Reduce(games, func(acc int, g game) int {
		return acc + (g.R * g.G * g.B)
	})
}
