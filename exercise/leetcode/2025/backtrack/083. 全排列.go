package backtrack

import (
	"fmt"
	"testing"
)

// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

func permute(nums []int) [][]int {

	var ans [][]int
	var path []int
	used := make([]bool, len(nums))
	var dfs func()
	dfs = func() {
		if len(path) == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, path)
			ans = append(ans, temp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			path = append(path, nums[i])
			used[i] = true
			dfs()
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	dfs()
	return ans
}

func TestPermute(t *testing.T) {
	fmt.Println(permute([]int{1, 2, 3}))
}
