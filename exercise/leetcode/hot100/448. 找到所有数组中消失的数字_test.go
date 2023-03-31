package hot100

import (
	"fmt"
	"testing"
)

/**
似曾相识的做出来了
利用负数和下标的技巧，替换hashmap
*/
/*
*
给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。请你找出所有在 [1, n] 范围内但没有出现在 nums 中的数字，并以数组的形式返回结果。

示例 1：

输入：nums = [4,3,2,7,8,2,3,1]
输出：[5,6]
示例 2：

输入：nums = [1,1]
输出：[2]
*/
func findDisappearedNumbers(nums []int) []int {

	for i := 0; i < len(nums); i++ {
		index := abs(nums[i])
		if nums[index-1] < 0 {
			continue
		}
		nums[index-1] = -nums[index-1]
	}

	var result []int
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			result = append(result, i+1)
		}
	}

	return result
}

func abs(num int) int {
	if num > 0 {
		return num
	}
	return -num
}

func TestFindDisappearedNumbers(t *testing.T) {

	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println(findDisappearedNumbers([]int{1, 1}))
}
