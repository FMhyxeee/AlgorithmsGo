package _5_01_cut_rod

import (
	"AlgorithmsGo/Utils"
	"math"
)

func CatRod(p []int, n int) (q int) {
	if n == 0 {
		return 0
	}
	q = math.MinInt32
	for i := 1; i <= n; i++ {
		q = Utils.Max(q, p[i]+CatRod(p, n-i))

	}

	return q
}
