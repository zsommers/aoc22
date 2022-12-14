package day13

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var rawInput = `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`

var input = strings.Split(rawInput, "\n")

func TestA(t *testing.T) {
	assert.Equal(t, 13, A(input))
}

func TestB(t *testing.T) {
	assert.Equal(t, 140, B(input))
}

func Test_compareLists(t *testing.T) {
	tests := []struct {
		left, right string
		want        int
	}{
		// {
		// 	left:  "1,1,3,1,1",
		// 	right: "1,1,5,1,1",
		// 	want:  1,
		// },
		{
			left:  "[1],[2,3,4]",
			right: "[1],4",
			want:  1,
		},
		{
			left:  "9",
			right: "[8,7,6]",
			want:  -1,
		},
		{
			left:  "[4,4],4,4",
			right: "[4,4],4,4,4",
			want:  1,
		},
		{
			left:  "7,7,7,7",
			right: "7,7,7",
			want:  -1,
		},
		{
			left:  "",
			right: "3",
			want:  1,
		},
		{
			left:  "[[]]",
			right: "[]",
			want:  -1,
		},
		{
			left:  "1,[2,[3,[4,[5,6,7]]]],8,9",
			right: "1,[2,[3,[4,[5,6,0]]]],8,9",
			want:  -1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s<>%s", tt.left, tt.right), func(t *testing.T) {
			assert.Equal(t, tt.want, compareLists(tt.left, tt.right))
		})
	}
}
