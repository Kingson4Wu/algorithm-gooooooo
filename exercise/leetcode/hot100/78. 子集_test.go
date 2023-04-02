package hot100

import (
	"fmt"
	"testing"
)

/**
自己想的，还是拆分子问题，递归

时间
4 ms
击败
16.7%
内存
2.2 MB
击败
13.41%
*/

/*
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

示例 1：

输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
示例 2：

输入：nums = [0]
输出：[[],[0]]
*/
func subsets(nums []int) [][]int {

	if len(nums) == 0 {
		return [][]int{}
	}
	if len(nums) == 1 {
		return [][]int{{}, {nums[0]}}
	}

	sub := subsets(nums[:len(nums)-1])

	var result [][]int

	for i := 0; i < len(sub); i++ {
		result = append(result, sub[i])

		ss := make([]int, len(sub[i])+1)
		for j := 0; j < len(sub[i]); j++ {
			ss[j] = sub[i][j]
		}
		ss[len(sub[i])] = nums[len(nums)-1]

		result = append(result, ss)
	}

	return result
}

func TestSubsets(t *testing.T) {
	//fmt.Println(subsets([]int{1, 2, 3}))
	//fmt.Println(subsets([]int{0}))
	//fmt.Println(subsets([]int{1, 2, 3, 4}))
	//fmt.Println(subsets([]int{1, 2}))
	fmt.Println(subsets([]int{9, 0, 3, 5, 7}))
	//为什么是[[] [7] [5] [5 7] [3] [3 7] [3 5] [3 5 7] [0] [0 7] [0 5] [0 5 7] [0 3] [0 3 7] [0 3 5] [0 3 5 7] [9] [9 7] [9 5] [9 5 7] [9 3] [9 3 7] [9 3 5] [9 3 5 7] [9 0] [9 0 7] [9 0 5] [9 0 5 7] [9 0 3] [9 0 3 7] [9 0 3 7] [9 0 3 7 7]]
}

/**
方法一：迭代法实现子集枚举
func subsets(nums []int) (ans [][]int) {
    n := len(nums)
    for mask := 0; mask < 1<<n; mask++ {
        set := []int{}
        for i, v := range nums {
            if mask>>i&1 > 0 {
                set = append(set, v)
            }
        }
        ans = append(ans, append([]int(nil), set...))
    }
    return
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/subsets/solutions/420294/zi-ji-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

方法二：递归法实现子集枚举

func subsets(nums []int) (ans [][]int) {
    set := []int{}
    var dfs func(int)
    dfs = func(cur int) {
        if cur == len(nums) {
            ans = append(ans, append([]int(nil), set...))
            return
        }
        set = append(set, nums[cur])
        dfs(cur + 1)
        set = set[:len(set)-1]
        dfs(cur + 1)
    }
    dfs(0)
    return
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/subsets/solutions/420294/zi-ji-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
