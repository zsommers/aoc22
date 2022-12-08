package day7

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var rawInput = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

var input = strings.Split(rawInput, "\n")

func dirStruct() *dir {
	root := dir{
		name: "/",
		files: []*file{
			{"b.txt", 14848514},
			{"c.dat", 8504156},
		},
		size: 48381165,
	}

	a := dir{
		name: "a",
		files: []*file{
			{"f", 29116},
			{"g", 2557},
			{"h.lst", 62596},
		},
		size:   94853,
		parent: &root,
	}
	root.dirs = append(root.dirs, &a)

	e := dir{
		name:   "e",
		files:  []*file{{"i", 584}},
		size:   584,
		parent: &a,
	}
	a.dirs = append(a.dirs, &e)

	d := dir{
		name: "d",
		files: []*file{
			{"j", 4060174},
			{"d.log", 8033020},
			{"d.ext", 5626152},
			{"k", 7214296},
		},
		size:   24933642,
		parent: &root,
	}
	root.dirs = append(root.dirs, &d)

	return &root
}

func TestA(t *testing.T) {
	assert.Equal(t, 95437, A(input))
}

func TestB(t *testing.T) {
	assert.Equal(t, 24933642, B(input))
}

func clearDirSizes(d *dir) {
	d.size = 0
	for _, sd := range d.dirs {
		clearDirSizes(sd)
	}
}

func Test_setDirSizes(t *testing.T) {
	root := dirStruct()
	clearDirSizes(root)
	setDirSizes(root)

	assert.Equal(t, 48381165, root.size)

	a, ok := root.getSubDir("a")
	require.True(t, ok)
	assert.Equal(t, 94853, a.size)

	e, ok := a.getSubDir("e")
	require.True(t, ok)
	assert.Equal(t, 584, e.size)

	d, ok := root.getSubDir("d")
	require.True(t, ok)
	assert.Equal(t, 24933642, d.size)
}

func Test_parseInput(t *testing.T) {
	expected := dirStruct()
	root := parseInput(input)

	assert.True(t, root.equals(expected))
}

func Test_dir_equals(t *testing.T) {
	tests := []struct {
		name string
		a, b *dir
		want bool
	}{
		{
			"Equal",
			dirStruct(),
			dirStruct(),
			true,
		},
		{
			"One empty",
			dirStruct(),
			&dir{},
			false,
		},
		{
			"Both empty",
			&dir{},
			&dir{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.a.equals(tt.b))
		})
	}
}
