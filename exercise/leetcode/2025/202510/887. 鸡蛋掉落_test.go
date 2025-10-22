package _02510

/*
*
定义：f(t, k) 表示在最多尝试 t 次、拥有 k 个鸡蛋的情况下，最多能测试的楼层数

转移方程：f(t, k) = 1 + f(t-1, k-1) + f(t-1, k)
最终寻找最小的 t，使得 f(t, k) >= n
*/
func superEggDrop(k int, n int) int {

	f := make([][]int, n+1)
	for t := 0; t <= n; t++ {
		f[t] = make([]int, k+1)
	}
	for t := 0; t <= n; t++ {
		f[t][1] = t
	}
	for kk := 0; kk <= k; kk++ {
		f[1][kk] = 1
	}
	ans := n
	if f[1][k] >= n {
		ans = 1
	}
	for t := 1; t <= n; t++ {
		for kk := 1; kk <= k; kk++ {
			f[t][kk] = 1 + f[t-1][kk] + f[t-1][kk-1]
		}
		if f[t][k] >= n {
			ans = t
			break
		}
	}
	return ans
}
