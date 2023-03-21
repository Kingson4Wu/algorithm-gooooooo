package sort

import (
	"algorithm-gooooooo/sort"
	"testing"
)

func adjust(nums []int, root, len int) {

	//left
	child := root*2 + 1

	for child < len {
		if child+1 < len && nums[child+1] > nums[child] {
			//right
			child++
		}

		//这步别忘了 ！！！
		if nums[child] <= nums[root] {
			break
		}

		nums[child], nums[root] = nums[root], nums[child]

		root = child
		//left
		child = root*2 + 1
	}
}

func heapSort(nums []int) {

	//init
	for i := len(nums) / 2; i >= 0; i-- {
		adjust(nums, i, len(nums))
	}

	//sort
	for i := len(nums) - 1; i > 0; i-- {
		//这里是交换！！！！
		nums[i], nums[0] = nums[0], nums[i]

		//这里的调整要指定长度，因为最后的几位已经排序了，不应该继续被堆调整影响
		//这里是i，因为是调整交换过前面的，第一次，i刚好是len(nums) - 1， 即最后一位不调整
		adjust(nums, 0, i)
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
