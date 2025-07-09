package _025

import (
	"fmt"
	"testing"
)

/*
*
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
4.57
MB
击败
66.03%

自己做的，调试中改了一次
*/
func combinationSum(candidates []int, target int) [][]int {

	var ans [][]int
	var path []int

	var dfs func(start, target int)
	dfs = func(start, target int) {
		if target == 0 {
			temp := make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
			return
		}
		//for i := start; i < len(candidates); i++ {
		if start < len(candidates) {
			if target >= candidates[start] {
				path = append(path, candidates[start])
				dfs(start, target-candidates[start])
				path = path[:len(path)-1]
			}
			dfs(start+1, target)
		}

		//}
	}
	dfs(0, target)

	return ans
}

func TestCombinationSum(t *testing.T) {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
	fmt.Println(combinationSum([]int{2}, 1))
}
