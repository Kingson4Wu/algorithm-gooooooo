package sort

import (
	"algorithm-gooooooo/sort"
	"fmt"
	"testing"
)

// 最大堆调整 （自上往下）
func adjust(nums []int, root, len int) {

	//left
	child := root*2 + 1

	// 左右子节点，找个最大的跟根节点交换！！
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

	//init （自底向上）
	//即从length/2 到0，每个根结点使用上述调整堆的方法
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

/** 最大堆， TopK， 取前K个最大值， 且有序 */
func heapSortTopK(nums []int, k int) []int {

	//init
	for i := len(nums) / 2; i >= 0; i-- {
		adjust(nums, i, len(nums))
	}

	//sort
	for i := len(nums) - 1; i > len(nums)-1-k; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		adjust(nums, 0, i)
	}

	return nums[len(nums)-k:]
}

func TestHeapSort(t *testing.T) {

	a := sort.GetArray(100)
	sort.Shuffle(a)
	heapSort(a)
	if !sort.IsAscending(a) {
		t.Fatal("排序错误")
	}

	b := sort.GetArray(100)
	sort.Shuffle(b)
	//前K个最大的
	c := heapSortTopK(b, 10)
	fmt.Println(c)

}
