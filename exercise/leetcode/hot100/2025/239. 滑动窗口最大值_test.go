package _025

import (
	"fmt"
	"testing"
)

/**
看完题解自己写

单调队列（按数值单调递减，存下标）
当滑动窗口向右移动时， i<j, 如果 nums[i]≤nums[j]， 那么可以删除nums[i]
使用一个队列存储所有还没有被移除的下标。在队列中，这些下标按照从小到大的顺序被存储，并且它们在数组 nums 中对应的值是严格单调递减的
遍历元素，先push进队列，然后比对队列从0开始下标是否符合窗口，不符合删除，队列是按数值单调递减的

*/

func maxSlidingWindow(nums []int, k int) []int {

	ans := make([]int, len(nums)-k+1)

	var queue []int
	push := func(i int) {
		for len(queue) > 0 && nums[i] >= nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}
	ans[0] = nums[queue[0]]
	for i := k; i < len(nums); i++ {
		push(i)
		//去掉超过k长度的窗口下标
		for queue[0] < i-k+1 {
			queue = queue[1:]
		}
		ans[i-k+1] = nums[queue[0]]
	}
	return ans
}

/**
解答错误
9 / 51 个通过的测试用例

官方题解
输入
nums =
[7,2,4]
k =
2

添加到测试用例
输出
[2,4]
预期结果
[7,4]
*/

func TestMaxSlidingWindow(t *testing.T) {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow([]int{7, 2, 4}, 2))
}

/**
执行用时分布
2
ms
击败
97.40%
复杂度分析
消耗内存分布
10.33
MB
击败
38.60%
*/
