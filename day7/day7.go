package day7

import (
	"fmt"
	"strings"

	"github.com/zsommers/aoc22/util"
)

type dir struct {
	name   string
	files  []*file
	dirs   []*dir
	size   int
	parent *dir
}

type file struct {
	name string
	size int
}

func (d *dir) getSubDir(s string) (*dir, bool) {
	for _, sd := range d.dirs {
		if sd.name == s {
			return sd, true
		}
	}
	return nil, false
}

func (d *dir) getFile(s string) (*file, bool) {
	for _, f := range d.files {
		if f.name == s {
			return f, true
		}
	}
	return nil, false
}

func (d1 *dir) equals(d2 *dir) bool {
	if len(d1.files) != len(d2.files) {
		return false
	}
	var f2 *file
	var ok bool
	for _, f1 := range d1.files {
		if f2, ok = d2.getFile(f1.name); !ok {
			return false
		}
		if f1.size != f2.size {
			return false
		}
	}

	if len(d1.dirs) != len(d2.dirs) {
		return false
	}
	var sd2 *dir
	for _, sd1 := range d1.dirs {
		if sd2, ok = d2.getSubDir(sd1.name); !ok {
			return false
		}
		if sd1.equals(sd2) == false {
			return false
		}
	}
	return true
}

func (d *dir) selectDirs(f func(*dir) bool) []*dir {
	selected := []*dir{}
	if f(d) {
		selected = append(selected, d)
	}
	for _, sd := range d.dirs {
		selected = append(selected, sd.selectDirs(f)...)
	}
	return selected
}

func setDirSizes(root *dir) {
	root.size = 0
	for _, d := range root.dirs {
		setDirSizes(d)
		root.size += d.size
	}
	for _, f := range root.files {
		root.size += f.size
	}
}

func parseInput(input []string) *dir {
	if input[0] != "$ cd /" {
		panic("Unknown start")
	}

	root := &dir{name: "/"}
	cwd := root
	for _, l := range input[1:] {
		ss := strings.Split(l, " ")
		switch ss[0] {
		case "$": // command
			switch ss[1] {
			case "cd":
				if ss[2] == ".." {
					cwd = cwd.parent
				} else {
					var ok bool
					if cwd, ok = cwd.getSubDir(ss[2]); !ok {
						panic(fmt.Sprintf("Count not find subdirectory '%s' in '%s'", ss[2], cwd.name))
					}
				}
			case "ls":
				// noop
			default:
				panic(fmt.Sprintf("Unknown command '%s'", ss[1]))
			}
		case "dir":
			cwd.dirs = append(cwd.dirs, &dir{
				name:   ss[1],
				parent: cwd,
			})
		default: // file
			cwd.files = append(cwd.files, &file{
				name: ss[1],
				size: util.Atoi(ss[0]),
			})
		}
	}

	setDirSizes(root)
	return root
}

func A(input []string) int {
	root := parseInput(input)
	selected := root.selectDirs(func(d *dir) bool {
		return d.size <= 100000
	})
	sum := 0
	for _, d := range selected {
		sum += d.size
	}
	return sum
}

func B(input []string) int {
	root := parseInput(input)
	neededSpace := root.size - 40000000
	selected := root.selectDirs(func(d *dir) bool {
		return d.size >= neededSpace
	})
	smallest := selected[0]
	for _, d := range selected[1:] {
		if d.size < smallest.size {
			smallest = d
		}
	}
	return smallest.size
}
