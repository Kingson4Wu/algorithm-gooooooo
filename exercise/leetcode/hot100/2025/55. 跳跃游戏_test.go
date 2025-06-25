package _025

import (
	"fmt"
	"testing"
)

/*
*
贪心算法

执行用时分布
3
ms
击败
17.22%
复杂度分析
消耗内存分布
8.76
MB
击败
43.54%
*/
func canJump(nums []int) bool {
	mostIndex := 0

	for i := 0; i < len(nums); i++ {
		//能到达
		if i <= mostIndex {
			if mostIndex < i+nums[i] {
				mostIndex = i + nums[i]
			}
			if mostIndex >= len(nums)-1 {
				return true
			}
		} else {
			//不能到达
			break
		}
	}
	return false
}

func TestCanJump(t *testing.T) {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}
