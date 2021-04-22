package _4_01_max_array

import (
	"fmt"
	"testing"
)

func TestFindMaxSubarray2(t *testing.T) {
	A := []int{1, 3, 4, 5, 6, 2, 5, 7}
	maxValue := FindMaxSubarray2(A)
	fmt.Println(maxValue)
}
