package _02510

import (
	"fmt"
	"sort"
	"testing"
)

// 凭回忆写
// 基点是第一个数，另外两个数初始为基点的右边的左右两边的指针！！！！
func threeSum(nums []int) [][]int {
	var ans [][]int
	if len(nums) < 3 {
		return ans
	}
	sort.Ints(nums)
	//以第一个为基地
	for k := 0; k < len(nums)-2; k++ {
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		left, right := k+1, len(nums)-1
		for left < right {
			v := nums[k] + nums[left] + nums[right]
			if v == 0 {
				ans = append(ans, []int{nums[k], nums[left], nums[right]})
				for left+1 < right && nums[left] == nums[left+1] {
					left++
				}
				left++
				for right-1 > left && nums[right] == nums[right-1] {
					right--
				}
				right--
			} else if v < 0 {
				for left+1 < right && nums[left] == nums[left+1] {
					left++
				}
				left++
			} else {
				for right-1 > left && nums[right] == nums[right-1] {
					right--
				}
				right--
			}
		}
	}
	return ans
}

/**
执行用时分布
37
ms
击败
21.62%
复杂度分析
消耗内存分布
8.48
MB
击败
90.54%
复杂度分析

*/

func TestThreeSum(t *testing.T) {

	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0, 1, 1}))
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4, -1, -1, -1}))
	fmt.Println(threeSum([]int{-1, 0, 1, 0}))
	fmt.Println(threeSum([]int{1, -1, -1, 0}))

}

/**
[[-1 -1 2] [-1 0 1]]
[]
[[-1 -1 2] [-1 0 1]]
[[-1 0 1]]
*/
