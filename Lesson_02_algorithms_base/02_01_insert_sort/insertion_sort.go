package _2_01_insert_sort


// 插入排序，遍历一个数据，发现一个数字位置不对时候，就和前面的比较，直到比前面的小
func InsertionSort (A []int) {
	for j := 1 ; j < len(A); j++ {
		key := A[j]
		i := j - 1
		for i >= 0 && A[i] > key {
			A[i+1] = A[i]
			i = i - 1
		}
		A[i + 1] = key
	}
}