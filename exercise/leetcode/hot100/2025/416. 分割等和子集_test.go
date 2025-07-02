package _025

import (
	"fmt"
	"sort"
	"testing"
)

/**
目前想法，排序后两个指针，首尾往中间移动并计算
这个算法是错的

1

2,2,1,1


*/

// 以下是错的
func canPartition(nums []int) bool {

	if len(nums) == 1 {
		return false
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	count := 0
	for i := 0; i < len(nums); i++ {
		if count >= 0 {
			count -= nums[i]
		} else {
			count += nums[i]
		}
	}
	return count == 0
}

/**
解答错误
75 / 147 个通过的测试用例

官方题解
输入
nums =
[2,2,1,1]

添加到测试用例
输出
false
预期结果
true

解答错误
74 / 147 个通过的测试用例

官方题解
输入
nums =
[1,3,4,4]

添加到测试用例
输出
true
预期结果
false
*/

func TestCanPartition(t *testing.T) {
	fmt.Println(canPartition([]int{1, 3, 4, 4}))
	fmt.Println(canPartition([]int{2, 2, 1, 1}))
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
	fmt.Println(canPartition([]int{1, 2, 3, 5}))
}
