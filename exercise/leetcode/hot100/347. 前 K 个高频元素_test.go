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
/**
用hashmap记录数字和对应的出现个数
因为是前k个，使用快排即可，不过快排最差时间n的平方

*/
func topKFrequent(nums []int, k int) []int {

	times := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		times[nums[i]]++
	}

	arr := make([][]int, len(times))
	i := 0
	for k, v := range times {
		arr[i] = []int{k, v}
		i++
	}

	partition := func(arr [][]int, low, high int) int {
		temp := arr[low]
		for low < high {
			for low < high && arr[high][1] <= arr[low][1] {
				high--
			}
			if low < high {
				arr[low] = arr[high]
			}
			for low < high && arr[low][1] > arr[high][1] {
				low--
			}
			if low < high {
				arr[high] = arr[low]
			}
			arr[low] = temp
		}
		return low
	}

	var quickSortTopK func(arr [][]int, start, end int)
	quickSortTopK = func(arr [][]int, start, end int) {
		if start >= end {
			return
		}

		mid := partition(arr, start, end)
		if mid >= k-1 {
			quickSortTopK(arr, start, mid)
		} else {
			quickSortTopK(arr, mid+1, end)
		}
	}
	quickSortTopK(arr, 0, len(arr)-1)
	ans := make([]int, k)
	for i := 0; i < k; i++ {
		ans[i] = arr[i][0]
	}
	return ans
}

/**
超出时间限制
4 / 22 个通过的测试用例
最后执行的输入
添加到测试用例
nums =
[3,0,1,0]
k =
1
*/

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
