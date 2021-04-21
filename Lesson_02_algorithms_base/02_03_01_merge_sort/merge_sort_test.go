package _2_03_01_merge_sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	A := []int{5, 2, 4, 7, 1, 3, 2, 6}
	MergeSort(A,0,len(A)-1)
	fmt.Println(A)
}

func TestCopy(t *testing.T) {
	A := []int{3, 1, 4, 5, 6}
	Ac := make([]int, len(A))
	Bc := make([]int, len(A))
	copy(Ac,A[0:3])
	copy(Bc,A[3:5])
	fmt.Println(Ac)
	fmt.Println(Bc)
}