package sort

import (
	"algorithm-gooooooo/sort"
	"fmt"
	"testing"
)

func bubbleSort(arr []int) {
	n := len(arr)

	if n <= 1 {
		return
	}

	for i := 0; i < n; i++ {
		unChange := true
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				unChange = false
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		if unChange {
			break
		}
	}

}

func TestSort(t *testing.T) {
	arr := []int{43, 7, 1, 6, 98, 1, 0, 35, 8, 3, 76, 98, 243, 572, 234, 2, 6, 9, 1, 4, 9, 6}
	bubbleSort(arr)
	fmt.Println(arr)

	sort.IsAscending(arr)

	a := sort.GetArray(100)
	sort.Shuffle(a)
	bubbleSort(a)
	sort.IsAscending(a)
}
