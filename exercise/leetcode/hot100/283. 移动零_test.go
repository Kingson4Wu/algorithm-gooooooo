package hot100

import (
	"fmt"
	"testing"
)

/*
*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

请注意 ，必须在不复制数组的情况下原地对数组进行操作。

示例 1:

输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]
示例 2:

输入: nums = [0]
输出: [0]
*/
func moveZeroes(nums []int) {

	if len(nums) <= 1 {
		return
	}

	index := 0
	preZero := false
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if !preZero {
				preZero = true
			}
			continue
		}
		if preZero {
			nums[i], nums[index] = nums[index], nums[i]
		}
		index++
	}
}

func TestMoveZeroes(t *testing.T) {

	arr := []int{0, 1, 0, 3, 12}
	moveZeroes(arr)
	fmt.Println(arr)

	arr2 := []int{0}
	moveZeroes(arr2)
	fmt.Println(arr2)

	arr3 := []int{1, 4, 2, 0, 4, 1, 0, 4, 5, 6, 0}
	moveZeroes(arr3)
	fmt.Println(arr3)

	arr4 := []int{0, 0, 0, 3, 4, 2, 1, 7, 0, 0, 0}
	moveZeroes(arr4)
	fmt.Println(arr4)

}
