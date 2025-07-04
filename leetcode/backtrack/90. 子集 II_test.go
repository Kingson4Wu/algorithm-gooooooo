package backtrack

import (
	"fmt"
	"sort"
	"testing"
)

/*
*
给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的 子集（幂集）。

解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。

示例 1：

输入：nums = [1,2,2]
输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
示例 2：

输入：nums = [0]
输出：[[],[0]]

提示：

1 <= nums.length <= 10
-10 <= nums[i] <= 10
*/
/**
写不出来
抄GPT的
*/
func subsetsWithDup(nums []int) (ans [][]int) {
	sort.Ints(nums) // 先排序以便去重
	var set []int
	var dfs func(int)
	dfs = func(cur int) {
		ans = append(ans, append([]int(nil), set...)) // 每一步都加入答案
		for i := cur; i < len(nums); i++ {
			// 如果当前元素等于前一个元素，且前一个元素已经被跳过，则跳过当前，避免重复子集
			if i > cur && nums[i] == nums[i-1] {
				continue
			}
			set = append(set, nums[i])
			dfs(i + 1)
			set = set[:len(set)-1]
		}
	}
	dfs(0)
	return
}

/**
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
4.35
MB
击败
26.09%
*/

func TestSubsetsWithDup(t *testing.T) {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}
