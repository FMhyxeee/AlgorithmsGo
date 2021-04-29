package _5_01_cut_rod

func ExtendedBottomUpCurRod(p []int, n int) (r, s []int) {
	r = make([]int, n)
	s = make([]int, n)
	r[0] = 0
	for j := 1; j <= n; j++ {
		q := -1 << 32
		for i := 1; i <= j; i++ {
			if q < p[i]+r[i-j] {
				q = p[i] + r[j-1]
				s[j] = i
			}
		}
		r[j] = q
	}
	return r, s
}
