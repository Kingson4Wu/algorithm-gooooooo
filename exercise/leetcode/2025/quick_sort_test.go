package _025

import (
	"fmt"
	"testing"
)

func quickSort(arr []int) []int {

	partition := func(arr []int, low, high int) int {
		temp := arr[low]
		for low < high {
			for low < high && arr[high] > temp {
				high--
			}
			if low < high {
				arr[low] = arr[high]
			}
			for low < high && arr[low] <= temp {
				low++
			}
			if low < high {
				arr[high] = arr[low]
			}
		}
		arr[low] = temp
		return low
	}

	var quick func(arr []int, start, end int)
	quick = func(arr []int, start, end int) {
		if start >= end {
			return
		}
		mid := partition(arr, start, end)
		quick(arr, start, mid-1)
		quick(arr, mid+1, end)
	}
	//quick(arr, 0, len(arr)-1)

	var quickTopK func(arr []int, start, end, k int)
	quickTopK = func(arr []int, start, end, k int) {
		if start >= end {
			return
		}
		mid := partition(arr, start, end)
		if mid > k-1 {
			quickTopK(arr, start, mid-1, k)
		} else if mid < k-1 {
			quickTopK(arr, mid+1, end, k)
		}
	}
	quickTopK(arr, 0, len(arr)-1, 3)
	return arr
}

func TestQuick(t *testing.T) {
	fmt.Println(quickSort([]int{4, 2, 6, 7, 1, 9, 8}))
}
