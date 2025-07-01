package hot100

import (
	"fmt"
	"testing"
)

/*
*
使用堆排序，获取第k个
快排也行
两个都是试一下
*/
//快排，大的排前面
/*func findKthLargest(nums []int, k int) int {

	partition := func(nums []int, low, high int) int {
		temp := nums[low]
		for low < high {
			for low < high && nums[high] <= temp {
				high--
			}
			if low < high {
				nums[low] = nums[high]
			}
			for low < high && nums[low] > temp {
				low++
			}
			if low < high {
				nums[high] = nums[low]
			}
		}
		nums[low] = temp
		return low
	}

	var quickSortTopK func(nums []int, start, end, k int)
	quickSortTopK = func(nums []int, start, end, k int) {
		if start >= end {
			return
		}
		mid := partition(nums, start, end)
		if mid >= k-1 {
			quickSortTopK(nums, start, mid, k)
		} else {
			quickSortTopK(nums, mid+1, end, k)
		}
	}
	quickSortTopK(nums, 0, len(nums)-1, k)
	return nums[k-1]
}*/

/**
执行用时分布
1679
ms
击败
6.61%
复杂度分析
消耗内存分布
11.06
MB
击败
7.11%
*/

// 堆排序，构建最大堆
func findKthLargest(nums []int, k int) int {

	// 从上到下（递归）
	/*var adjust func(nums []int, root, length int)
	adjust = func(nums []int, root, length int) {
		//左结点
		index := root*2 + 1
		if index < length {
			if index+1 < length && nums[index+1] > nums[index] {
				index++
			}
			if nums[root] >= nums[index] {
				return
			}
			nums[root], nums[index] = nums[index], nums[root]
			adjust(nums, index, length)
		}
	}*/
	// 从上到下（递归） 非递归
	adjust := func(nums []int, root, length int) {
		//左结点
		index := root*2 + 1
		for index < length {
			if index+1 < length && nums[index+1] > nums[index] {
				index++
			}
			if nums[root] >= nums[index] {
				return
			}
			nums[root], nums[index] = nums[index], nums[root]
			root = index
			index = root*2 + 1
		}
	}

	// 初始化最大堆(从下到上)
	for i := len(nums)/2 - 1; i >= 0; i-- {
		adjust(nums, i, len(nums))
	}
	// 排序，前k个
	for i := len(nums) - 1; i >= len(nums)-k; i-- {
		//第一个和最后一个换
		nums[0], nums[i] = nums[i], nums[0]
		adjust(nums, 0, i)
	}
	return nums[len(nums)-k]
}

/**
执行用时分布
15
ms
击败
65.80%
复杂度分析
消耗内存分布
9.59
MB
击败
73.96%


*/

/**
输入
nums =
[0,1,2,3,4,5,9999]
k =
1

添加到测试用例
输出
5
预期结果
9999
*/

func TestFindKthLargest(t *testing.T) {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 1))
	fmt.Println(findKthLargest([]int{3}, 1))
	fmt.Println(findKthLargest([]int{3, 2}, 2))
	fmt.Println(findKthLargest([]int{0, 1, 2, 3, 4, 5, 9999}, 1))
}
