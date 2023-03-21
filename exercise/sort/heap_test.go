package sort

import (
	"algorithm-gooooooo/sort"
	"testing"
)

func adjust(nums []int, root int) {

	length := len(nums)
	//left
	child := root*2 + 1

	for child < length {
		if child+1 < length && nums[child+1] > nums[child] {
			//right
			child++
		}

		if nums[child] < nums[root] {
			nums[child], nums[root] = nums[root], nums[child]
		}

		root = child
		//left
		child = root*2 + 1
	}
}

func heapSort(nums []int) {

	//init
	for i := len(nums) / 2; i >= 0; i-- {
		adjust(nums, i)
	}

	//sort
	for i := len(nums) - 1; i > 0; i-- {
		nums[i] = nums[0]
		adjust(nums, 0)
	}

}

func TestHeapSort(t *testing.T) {

	a := sort.GetArray(100)
	sort.Shuffle(a)
	heapSort(a)
	if !sort.IsAscending(a) {
		t.Fatal("排序错误")
	}

	/*b := sort.GetArray(100)
	sort.Shuffle(b)
	quickSortTopK(b, 10)
	fmt.Println(b)*/

}
