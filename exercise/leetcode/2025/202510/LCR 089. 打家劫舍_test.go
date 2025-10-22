package _02510

import (
	"fmt"
	"testing"
)

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	first, second := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		first, second = second, max(first+nums[i], second)
	}
	return second
}

/**
执行用时分布
1
ms
击败
59.26%
复杂度分析
消耗内存分布
3.71
MB
击败
81.48%

*/

func TestRob(t *testing.T) {
	fmt.Println(rob([]int{2, 1, 1, 2}))
}
