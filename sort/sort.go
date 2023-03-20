package sort

import (
	"math/rand"
	"time"
)

func Shuffle(nums []int) {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := len(nums) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func GetArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	return arr
}

func IsAscending(arr []int) bool {
	n := len(arr)
	if n <= 1 {
		return true
	}
	for i := 1; i < n; i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}
	return true
}

func IsDescending(arr []int) bool {
	n := len(arr)
	if n <= 1 {
		return true
	}
	for i := 1; i < n; i++ {
		if arr[i-1] < arr[i] {
			return false
		}
	}
	return true
}
