package _024

import (
	"fmt"
	"testing"
)

/**
题目的思路是很简单，但总感觉自己的实现有点复杂
*/

func removeElement(nums []int, val int) int {

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
	fmt.Println(removeElement([]int{1, 2, 2, 4}, 2))
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}
