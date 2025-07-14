package topinterview150

import (
	"math/bits"
	"sort"
)

/*
*
给你一个整数数组 nums ，按要求返回一个新数组 counts 。数组 counts 有该性质： counts[i] 的值是  nums[i] 右侧小于 nums[i] 的元素的数量。

示例 1：

输入：nums = [5,2,6,1]
输出：[2,1,1,0]
解释：
5 的右侧有 2 个更小的元素 (2 和 1)
2 的右侧仅有 1 个更小的元素 (1)
6 的右侧有 1 个更小的元素 (1)
1 的右侧有 0 个更小的元素
示例 2：

输入：nums = [-1]
输出：[0]
示例 3：

输入：nums = [-1,-1]
输出：[0,0]

提示：

1 <= nums.length <= 105
-104 <= nums[i] <= 104
*/
/**
模仿题解抄的
*/
func countSmaller(nums []int) []int {

	//离散化
	numToId := func(nums []int) map[int]int {
		set := make(map[int]struct{})
		for _, num := range nums {
			set[num] = struct{}{}
		}
		a := make([]int, 0, len(nums))
		for num := range set {
			a = append(a, num)
		}
		sort.Ints(a)
		m := make(map[int]int)
		for i, num := range a {
			m[num] = i + 1
		}
		return m
	}(nums)
	c := make([]int, 1<<(bits.Len(uint(len(numToId))))+1)
	lowBit := func(x int) int {
		return x & -x
	}
	query := func(pos int) int {
		ret := 0
		for pos > 0 {
			ret += c[pos]
			pos -= lowBit(pos)
		}
		return ret
	}
	update := func(pos int) {
		for pos < len(c) {
			c[pos]++
			pos += lowBit(pos)
		}
	}

	ans := make([]int, len(nums))
	c = make([]int, len(nums)+5)
	//从后往前计算
	for i := len(nums) - 1; i >= 0; i-- {
		id := numToId[nums[i]]
		ans[i] = query(id - 1)
		update(id)
	}
	return ans
}

/**
执行用时分布
49
ms
击败
79.45%
消耗内存分布
11.07
MB
击败
69.86%

*/
