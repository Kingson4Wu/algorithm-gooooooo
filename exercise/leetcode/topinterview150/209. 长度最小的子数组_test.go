package topinterview150

import (
	"fmt"
	"math"
	"testing"
)

/**
完全没想法

*/
/**
给定一个含有 n 个正整数的数组和一个正整数 target 。

找出该数组中满足其总和大于等于 target 的长度最小的 子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。


示例 1：

输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。
示例 2：

输入：target = 4, nums = [1,4,4]
输出：1
示例 3：

输入：target = 11, nums = [1,1,1,1,1,1,1,1]
输出：0

子数组必须是连续的元素

自己想的，滑动窗口
遍历一直加，符合条件左边一直减，直到不符合，继续遍历

执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
8.86
MB
击败
54.81%
复杂度分析


*/

func minSubArrayLen(target int, nums []int) int {

	minLength := math.MaxInt32
	sum := 0
	start := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return 1
		}
		sum += nums[i]
		for sum >= target && start <= i {
			if i-start+1 < minLength {
				minLength = i - start + 1
			}
			sum -= nums[start]
			start++
		}
	}
	if minLength == math.MaxInt32 {
		return 0
	}
	return minLength
}

func TestMinSubArrayLen(t *testing.T) {

	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))

}
