package backtrack

import (
	"fmt"
	"testing"
)

/*
*
给定一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
*/
/**
自己做的，其实子集就是组合
*/
func subsets(nums []int) [][]int {

	var ans [][]int
	var path []int

	var dfs func(cur int)
	dfs = func(cur int) {
		temp := make([]int, len(path))
		copy(temp, path)
		ans = append(ans, temp)
		for i := cur; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)

	return ans
}

func TestSubsets(t *testing.T) {
	fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsets([]int{0}))
	fmt.Println(subsets([]int{1, 2, 3, 4}))
	fmt.Println(subsets([]int{1, 2}))
	fmt.Println(subsets([]int{9, 0, 3, 5, 7}))
	//为什么是[[] [7] [5] [5 7] [3] [3 7] [3 5] [3 5 7] [0] [0 7] [0 5] [0 5 7] [0 3] [0 3 7] [0 3 5] [0 3 5 7] [9] [9 7] [9 5] [9 5 7] [9 3] [9 3 7] [9 3 5] [9 3 5 7] [9 0] [9 0 7] [9 0 5] [9 0 5 7] [9 0 3] [9 0 3 7] [9 0 3 7] [9 0 3 7 7]]
}
