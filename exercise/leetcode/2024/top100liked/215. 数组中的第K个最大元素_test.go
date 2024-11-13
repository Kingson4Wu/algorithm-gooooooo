package top100liked

import (
	"fmt"
	"testing"
)

/**
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。



示例 1:

输入: [3,2,1,5,6,4], k = 2
输出: 5
示例 2:

输入: [3,2,3,1,2,4,5,5,6], k = 4
输出: 4


提示：

1 <= k <= nums.length <= 105
-104 <= nums[i] <= 104
*/

/*
*
学习堆排序完第二天凭记忆完成。。。

执行用时分布
25
ms
击败
42.28%
复杂度分析
消耗内存分布
9.01
MB
击败
32.59%
*/
func findKthLargest(nums []int, k int) int {

	adjust := func(nums []int, root int, length int) {
		child := root*2 + 1
		for child < length {
			rightChild := child + 1
			if rightChild < length && nums[rightChild] > nums[child] {
				child = rightChild
			}
			if nums[child] > nums[root] {
				nums[child], nums[root] = nums[root], nums[child]
			}
			root = child
			child = root*2 + 1
		}
	}

	for i := len(nums)/2 - 1; i >= 0; i-- {
		adjust(nums, i, len(nums))
	}
	for i := 0; i < k; i++ {
		nums[len(nums)-1-i], nums[0] = nums[0], nums[len(nums)-1-i]
		//注意这里最大值拿出来到最后，然后长度就要减1了
		adjust(nums, 0, len(nums)-1-i)
	}
	return nums[len(nums)-k]
}

func TestFindKthLargest(t *testing.T) {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 1))
	fmt.Println(findKthLargest([]int{3}, 1))
	fmt.Println(findKthLargest([]int{3, 2}, 2))
}

/**
调整方法（寻找左右节点最大的那个跟根节点交换），自上而下

1.初始化堆
从最后一个根节点开始遍历，自下而上调调整方法
2.遍历排序（TopK）
第一个根节点开始遍历到根节点，和最后一个节点交换，再调调整方法
*/
