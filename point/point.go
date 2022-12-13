package point

import "github.com/zsommers/aoc22/util"

type Point struct {
	X, Y int
}

func New(x, y int) *Point {
	return &Point{X: x, Y: y}
}

func (a *Point) Equal(b *Point) bool {
	return a.X == b.X && a.Y == b.Y
}

func (a *Point) Add(b *Point) Point {
	return Point{a.X + b.X, a.Y + b.Y}
}

func (a *Point) AddInPlace(b *Point) {
	a.X += b.X
	a.Y += b.Y
}

func (a *Point) Distance(b *Point) int {
	return util.Abs(a.X-b.X) + util.Abs(a.Y-b.Y)
}
