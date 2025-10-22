package _02510

import (
	"fmt"
	"testing"
)

func rob2(nums []int) int {

	// 长度为1要注意判断！！！！
	if len(nums) == 1 {
		return nums[0]
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	calculate := func(nums []int) int {
		if len(nums) == 0 {
			return 0
		}
		if len(nums) == 1 {
			return nums[0]
		}
		first, second := nums[0], max(nums[0], nums[1])
		for i := 2; i < len(nums); i++ {
			first, second = second, max(first+nums[i], second)
		}
		return second
	}
	return max(calculate(nums[1:]), calculate(nums[:len(nums)-1]))
}

/**
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
3.81
MB
击败
23.81%
*/

func TestRob2(t *testing.T) {
	fmt.Println(rob2([]int{1}))
}
