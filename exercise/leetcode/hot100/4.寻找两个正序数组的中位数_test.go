package hot100

import (
	"fmt"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	fmt.Println(findMedianSortedArrays([]int{}, []int{}))
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{}))
	fmt.Println(findMedianSortedArrays([]int{}, []int{1, 3}))
	fmt.Println(findMedianSortedArrays([]int{2}, []int{}))
	fmt.Println(findMedianSortedArrays([]int{}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2, 3}, []int{}))
	fmt.Println(findMedianSortedArrays([]int{}, []int{1, 2, 3}))

	fmt.Println("===")

	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))

	fmt.Println("===")
	fmt.Println(findMedianSortedArrays([]int{1}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{2}, []int{3}))

	fmt.Println("===///")
	fmt.Println(findMedianSortedArrays([]int{1, 4, 5, 6, 7}, []int{2, 3}))
	fmt.Println(findMedianSortedArrays([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9}))
}
