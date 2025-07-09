package hot100

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

/**
自己总结出来的规律

1、从后面开始找，开始降序的第一个元素，然后找到比这个元素大的最小元素，进行交换，然后对后面的序列升序排序
2、本来是逆序的，直接一步升序处理

时间
4 ms
击败
50.47%
内存
2.3 MB
击败
51.83%

*/

/*
*
整数数组的一个 排列  就是将其所有成员以序列或线性顺序排列。

例如，arr = [1,2,3] ，以下这些都可以视作 arr 的排列：[1,2,3]、[1,3,2]、[3,1,2]、[2,3,1] 。
整数数组的 下一个排列 是指其整数的下一个字典序更大的排列。更正式地，如果数组的所有排列根据其字典顺序从小到大排列在一个容器中，那么数组的 下一个排列 就是在这个有序容器中排在它后面的那个排列。如果不存在下一个更大的排列，那么这个数组必须重排为字典序最小的排列（即，其元素按升序排列）。

例如，arr = [1,2,3] 的下一个排列是 [1,3,2] 。
类似地，arr = [2,3,1] 的下一个排列是 [3,1,2] 。
而 arr = [3,2,1] 的下一个排列是 [1,2,3] ，因为 [3,2,1] 不存在一个字典序更大的排列。
给你一个整数数组 nums ，找出 nums 的下一个排列。

必须 原地 修改，只允许使用额外常数空间。

示例 1：

输入：nums = [1,2,3]
输出：[1,3,2]
示例 2：

输入：nums = [3,2,1]
输出：[1,2,3]
示例 3：

输入：nums = [1,1,5]
输出：[1,5,1]

提示：

1 <= nums.length <= 100
0 <= nums[i] <= 100
*/
func nextPermutation(nums []int) {

	if len(nums) <= 1 {
		return
	}

	index := -1
	/** 找到逆序的前一个元素下标 */
	for i := len(nums) - 2; i >= 0; i-- {

		if nums[i] >= nums[i+1] {
			continue
		}
		index = i
		break
	}

	if index > -1 {
		/** 找到比index大的最小元素，进行交换，并对剩下的升序 */
		swapIndex := -1
		swapValue := math.MaxInt32

		for i := index + 1; i < len(nums); i++ {
			if nums[i] > nums[index] && nums[i] < swapValue {
				swapValue = nums[i]
				swapIndex = i
			}
		}
		nums[index], nums[swapIndex] = nums[swapIndex], nums[index]

		sort.Ints(nums[index+1:])
	} else {
		/** 本来全部降序的，下一个就是全部升序 */
		sort.Ints(nums)
	}

}

func TestNextPermutation(t *testing.T) {
	/*a := []int{1, 2, 3}
	nextPermutation(a)
	fmt.Println(a)
	b := []int{3, 2, 1}
	nextPermutation(b)
	fmt.Println(b)
	c := []int{1, 1, 5}
	nextPermutation(c)
	fmt.Println(c)

	d := []int{1, 2, 3, 4}
	for i := 0; i < 20; i++ {
		nextPermutation(d)
		fmt.Println(d)
	}*/

	fmt.Println("======")
	aa := []int{1, 3, 4, 2}
	nextPermutation(aa)
	fmt.Println(aa)

	/*fmt.Println("======")
	f := []int{1, 1, 3, 5}
	for i := 0; i < 20; i++ {
		nextPermutation(f)
		fmt.Println(f)
	}*/

}
