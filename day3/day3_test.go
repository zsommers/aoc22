package day3

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var rawInput = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

var input = strings.Split(rawInput, "\n")

func TestA(t *testing.T) {
	assert.Equal(t, 157, A(input))
}

func TestB(t *testing.T) {
	assert.Equal(t, 70, B(input))
}
