package hot100

import (
	"fmt"
	"testing"
)

/**
看了题解，尝试自己来

时间
0 ms
击败
100%
内存
2.5 MB
击败
76.83%
*/

/*
*
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

示例 1：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
示例 2：

输入：nums = [0,1]
输出：[[0,1],[1,0]]
示例 3：

输入：nums = [1]
输出：[[1]]

提示：

1 <= nums.length <= 6
-10 <= nums[i] <= 10
nums 中的所有整数 互不相同
*/
/*func permute(nums []int) [][]int {

	var ans [][]int
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int{}, nums...))
			return
		}
		for i := cur; i < len(nums); i++ {
			nums[cur], nums[i] = nums[i], nums[cur]
			dfs(cur + 1)
			nums[cur], nums[i] = nums[i], nums[cur]
			//dfs(cur + 1)
		}
	}
	dfs(0)
	return ans
}*/

func permute(nums []int) [][]int {
	var res [][]int
	var path []int
	used := make([]bool, len(nums))

	var dfs func()
	dfs = func() {
		if len(path) == len(nums) {
			res = append(res, append([]int(nil), path...))
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
	return res
}

func TestPermute(t *testing.T) {
	fmt.Println(permute([]int{1, 2, 3}))
}
