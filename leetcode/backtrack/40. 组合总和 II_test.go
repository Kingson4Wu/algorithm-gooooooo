package backtrack

import (
	"fmt"
	"sort"
	"testing"
)

/*
*
给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用 一次 。

注意：解集不能包含重复的组合。

示例 1:

输入: candidates = [10,1,2,7,6,1,5], target = 8,
输出:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]
示例 2:

输入: candidates = [2,5,2,1,2], target = 5,
输出:
[
[1,2,2],
[5]
]

提示:

1 <= candidates.length <= 100
1 <= candidates[i] <= 50
1 <= target <= 30
*/

/*
*
回溯
不会写

关键点
1.排序数组：先对 candidates 排序，可以方便我们去重。
2.回溯 + 剪枝：
  - 每一层循环从当前下标开始往后遍历。
  - 如果当前数字和上一个数字相同，并且在同一层递归中，跳过，避免重复组合。

3.元素只使用一次：递归时 start 要传 i+1，不能回头用前面用过的数字。

| 要点      | 内容                                            |
| ------- | --------------------------------------------- |
| 是否排序    | 必须排序，方便去重和剪枝                                  |
| 如何去重    | 同层中跳过相同值（i > start && nums[i] == nums[i-1]） |
| 元素可重复用否 | ❌ 不可重复使用，同一次回溯中每个元素只能用一次                      |
| 回溯边界条件  | `sum == target` 时收集结果                         |
*/
func combinationSum2(candidates []int, target int) [][]int {

	var ans [][]int
	var path []int

	sort.Ints(candidates) // 先排序，方便去重

	var backtrack func(start, sum int)
	backtrack = func(start, sum int) {

		if sum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
			return
		}
		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i-1] == candidates[i] {
				// 剪枝：跳过重复数字（同层不能选重复数字）
				continue
			}
			if sum+candidates[i] > target {
				// 因为排序了，后面更大，直接剪枝
				break
			}
			path = append(path, candidates[i])
			backtrack(i+1, sum+candidates[i])
			path = path[:len(path)-1]
		}
	}
	backtrack(0, 0)

	return ans
}

/**
执行用时分布
1
ms
击败
26.85%
复杂度分析
消耗内存分布
4.38
MB
击败
47.83%
*/

func TestCombinationSum2(t *testing.T) {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
