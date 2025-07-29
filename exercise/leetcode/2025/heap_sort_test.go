package _025

import (
	"fmt"
	"testing"
)

func heapSort(arr []int) []int {

	//从上到下的调整函数
	//最大堆
	adjust := func(arr []int, root, size int) {
		index := root*2 + 1
		for index < size {
			if index+1 < size && arr[index] < arr[index+1] {
				index++
			}
			if arr[index] > arr[root] {
				arr[index], arr[root] = arr[root], arr[index]
			} else {
				break
			}
			root = index
			index = root*2 + 1
		}
	}
	//从最后一个跟结点开始初始堆
	for i := len(arr)/2 - 1; i >= 0; i-- {
		adjust(arr, i, len(arr))
	}
	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		adjust(arr, 0, i)
	}
	return arr
}

func TestHeap(t *testing.T) {
	fmt.Println(heapSort([]int{4, 2, 6, 7, 1, 9, 8}))
}
