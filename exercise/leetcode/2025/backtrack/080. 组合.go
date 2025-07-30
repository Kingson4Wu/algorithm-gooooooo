package backtrack

//给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

func combine(n int, k int) [][]int {

	var ans [][]int
	var path []int

	var dfs func(start int)
	dfs = func(start int) {
		if len(path) == k {
			temp := make([]int, k)
			copy(temp, path)
			ans = append(ans, temp)
			return
		}
		if start > n {
			return
		}
		/*for i := start; i <= n; i++ {
			path = append(path, i)
			dfs(start + 1)
			path = path[:len(path)-1]
		}*/
		path = append(path, start)
		dfs(start + 1)
		path = path[:len(path)-1]
		dfs(start + 1)

	}
	dfs(1)
	return ans
}
