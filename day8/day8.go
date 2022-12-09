package day8

import (
	"fmt"

	"github.com/zsommers/aoc22/util"
)

func parseInput(input []string) [][]int {
	grid := make([][]int, len(input))
	for x := range grid {
		grid[x] = make([]int, len(input[0]))
	}
	for y, line := range input {
		for x, r := range line {
			grid[x][y] = int(r - '0')
		}
	}
	return grid
}

func isVisable(x, y int, grid [][]int) bool {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) {
		panic(fmt.Sprintf("(%d, %d) is out of range", x, y))
	}
	if x == 0 || y == 0 || x == len(grid)-1 || y == len(grid[0])-1 {
		return true
	}

	var east, west, north, south bool
	for curX := 0; curX < x; curX++ {
		if grid[curX][y] >= grid[x][y] {
			west = true
		}
	}
	if !west {
		return true
	}

	for curX := x + 1; curX < len(grid); curX++ {
		if grid[curX][y] >= grid[x][y] {
			east = true
		}
	}
	if !east {
		return true
	}

	for curY := 0; curY < y; curY++ {
		if grid[x][curY] >= grid[x][y] {
			north = true
		}
	}
	if !north {
		return true
	}

	for curY := y + 1; curY < len(grid[0]); curY++ {
		if grid[x][curY] >= grid[x][y] {
			south = true
		}
	}
	return !south
}

func scenicScore(x, y int, grid [][]int) int {
	score := 1

	found := false
	for curX := x - 1; curX >= 0; curX-- {
		if grid[curX][y] >= grid[x][y] {
			score *= x - curX
			found = true
			break
		}
	}
	if !found {
		score *= x
	}

	found = false
	for curX := x + 1; curX < len(grid); curX++ {
		if grid[curX][y] >= grid[x][y] {
			score *= curX - x
			found = true
			break
		}
	}
	if !found {
		score *= len(grid) - x - 1
	}

	found = false
	for curY := y - 1; curY >= 0; curY-- {
		if grid[x][curY] >= grid[x][y] {
			score *= y - curY
			found = true
			break
		}
	}
	if !found {
		score *= y
	}

	found = false
	for curY := y + 1; curY < len(grid[0]); curY++ {
		if grid[x][curY] >= grid[x][y] {
			score *= curY - y
			found = true
			break
		}
	}
	if !found {
		score *= len(grid[0]) - y - 1
	}

	return score
}

func A(input []string) int {
	grid := parseInput(input)
	sum := 2*len(grid) + 2*len(grid[0]) - 4
	for x := 1; x < len(grid)-1; x++ {
		for y := 1; y < len(grid[0])-1; y++ {
			if isVisable(x, y, grid) {
				sum++
			}
		}
	}
	return sum
}

func B(input []string) int {
	grid := parseInput(input)
	max := 0
	for x := 1; x < len(grid)-1; x++ {
		for y := 1; y < len(grid[0])-1; y++ {
			max = util.Max(max, scenicScore(x, y, grid))
		}
	}
	return max
}
