package day4

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var rawInput = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

var input = strings.Split(rawInput, "\n")

func TestA(t *testing.T) {
	assert.Equal(t, 2, A(input))
}

func TestB(t *testing.T) {
	assert.Equal(t, 4, B(input))
}
