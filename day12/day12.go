package day12

import (
	"fmt"

	"github.com/zsommers/aoc22/point"
	"github.com/zsommers/aoc22/util"
)

func A(input []string) int {
	var start, end *point.Point
	for y, l := range input {
		for x, c := range l {
			if c == 'S' {
				start = point.New(x, y)
			}
			if c == 'E' {
				end = point.New(x, y)
			}
			if start != nil && end != nil {
				break
			}
		}
		if start != nil && end != nil {
			break
		}
	}

	return findPath(input, start, end)
}

func B(input []string) int {
	var end *point.Point
	starts := []*point.Point{}
	for y, l := range input {
		for x, c := range l {
			if c == 'S' || c == 'a' {
				starts = append(starts, point.New(x, y))
			}
			if c == 'E' {
				end = point.New(x, y)
			}
		}
	}

	min := util.MaxInt
	for i, start := range starts {
		distance := findPath(input, start, end)
		fmt.Printf("Path %3d of %3d: %3d\n", i+1, len(starts), distance)
		min = util.Min(min, distance)
	}
	return min
}

func makeVisited(topo []string) [][]bool {
	visited := [][]bool{}
	for range topo {
		visited = append(visited, make([]bool, len(topo[0])))
	}
	return visited
}

func makeDistance(topo []string) [][]int {
	distance := [][]int{}
	for range topo {
		d := []int{}
		for range topo[0] {
			d = append(d, util.MaxInt-1)
		}
		distance = append(distance, d)
	}
	return distance
}

func canVisit(current, destination *point.Point, topo []string, visited [][]bool) bool {
	// Inside bounds?
	if destination.X < 0 || destination.Y < 0 || destination.X >= len(topo[0]) || destination.Y >= len(topo) {
		return false
	}
	//
	if current.Distance(destination) > 1 {
		return false
	}
	if visited[destination.Y][destination.X] {
		return false
	}
	cVal := topo[current.Y][current.X]
	dVal := topo[destination.Y][destination.X]
	if cVal == 'S' {
		return dVal == 'a' || dVal == 'b'
	}
	if dVal == 'E' {
		return cVal == 'y' || cVal == 'z'
	}
	return dVal == cVal || dVal-1 <= cVal
}

func visit(current *point.Point, topo []string, visited [][]bool, distance [][]int) {
	// fmt.Printf("Visiting (%d, %d)<%c>", current.X, current.Y, topo[current.Y][current.X])
	neighbors := []*point.Point{}
	for _, x := range []int{current.X + 1, current.X - 1} {
		neighbors = append(neighbors, point.New(x, current.Y))
	}
	for _, y := range []int{current.Y + 1, current.Y - 1} {
		neighbors = append(neighbors, point.New(current.X, y))
	}
	for _, n := range neighbors {
		newD := distance[current.Y][current.X] + 1
		if canVisit(current, n, topo, visited) && newD < distance[n.Y][n.X] {
			// fmt.Printf(" - Setting (%d, %d)<%c> to %d", n.X, n.Y, topo[n.Y][n.X], newD)
			distance[n.Y][n.X] = newD
		}
	}
	visited[current.Y][current.X] = true
	// fmt.Println()
}

func findPath(topo []string, start, end *point.Point) int {
	visited := makeVisited(topo)
	distance := makeDistance(topo)

	distance[start.Y][start.X] = 0

	current := start
	for {
		visit(current, topo, visited, distance)
		dist := util.MaxInt
		allDone := true
		for y := range topo {
			for x := range topo[0] {
				if !visited[y][x] && distance[y][x] < dist {
					allDone = false
					current = point.New(x, y)
					dist = distance[y][x]
				}
			}
		}
		// for y := range topo {
		// 	for x := range topo[0] {
		// 		if distance[y][x] > 1000 {
		// 			fmt.Print(" X")
		// 		} else {
		// 			fmt.Printf("%2d", distance[y][x]%100)
		// 		}
		// 	}
		// 	fmt.Println()
		// }
		if allDone {
			break
		}
	}
	return distance[end.Y][end.X]
}

//
// Second implementation
//

type square struct {
	height    byte
	visited   bool
	neighbors []*square
	distance  int
}

type grid [][]*square

func buildGrid(input []string) grid {
	g := grid{}
	for x := range input[0] {
		col := []*square{}
		for y := range input {
			col = append(col, &square{
				height:    input[y][x],
				distance:  util.MaxInt - 1,
				neighbors: []*square{},
			})
		}
		g = append(g, col)
	}
	return g
}

func (g grid) setNeighbors() {
	for x := range g {
		for y := range g[0] {
			s := g[x][y]
			adj := []struct{ x, y int }{
				{x + 1, y},
				{x - 1, y},
				{x, y + 1},
				{x, y - 1},
			}
			for _, a := range adj {
				if a.x < 0 || a.x >= len(g) || a.y < 0 || a.y >= len(g[0]) {
					continue
				}
				d := g[a.x][a.y]
				switch {
				case s.height == 'S':
					if d.height == 'a' {
						s.neighbors = append(s.neighbors, d)
					}
				case d.height == 'E':
					if s.height == 'z' {
						s.neighbors = append(s.neighbors, d)
					}
				case d.height-1 <= s.height:
					s.neighbors = append(s.neighbors, d)
				}
			}
		}
	}
}

func (g grid) getNext() *square {
	minDist := util.MaxInt
	var current *square

	for x := range g {
		for y := range g[0] {
			s := g[x][y]
			if !s.visited && s.distance < minDist {
				minDist = s.distance
				current = s
			}
		}
	}
	return current
}

func (g grid) getEnds() (start, end *square) {
	for x := range g {
		for y := range g[0] {
			switch g[x][y].height {
			case 'S':
				start = g[x][y]
			case 'E':
				end = g[x][y]
			}
		}
	}
	return
}

func (s *square) visit() {
	for _, n := range s.neighbors {
		if !n.visited && s.distance+1 < n.distance {
			n.distance = s.distance + 1
		}
	}
	s.visited = true
}

func newPath(input []string) int {
	g := buildGrid(input)
	g.setNeighbors()
	start, end := g.getEnds()
	start.distance = 0
	current := start
	for current != nil {
		current.visit()
		current = g.getNext()
	}
	return end.distance
}
