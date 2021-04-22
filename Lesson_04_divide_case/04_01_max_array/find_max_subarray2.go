package _4_01_max_array

import "math"

/**
使用如下思想设计一个非递归、线性时间的算法。
数组从左边开始，由左到右进行处理，记录到现在位置最大子数组。
若已知[0:j]为最大数组，基于如下性质将解扩展到[0:j+1]为最大数组。
A[0:j+1]的最大子数组要不是A[0:j]的最大子数组，要么是某个子数组[i:j+1] 在已知A[0:j]的最大子数组情况下，可以在线性时间内找到A[i:j+1]的最大子数组

首先 最大子数组是在数组中连续的一串数字，有2个边界 i j sum(A[i:j])
2个边界必定都为正数
寻找股票最低点的位置，然后在最低点买入再寻找一个和最低点插值最大的点卖出就行了
*/

func FindMaxSubarray2(A []int) int {
	minprice := math.MaxInt32
	maxprofit := 0
	for i := 0; i < len(A); i++ {
		if A[i] < minprice {
			minprice = A[i]
		} else if (A[i] - minprice) > maxprofit {
			maxprofit = A[i] - minprice
		}
	}
	return maxprofit
}
