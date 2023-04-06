package hot100

import (
	"fmt"
	"testing"
)

/**
看了题解做的

使用hash存
1、遍历，从当前数逐步加一，到找不到停止，则为由当前数开始的最长
2、判断是否存在更小的数，由更小的数来迭代，避免重复计算

时间
424 ms
击败
27.98%
内存
9 MB
击败
86.73%
*/
/*
*
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

请你设计并实现时间复杂度为 O(n) 的算法解决此问题。

示例 1：

输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
示例 2：

输入：nums = [0,3,7,2,5,8,4,6,0,1]
输出：9
*/
func longestConsecutive(nums []int) int {

	m := make(map[int]bool, len(nums))

	for i := 0; i < len(nums); i++ {
		m[nums[i]] = true
	}

	maxCount := 0
	for i := 0; i < len(nums); i++ {
		if m[nums[i]-1] {
			//由更小的数来迭代，避免重复计算
			continue
		}
		count := 0
		j := 0
		for {
			if m[nums[i]+j] {
				count++
				j++
				continue
			}
			break
		}
		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount
}

func TestLongestConsecutive(t *testing.T) {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
}
