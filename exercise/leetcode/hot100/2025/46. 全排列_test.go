package _025

import (
	"fmt"
	"testing"
)

/*
*
自己根据回忆写出来了
*/
func permute(nums []int) [][]int {

	var ans [][]int
	var path []int
	used := make([]bool, len(nums))

	var dfs func()
	dfs = func() {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
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
