package _02510

import (
	"fmt"
	"testing"
)

func topKFrequent(nums []int, k int) []int {

	m := map[int]int{}
	for _, num := range nums {
		m[num]++
	}
	arr := make([][]int, len(m))
	i := 0
	for k, v := range m {
		arr[i] = make([]int, 2)
		arr[i][0] = k
		arr[i][1] = v
		i++
	}

	partition := func(arr [][]int, low, high int) int {
		temp := arr[low]
		for low < high {
			for low < high && arr[high][1] <= temp[1] {
				high--
			}
			if low < high {
				arr[low] = arr[high]
			}
			for low < high && arr[low][1] > temp[1] {
				low++
			}
			if low < high {
				arr[high] = arr[low]
			}
		}
		arr[low] = temp
		return low
	}

	var quickTopK func(arr [][]int, start, end int)
	quickTopK = func(arr [][]int, start, end int) {
		if start >= end {
			return
		}
		mid := partition(arr, start, end)
		if mid >= k-1 {
			quickTopK(arr, 0, mid)
		} else {
			quickTopK(arr, mid+1, end)
		}
	}
	quickTopK(arr, 0, len(arr)-1)
	ans := make([]int, k)
	for i := 0; i < k; i++ {
		ans[i] = arr[i][0]
	}
	return ans
}

/**
执行用时分布
16
ms
击败
6.80%
复杂度分析
消耗内存分布
7.66
MB
击败
31.07%
复杂度分析

*/

func TestTopKFrequent(t *testing.T) {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println(topKFrequent([]int{1}, 1))
}
