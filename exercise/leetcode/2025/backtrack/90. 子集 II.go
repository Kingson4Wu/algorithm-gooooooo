package backtrack

import "sort"

/*
*
给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的 子集（幂集）。

解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。

跟079. 子集 相比，排序和去重即可
*/
func subsetsWithDup(nums []int) [][]int {

	var ans [][]int
	var path []int
	sort.Ints(nums)

	var dfs func(cur int)
	dfs = func(cur int) {
		temp := make([]int, len(path))
		copy(temp, path)
		ans = append(ans, temp)
		for i := cur; i < len(nums); i++ {
			if i > cur && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return ans
}
