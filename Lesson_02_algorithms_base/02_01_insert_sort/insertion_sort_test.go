package _2_01_insert_sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestInsertionSort(t *testing.T) {
	A := []int{3, 1, 4, 5, 6}
	InsertionSort(A)
	assert.Equal(t, A, []int {1,3,4,5,6})

}