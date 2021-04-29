package _5_01_cut_rod

import (
	"AlgorithmsGo/Utils"
	"math"
)

func MemoizedCatRod(p []int, n int) (q int) {
	r := make([]int, n)
	for i := 0; i <= n; i++ {
		r[i] = math.MinInt32
	}

	return memoizedCutRodAux(p, n, r)
}

func memoizedCutRodAux(p []int, n int, r []int) (q int) {
	if r[n] >= 0 {
		return r[n]
	}
	if n == 0 {
		q = 0
	}
	if q == math.MinInt32 {
		for i := 1; i <= n; i++ {
			q = Utils.Max(q, p[i]+memoizedCutRodAux(p, n-i, r))
		}
	}
	r[n] = q
	return q
}
