package _2_03_01_merge_sort

func MergeSort(A []int,p,r int) {
	if p < r {
		q := (r+p)/2
		MergeSort(A,p,q)
		MergeSort(A,q+1,r)
		merge(A,p,q,r)
	}

}

// merge 将2个已经排序的数组进行合并，就像2列扑克牌一样，对比最上面的2张排就ok
// A[p:q] 和 A[q+1:r] 都是已经排好序的
func merge(A []int, p, q, r int) {

	llen := q-p+1
	rlen := r-q
	la := make([]int, llen)
	ra := make([]int, rlen)
	//copy(la,A[p:q+1])
	//copy(ra,A[q+1:r])
	//fmt.Println(la)
	//fmt.Println(ra)

	for i := 0; i < llen; i++ {
		la[i] = A[i+p]
	}

	for i := 0; i < rlen; i++ {
		ra[i] = A[i+q+1]
	}
	li, ri := 0,0
	for i := p; i <= r; i++ {
		if li == llen {
			A[i] = ra[ri]
			ri++
		}else if ri == rlen{
			A[i] = la[li]
			li++
		}else{
			if la[li] <= ra[ri] {
				A[i] = la[li]
				li++
			}else {
				A[i] = ra[ri]
				ri++
			}
		}
	}

}