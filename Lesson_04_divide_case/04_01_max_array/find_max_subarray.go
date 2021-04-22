package _4_01_max_array

import (
	"math"
)
/**
考虑用分治的方法求最大子数组问题。
假设我们要寻找的子数组A[low:high+1]的最大子数组，使用分治技术意味着我们要将子数组划分为两个规模尽量相等的子数组，先找到数组的中央位置，比如mid，考虑求解2个子数组
最大子数组必然位于以下三种情况：
1. 全部位于A[low:mid]中
2. 跨越2个子数组
3. 全部位于A[mid+1: high+1]中
 */
func FindMaximumSubarray(A []int, low, high int) (int, int, int) {
	// 只有一个元素
	if low == high {
		return low, high, A[low]
	}
	mid := (low+high)/2
	leftLow, leftHigh, leftSum := FindMaximumSubarray(A, low, mid)
	rightLow, rightHigh, rightSum := FindMaximumSubarray(A, mid+1, high)
	crossLow, crossHigh, crossSum := findMaxCrossingSubarray(A, low, mid, high)
	if leftSum >= rightSum && leftSum >= crossSum {
		return leftLow, leftHigh, leftSum
	}else if rightSum >= leftSum && rightSum >= crossSum {
		return rightLow, rightHigh, rightSum
	}else {
		return crossLow, crossHigh, crossSum
	}


}

func findMaxCrossingSubarray(A []int, low, mid, high int) (maxLeft, maxRight, maxSum int) {
	leftSum  := math.MinInt32
	rightSum := math.MinInt32
	maxLeft  = -1
	maxRight = -1
	sum := 0

	for i := mid; i >= low; i-- {
		sum += A[i]
		if sum > leftSum {
			leftSum = sum
			maxLeft = i
		}
	}
	sum = 0
	for i := mid + 1; i <= high; i++ {
		sum += A[i]
		if sum > rightSum {
			rightSum = sum
			maxRight = i
		}
	}

	return maxLeft, maxRight, leftSum + rightSum
}


