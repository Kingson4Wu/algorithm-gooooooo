package top100liked

import (
	"fmt"
	"testing"
)

/*
*
给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

示例 1:

输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]
示例 2:

输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右轮转 1 步: [99,-1,-100,3]
向右轮转 2 步: [3,99,-1,-100]
*/
/**
自己画图找出来的规律

（1）
执行出错
2 / 38 个通过的测试用例
panic: runtime error: integer divide by zero
main.rotate({0xc0000141f8?, 0xc000014200?, 0x1?}, 0x4cd580?)
solution.go, line 14
main.__helper__(...)
solution.go, line 43
main.main()
solution.go, line 84
最后执行的输入
添加到测试用例
nums =
[1]
k =
0

（2）

解答错误
34 / 38 个通过的测试用例

官方题解
输入
nums =
[1,2,3,4,5,6]
k =
4

添加到测试用例
输出
[5,2,1,4,1,6]
预期结果
[3,4,5,6,1,2]

执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
10.16
MB
击败
5.22%

提交了三次才通过
*/
/*func rotate(nums []int, k int) {

	if len(nums) == 0 {
		return
	}
	if k >= len(nums) {
		k %= len(nums)
	}
	// 忘记k==0的边界！
	if k == 0 {
		return
	}

	if len(nums)%k == 0 {
		times := len(nums) / k
		index := 0
		for index < k {
			temp := nums[index]
			for i := 0; i < times; i++ {
				next := (index + k) % len(nums)
				nums[next], temp = temp, nums[next]
				index = next
			}
			index++
		}
	} else {
		count := 0
		index := 0
		temp := nums[0]
		for count < len(nums) {
			next := (index + k) % len(nums)
			nums[next], temp = temp, nums[next]
			index = next
			count++
		}
	}

}*/

func rotate(nums []int, k int) {

	if len(nums) == 0 {
		return
	}
	if k >= len(nums) {
		k %= len(nums)
	}
	// 忘记k==0的边界！
	if k == 0 {
		return
	}

	count := 0
	index := 0
	temp := nums[0]
	start := 0
	for count < len(nums) {
		next := (index + k) % len(nums)
		nums[next], temp = temp, nums[next]
		index = next
		count++
		if index == start {
			start++
			index = start
			temp = nums[index]
		}
	}

}

func TestRotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)

	nums = []int{-1, -100, 3, 99}
	rotate(nums, 2)
	fmt.Println(nums)

	nums = []int{-1, -100, 3, 99}
	rotate(nums, 3)
	fmt.Println(nums)

	nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	rotate(nums, 4)
	fmt.Println(nums)

	nums = []int{1}
	rotate(nums, 0)
	fmt.Println(nums)

	nums = []int{1, 2, 3, 4, 5, 6}
	rotate(nums, 4)
	fmt.Println(nums)
}
