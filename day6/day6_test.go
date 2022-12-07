package day6

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var rawInput = `mjqjpqmgbljsphdztnvjfqwrcgsmlb
bvwbjplbgvbhsrlpgdmjqwftvncz
nppdvjthqldpwncqszvftbrmjlhg
nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg
zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`

var input = strings.Split(rawInput, "\n")

func TestA(t *testing.T) {
	answers := []int{7, 5, 6, 10, 11}
	type test = struct {
		input []string
		want  int
	}
	tests := []test{}
	for i, s := range input {
		tests = append(tests, test{[]string{s}, answers[i]})
	}
	for _, tt := range tests {
		t.Run(tt.input[0], func(t *testing.T) {
			var result int
			require.NotPanics(t, func() { result = A(tt.input) })
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestB(t *testing.T) {
	answers := []int{19, 23, 23, 29, 26}
	type test = struct {
		input []string
		want  int
	}
	tests := []test{}
	for i, s := range input {
		tests = append(tests, test{[]string{s}, answers[i]})
	}
	for _, tt := range tests {
		t.Run(tt.input[0], func(t *testing.T) {
			var result int
			require.NotPanics(t, func() { result = A(tt.input) })
			assert.Equal(t, tt.want, result)
		})
	}
}
