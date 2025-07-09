package _025

import (
	"fmt"
	"testing"
)

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
		for i := start; i < len(candidates); i++ {
			if target >= candidates[i] {
				path = append(path, candidates[i])
				dfs(i, target-candidates[i])
				path = path[:len(path)-1]
			}
			dfs(i+1, target)
		}
	}
	dfs(0, target)

	return ans
}

func TestCombinationSum(t *testing.T) {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
	fmt.Println(combinationSum([]int{2}, 1))
}
