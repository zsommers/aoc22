package point

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoint_Distance(t *testing.T) {
	tests := []struct {
		b    Point
		want int
	}{
		{Point{0, 0}, 0},
		{Point{1, 0}, 1},
		{Point{0, 1}, 1},
		{Point{-1, 0}, 1},
		{Point{0, -1}, 1},
		{Point{1, 1}, 2},
		{Point{1, -1}, 2},
		{Point{-1, -1}, 2},
		{Point{-1, 1}, 2},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.b), func(t *testing.T) {
			assert.Equal(t, tt.want, (&Point{}).Distance(&tt.b))
		})
	}
}
