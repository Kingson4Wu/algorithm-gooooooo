package _024

import (
	"fmt"
	"testing"
)

/**
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素。元素的顺序可能发生改变。然后返回 nums 中与 val 不同的元素的数量。

假设 nums 中不等于 val 的元素数量为 k，要通过此题，您需要执行以下操作：

更改 nums 数组，使 nums 的前 k 个元素包含不等于 val 的元素。nums 的其余元素和 nums 的大小并不重要。
返回 k。
*/

/**
题目的思路是很简单，但总感觉自己的实现有点复杂; 还答错过一次，没考虑数组长度为1的情况！！！边界条件要注意
*/

/**
执行用时分布
1
ms
击败
33.08%
复杂度分析
消耗内存分布
2.18
MB
击败
99.65%

*/

func removeElement(nums []int, val int) int {

	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		if nums[0] == val {
			return 0
		}
		return 1
	}

	head := 0
	tail := len(nums) - 1
	k := 0
	for head < tail {

		for tail > head {
			if nums[tail] == val {
				k++
				tail--
			} else {
				break
			}

		}
		for head < tail {
			if nums[head] == val {
				k++
				nums[head], nums[tail] = nums[tail], nums[head]
				head++
				tail--
				break
			}
			head++
		}
		if head == tail {
			if nums[head] == val {
				k++
			}
		}
	}
	return len(nums) - k
}

func TestRemoveElement(t *testing.T) {
	//答错了一次，没考虑元素长度是1的情况
	fmt.Println(removeElement([]int{1}, 1))
	fmt.Println(removeElement([]int{1, 2}, 1))
	fmt.Println(removeElement([]int{1, 2, 2, 4}, 2))
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}
