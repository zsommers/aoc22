package day9

import (
	"fmt"
	"strings"

	"github.com/zsommers/aoc22/set"
	"github.com/zsommers/aoc22/util"
)

type intPair struct {
	x, y int
}

func (a *intPair) add(b *intPair) intPair {
	return intPair{a.x + b.x, a.y + b.y}
}

func (a *intPair) addInPlace(b *intPair) {
	a.x += b.x
	a.y += b.y
}

func (a *intPair) distance(b *intPair) int {
	return util.Abs(a.x-b.x) + util.Abs(a.y-b.y)
}

func (a *intPair) equal(b *intPair) bool {
	return a.x == b.x && a.y == b.y
}

func unit(i int) int {
	switch {
	case i > 0:
		return 1
	case i < 0:
		return -1
	default:
		return 0
	}
}

func parseInput(input []string) []intPair {
	steps := []intPair{}
	for _, line := range input {
		ss := strings.Split(line, " ")
		i := util.Atoi(ss[1])
		var move intPair
		switch ss[0] {
		case "U":
			move = intPair{0, i}
		case "D":
			move = intPair{0, -1 * i}
		case "L":
			move = intPair{-1 * i, 0}
		case "R":
			move = intPair{i, 0}
		default:
			panic(fmt.Sprintf("'%s' not a valid direction", ss[0]))
		}
		for move.x != 0 || move.y != 0 {
			switch {
			case move.x > 0:
				steps = append(steps, intPair{1, 0})
				move.x--
			case move.x < 0:
				steps = append(steps, intPair{-1, 0})
				move.x++
			case move.y > 0:
				steps = append(steps, intPair{0, 1})
				move.y--
			case move.y < 0:
				steps = append(steps, intPair{0, -1})
				move.y++
			default:
				panic("This shouldn't happen")
			}
		}
	}
	return steps
}

func moveTail(head, tail *intPair, tailTrail *set.Set[intPair]) {
	switch head.distance(tail) {
	case 0:
		fallthrough
	case 1:
		// Close enough
	case 2:
		if head.x != tail.x && head.y != tail.y {
			// Diagonally adjacent
			break
		}
		tail.x += unit(head.x - tail.x)
		tail.y += unit(head.y - tail.y)
		tailTrail.Insert(*tail)
	case 3:
		fallthrough
	case 4:
		// Move diagonally
		tail.x += unit(head.x - tail.x)
		tail.y += unit(head.y - tail.y)
		tailTrail.Insert(*tail)
	default:
		panic(fmt.Sprintf("Bad distance from %v to %v", &head, &tail))
	}
}

func A(input []string) int {
	moves := parseInput(input)
	head := intPair{}
	tail := intPair{}
	trail := set.New[intPair]()
	trail.Insert(tail)
	for _, m := range moves {
		head.addInPlace(&m)
		moveTail(&head, &tail, trail)
	}
	return trail.Len()
}

func B(input []string) int {
	moves := parseInput(input)
	rope := []*intPair{}
	trail := []*set.Set[intPair]{}
	for i := 0; i < 10; i++ {
		rope = append(rope, &intPair{})
		trail = append(trail, set.New[intPair]())
		trail[i].Insert(intPair{})
	}
	for _, m := range moves {
		rope[0].addInPlace(&m)
		for i, r := range rope[1:] {
			moveTail(rope[i], r, trail[i+1])
		}
	}
	return trail[9].Len()
}
