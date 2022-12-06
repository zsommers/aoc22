package day2

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var rawInput = `A Y
B X
C Z`

var input = strings.Split(rawInput, "\n")

func TestA(t *testing.T) {
	assert.Equal(t, 15, A(input))
}

func TestB(t *testing.T) {
	assert.Equal(t, 12, B(input))
}
