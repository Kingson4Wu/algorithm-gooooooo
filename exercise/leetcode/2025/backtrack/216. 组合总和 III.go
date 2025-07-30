package backtrack

/*
*
找出所有相加之和为 n 的 k 个数的组合，且满足下列条件：

只使用数字1到9
每个数字 最多使用一次
返回 所有可能的有效组合的列表 。该列表不能包含相同的组合两次，组合可以以任何顺序返回。
*/
/**
082. 组合总和 II 的 破产版，逻辑套路一样
*/
func combinationSum3(k int, n int) [][]int {
	var ans [][]int
	var path []int

	var dfs func(cur, target int)
	dfs = func(cur, target int) {
		if target == 0 {
			if len(path) == k {
				temp := make([]int, len(path))
				copy(temp, path)
				ans = append(ans, temp)
			}
			return
		}
		for i := cur; i <= 9; i++ {
			if target < i {
				break
			}
			path = append(path, i)
			dfs(i+1, target-i)
			path = path[:len(path)-1]
		}
	}
	dfs(1, n)
	return ans
}
