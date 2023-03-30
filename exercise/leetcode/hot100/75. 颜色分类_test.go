package hot100

import (
	"algorithm-gooooooo/sort"
	"fmt"
	"testing"
)

/**
个人想法：
三个数字，刚好可以使用三指针
一个指针p，left，right； p指针遍历for p<=right; left和right分别指向非0和非2的数组下标
*/

/*
*
给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

必须在不使用库内置的 sort 函数的情况下解决这个问题。

示例 1：

输入：nums = [2,0,2,1,1,0]
输出：[0,0,1,1,2,2]
示例 2：

输入：nums = [2,0,1]
输出：[0,1,2]

提示：

n == nums.length
1 <= n <= 300
nums[i] 为 0、1 或 2

进阶：

你能想出一个仅使用常数空间的一趟扫描算法吗？
*/
func sortColors(nums []int) {

	if len(nums) <= 1 {
		return
	}

	left, p, right := 0, 0, len(nums)-1

	for p <= right {
		if nums[p] == 1 {
			p++
			continue
		}
		if nums[p] == 0 {
			nums[p], nums[left] = nums[left], nums[p]
			left++

			for left < right && nums[left] == 0 {
				left++
			}
			if p < left {
				p = left
			}
			continue
		}
		if nums[p] == 2 {
			nums[p], nums[right] = nums[right], nums[p]
			right--
			for right > left && nums[right] == 2 {
				right--
			}
			continue
		}
	}
}

func TestSortColors(t *testing.T) {
	verify([]int{1, 1, 0, 2, 2, 0, 0, 1, 1, 1, 0, 1, 2}, t)
	verify([]int{0, 0, 0, 1, 1, 1, 1, 2, 2, 2, 1, 0, 0, 0, 1, 2, 2, 2, 1, 1, 1, 0, 0, 0, 1, 0, 2, 2, 2, 2}, t)
	verify([]int{0, 0, 0, 1, 1, 1, 1}, t)
	verify([]int{1, 1, 1, 1, 0, 0, 1, 1, 1, 0}, t)
	verify([]int{2, 2, 2, 1, 1, 1, 1, 2, 2, 2, 1}, t)
	verify([]int{1, 1, 1, 1, 1, 1}, t)
	verify([]int{0, 0, 0, 0, 0, 0}, t)
	verify([]int{2, 1, 0}, t)
	verify([]int{2, 2, 1, 0}, t)

	for i := 0; i < 30; i++ {
		arr := []int{2, 1, 0}
		sort.Shuffle(arr)
		verify(arr, t)
	}
	for i := 0; i < 30; i++ {
		arr := []int{2, 2, 2, 1, 1, 0, 0}
		sort.Shuffle(arr)
		verify(arr, t)
	}
	for i := 0; i < 30; i++ {
		arr := []int{2, 2, 2, 1, 1, 1, 2}
		sort.Shuffle(arr)
		verify(arr, t)
	}
	for i := 0; i < 30; i++ {
		arr := []int{0, 0, 0, 1, 1, 1, 1, 0}
		sort.Shuffle(arr)
		verify(arr, t)
	}

}

func verify(arr []int, t *testing.T) {
	sortColors(arr)
	fmt.Println(arr)
	if !sort.IsAscending(arr) {
		t.Fatal("排序错误")
	}
}
