package topinterview145

import (
	"sort"
	"testing"
)

/**
给你一个整数数组 nums，将它重新排列成 nums[0] < nums[1] > nums[2] < nums[3]... 的顺序。

你可以假设所有输入数组都可以得到满足题目要求的结果。



示例 1：

输入：nums = [1,5,1,1,6,4]
输出：[1,6,1,5,1,4]
解释：[1,4,1,5,1,6] 同样是符合题目要求的结果，可以被判题程序接受。
示例 2：

输入：nums = [1,3,2,2,3,1]
输出：[2,3,1,3,1,2]


提示：

1 <= nums.length <= 5 * 104
0 <= nums[i] <= 5000
题目数据保证，对于给定的输入 nums ，总能产生满足题目要求的结果


进阶：你能用 O(n) 时间复杂度和 / 或原地 O(1) 额外空间来实现吗？
*/
/**
题解

排序 + 交错填充
步骤：
1. 排序原数组（升序）；
2.将排好序的数组，从后往前分成两半：
	- 一半放在 奇数位置（大的数）；
	- 另一半放在 偶数位置（小的数）；
3. 为什么这么做？
这样能保证 小 < 大 > 小 < 大 的交错结构。

分两段，倒着穿插交叉插入
1. 从前半部分末尾向前拿 → 填到 偶数位（0,2,4,...）
2. 从后半部分末尾向前拿 → 填到 奇数位（1,3,5,...）
这样可以保证小的和大的交错，且重复值分开，避免相邻相等。

*/
func wiggleSort(nums []int) {

	ans := make([]int, len(nums))
	copy(ans, nums)
	sort.Slice(ans, func(i, j int) bool {
		return ans[i] < ans[j]
	})

	left, right := (len(ans)-1)/2, len(ans)-1
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			nums[i] = ans[left]
			left--
		} else {
			nums[i] = ans[right]
			right--
		}
	}
}

func TestWiggleSort(t *testing.T) {
	wiggleSort([]int{1, 5, 1, 1, 6, 4})
	wiggleSort([]int{1, 3, 2, 2, 3, 1})
	wiggleSort([]int{1, 3, 2, 2, 3, 1})
}

/**
解答错误
29 / 52 个通过的测试用例

官方题解
输入
nums =
[1,1,2,1,2,2,1]

添加到测试用例
输出
[1,1,1,2,1,2,1]
预期结果
[1,2,1,2,1,2,1]
*/

/**
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
7.96
MB
击败
45.10%
*/
