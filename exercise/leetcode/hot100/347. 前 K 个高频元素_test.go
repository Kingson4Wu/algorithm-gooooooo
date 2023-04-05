package hot100

import (
	"fmt"
	"testing"
)

/*
*
给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。

示例 1:

输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
示例 2:

输入: nums = [1], k = 1
输出: [1]
*/
func topKFrequent(nums []int, k int) []int {

	var ans []int

	return ans
}

/**
好烦，直接hashmap
*/
//题目理解错误，前k高，不是出现次数大于等于k
/*func topKFrequent(nums []int, k int) []int {

	var ans []int
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]] += 1
		if m[nums[i]] == k {
			ans = append(ans, nums[i])
		}
	}
	return ans
}*/

func TestTopKFrequent(t *testing.T) {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println(topKFrequent([]int{1}, 1))
}

/**
nums =
[1,2]
k =
2
添加到测试用例
输出
[]
预期结果
[1,2]
*/
