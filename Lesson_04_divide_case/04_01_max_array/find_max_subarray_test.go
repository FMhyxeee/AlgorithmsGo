package _4_01_max_array

import (
	"fmt"
	"testing"
)

func TestFindMaximumSubarray(t *testing.T) {
	A := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	low, high, sum := FindMaximumSubarray(A, 0, len(A)-1)
	fmt.Println(low)
	fmt.Println(high)
	fmt.Println(sum)
}
