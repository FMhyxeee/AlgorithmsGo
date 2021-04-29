package _5_01_cut_rod

import (
	"AlgorithmsGo/Utils"
	"math"
)

func bottomUpCutRod(p []int, n int) (q int) {
	r := make([]int, n)
	r[0] = 0
	for j := 1; j <= n; j++ {
		q = math.MinInt32
		for i := 1; i <= j; i++ {
			q = Utils.Max(q, p[i]+r[j-i])
		}
		r[j] = q
	}
	return r[n]
}
