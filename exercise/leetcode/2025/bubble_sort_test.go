package _025

import (
	"fmt"
	"testing"
)

// 20250729 - 回忆写
// 双重循环 - 两两交换
func bubble(arr []int) []int {
	for i := len(arr) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func TestBubble(t *testing.T) {
	fmt.Println(bubble([]int{4, 2, 6, 7, 1, 9, 8}))
}
