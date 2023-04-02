package hot100

import (
	"fmt"
	"testing"
)

func subsetsDfs(nums []int) (ans [][]int) {

	var set []int
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		set = append(set, nums[cur])//选
		dfs(cur + 1)
		set = set[:len(set)-1]//不选
		dfs(cur + 1)
	}
	dfs(0)
	return
}

func TestSubsetsDfs(t *testing.T) {
	fmt.Println(subsetsDfs([]int{1, 2, 3}))
}
