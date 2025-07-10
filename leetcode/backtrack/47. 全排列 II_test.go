package backtrack

import (
	"fmt"
	"sort"
	"testing"
)

/*
*
给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。

示例 1：

输入：nums = [1,1,2]
输出：
[[1,1,2],

	[1,2,1],
	[2,1,1]]

示例 2：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

提示：

1 <= nums.length <= 8
-10 <= nums[i] <= 10
*/

/*
*
不会写

回溯 + 排序 + used 数组 + 去重
*/
/**
凭回忆写得，相比46.全排列I:
增加排序
增加重复剪枝

还是没写对！！！
*/
func permuteUnique(nums []int) [][]int {

	var ans [][]int
	var path []int
	used := make([]bool, len(nums))

	sort.Ints(nums)

	var dfs func()
	dfs = func() {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] && !used[i] {
				continue
			}
			/*if used[i] {
				continue
			}*/
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

func TestPermuteUnique(t *testing.T) {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
}
