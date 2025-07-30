package backtrack

/**
给定一个无重复元素的正整数数组 candidates 和一个正整数 target ，找出 candidates 中所有可以使数字和为目标数 target 的唯一组合。

candidates 中的数字可以无限制重复被选取。如果至少一个所选数字数量不同，则两种组合是不同的。
*/
/**
自己写的，不过target < candidates[i]写反了
注意for循环只关注选！！！
*/
func combinationSum(candidates []int, target int) [][]int {

	var ans [][]int
	var path []int
	var dfs func(cur, target int)
	dfs = func(cur, target int) {
		if target == 0 {
			temp := make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
			return
		}
		for i := cur; i < len(candidates); i++ {
			if target < candidates[i] {
				continue
			}
			path = append(path, candidates[i])
			dfs(i, target-candidates[i])
			path = path[:len(path)-1]
		}
	}
	dfs(0, target)
	return ans
}
