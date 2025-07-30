package backtrack

import "sort"

// 给定一个可包含重复数字的整数集合 nums ，按任意顺序 返回它所有不重复的全排列。
func permuteUnique(nums []int) [][]int {

	var ans [][]int
	var path []int
	used := make([]bool, len(nums))

	//漏了这步先排序！！！！
	sort.Ints(nums)

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
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
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
