package LCR

/*
*
根据回忆写完

执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
5.98
MB
击败
31.43%
*/
func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1

	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}

	index2, index3, index5 := 0, 0, 0

	for i := 1; i < n; i++ {
		r2, r3, r5 := dp[index2]*2, dp[index3]*3, dp[index5]*5
		v := min(r2, min(r3, r5))
		if v == r2 {
			index2++
		}
		if v == r3 {
			index3++
		}
		if v == r5 {
			index5++
		}
		dp[i] = v
	}
	return dp[n-1]
}
