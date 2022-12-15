package point

import (
	"fmt"
	"strings"

	"github.com/zsommers/aoc22/util"
)

type Point struct {
	X, Y int
}

var Zero = &Point{0, 0}

func New(x, y int) *Point {
	return &Point{X: x, Y: y}
}

func Parse(s string) *Point {
	parts := strings.Split(strings.TrimSpace(s), ",")
	if len(parts) != 2 {
		panic(fmt.Sprintf("Expected 'x,y' but got '%s'", s))
	}
	return &Point{util.Atoi(parts[0]), util.Atoi(parts[1])}
}

func (a *Point) Equal(b *Point) bool {
	return a.X == b.X && a.Y == b.Y
}

func (a *Point) Add(b *Point) *Point {
	return &Point{a.X + b.X, a.Y + b.Y}
}

func (a *Point) AddInPlace(b *Point) {
	a.X += b.X
	a.Y += b.Y
}

func (a *Point) Subtract(b *Point) *Point {
	return &Point{a.X - b.X, a.Y - b.Y}
}

func (a *Point) SubtractInPlace(b *Point) {
	a.X -= b.X
	a.Y -= b.Y
}

func (a *Point) Unit() *Point {
	unit := func(i int) int {
		switch {
		case i > 0:
			return 1
		case i < 0:
			return -1
		default:
			return 0
		}
	}
	return &Point{unit(a.X), unit(a.Y)}
}

func (a *Point) Distance(b *Point) int {
	return util.Abs(a.X-b.X) + util.Abs(a.Y-b.Y)
}
