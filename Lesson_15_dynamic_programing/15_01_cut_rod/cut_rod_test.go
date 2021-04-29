package _5_01_cut_rod

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCatRod(t *testing.T) {
	p := []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	assert.Equal(t, CatRod(p, 4), 10)
}
