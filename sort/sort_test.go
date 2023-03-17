package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {

	arr := []int{43, 7, 1, 6, 98, 1, 0, 35, 8, 3, 76, 98, 243, 572, 234, 2, 6, 9, 1, 4, 9, 6}
	BubbleSort(arr, len(arr))
	fmt.Println(arr)
}
