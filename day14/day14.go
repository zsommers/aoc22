package day14

import (
	"strings"

	"github.com/zsommers/aoc22/point"
	"github.com/zsommers/aoc22/util"
)

type grid map[point.Point]byte

var (
	sand byte = 'o'
	rock byte = '#'
)

func A(input []string) int {
	count := 0
	g := parseInput(input)
	for dropSand(g) {
		count++
	}
	return count
}

func B(input []string) int {
	count := 0
	g := parseInput(input)
	for ok := false; !ok; _, ok = g[point.Point{X: 500, Y: 0}] {
		dropSand(g)
		count++
	}
	return count
}

func (g grid) getBounds(xMin, xMax, yMin, yMax int) {
	xMin = util.MaxInt
	yMin = util.MaxInt
	for p := range g {
		xMin = util.Min(xMin, p.X)
		xMax = util.Max(xMax, p.X)
		yMin = util.Min(yMin, p.Y)
		yMax = util.Max(yMax, p.Y)
	}
	return
}

func (g grid) floor() (yMax int) {
	for p, v := range g {
		if v == rock {
			yMax = util.Max(yMax, p.Y)
		}
	}
	yMax += 2
	return
}

// returns whether sand is at rest
func dropSand(g grid) bool {
	moves := []*point.Point{
		point.New(0, 1),
		point.New(-1, 1),
		point.New(1, 1),
	}
	loc := point.New(500, 0)
	floor := g.floor()
	for {
		if loc.Y >= floor-1 {
			g[*loc] = sand
			return false
		}
		moved := false
		for _, move := range moves {
			l := loc.Add(move)
			if _, ok := g[*l]; !ok {
				loc = l
				moved = true
				break
			}
		}
		if !moved {
			g[*loc] = sand
			return true
		}
	}
}

func parseInput(input []string) grid {
	m := grid{}

	for _, l := range input {
		points := strings.Split(l, " -> ")
		cur := point.Parse(points[0])
		m[*cur] = rock
		for _, n := range points[1:] {
			next := point.Parse(n)
			diff := cur.Subtract(next)
			for !diff.Equal(point.Zero) {
				cur.SubtractInPlace(diff.Unit())
				m[*cur] = rock
				diff = cur.Subtract(next)
			}
		}
	}

	return m
}
