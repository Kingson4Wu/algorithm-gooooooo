package backtrack

import "sort"

/*
*
给定一个可能有重复数字的整数数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用一次，解集不能包含重复的组合。
*/
/**
写错两个地方：
1、❌ 去重条件错误	i > 0 会错跳过第一层不同分支	改为 i > start
2、❌ target 剪枝错误	用了 continue，浪费性能	改为 break
*/
func combinationSum2(candidates []int, target int) [][]int {

	var ans [][]int
	var path []int

	sort.Ints(candidates)

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
				//continue
				break
			}
			if i > cur && candidates[i] == candidates[i-1] {
				//if i > 0 && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			dfs(i+1, target-candidates[i])
			path = path[:len(path)-1]
		}
	}
	dfs(0, target)
	return ans
}
