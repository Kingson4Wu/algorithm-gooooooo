package sort

import (
	"algorithm-gooooooo/sort"
	"fmt"
	"testing"
)

func quickSort(arr []int) {

	var quick func(arr []int, start, end int)
	quick = func(arr []int, start, end int) {
		if start >= end {
			return
		}
		mid := partition(arr, start, end)
		quick(arr, start, mid)
		quick(arr, mid+1, end)
	}

	quick(arr, 0, len(arr)-1)

}

// low < high 的判断每一步都要，不要忘了！！
func partition(arr []int, low, high int) int {

	temp := arr[low]
	for low < high {
		for low < high && arr[high] >= temp {
			high--
		}
		if low < high {
			arr[low] = arr[high]
		}

		for low < high && arr[low] < temp {
			low++
		}
		if low < high {
			arr[high] = arr[low]
		}

	}
	arr[low] = temp

	return low
}

func quickSortTopK(arr []int, k int) {

	var quick func(arr []int, start, end, k int)
	quick = func(arr []int, start, end, k int) {
		if start >= end {
			return
		}
		mid := partition(arr, start, end)
		if mid > k {
			quick(arr, start, mid, k)
		} else if mid < k {
			quick(arr, mid+1, end, k)
		}
	}

	quick(arr, 0, len(arr)-1, k)

}

func TestQuickSort(t *testing.T) {

	a := sort.GetArray(100)
	sort.Shuffle(a)
	quickSort(a)
	if !sort.IsAscending(a) {
		t.Fatal("排序错误")
	}

	b := sort.GetArray(100)
	sort.Shuffle(b)
	//前K个最小的
	quickSortTopK(b, 10)
	fmt.Println(b)

}
