package hot100

import (
	"fmt"
	"testing"
)

/**
看过题解做的

两个dp，保存正负的成积
dp1 以i结尾的成绩最大值
dp2 以i结尾的成绩最小值

dp1[i] = max(dp1[i-1] * num[i], dp2[i-1] * num[i], num[i])
dp2[i] = min(dp2[i-1] * num[i], dp2[i-1] * num[i], num[i])

时间
8 ms
击败
31.86%
内存
3.8 MB
击败
8.2%

*/
/*
*
给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

测试用例的答案是一个 32-位 整数。

子数组 是数组的连续子序列。

示例 1:

输入: nums = [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。
示例 2:

输入: nums = [-2,0,-1]
输出: 0
解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
*/
func maxProduct(nums []int) int {

	dp1 := make([]int, len(nums)+1)
	dp2 := make([]int, len(nums)+1)
	dp1[0], dp2[0] = 1, 1

	maxValue := nums[0]

	for i := 1; i <= len(nums); i++ {

		a := dp1[i-1] * nums[i-1]
		b := dp2[i-1] * nums[i-1]

		dp1[i] = max(max(a, b), nums[i-1])
		dp2[i] = min(min(a, b), nums[i-1])

		if dp1[i] > maxValue {
			maxValue = dp1[i]
		}
	}

	return maxValue
}

func TestMaxProduct(t *testing.T) {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	fmt.Println(maxProduct([]int{-2, 0, -1}))

	fmt.Println(maxProduct([]int{-2}))
	fmt.Println(maxProduct([]int{-2, -4, 2}))

	fmt.Println(maxProduct([]int{-4, 2}))

	fmt.Println(maxProduct([]int{-4, 2, 1, -1, -1, 2, 4, 1}))
}
