package sort

import (
	"fmt"
	"testing"
)

var arr = []int{43, 7, 1, 6, 98, 1, 0, 35, 8, 3, 76, 98, 243, 572, 234, 2, 6, 9, 1, 4, 9, 6}

func TestBubbleSort(t *testing.T) {

	BubbleSort(arr, len(arr))
	fmt.Println(arr)
}

func TestHeapSort(t *testing.T) {

	HeadSort(arr)
	fmt.Println(arr)
}

func TestQuickSortTopk_1(t *testing.T) {
	a := GetArray(100)
	Shuffle(a)
	r := QuickSortTopk_1(a, 33)
	//[2 3 0 1 4 5 6 7 9 8 10]
	fmt.Println(r)

	aa := []int{5, 2, 8, 3, 1, 6, 4, 9, 7}
	rr := QuickSortTopk_1(aa, 6)
	fmt.Println(rr)
}

func TestHeadSort(t *testing.T) {
	a := GetArray(100)
	Shuffle(a)

	if IsSort(a) {
		t.Fatal("数组已经有序")
	}

	HeadSort(a)

	if !IsSort(a) {
		t.Fatal("数组排序失败")
	}
}

func TestHeadSort2(t *testing.T) {
	a := []int{7, 3, 4, 9, 2, 8, 1, 5, 6}
	//{0, 1, 2, 3, 4, 5, 6, 7, 8}

	HeadSort(a)

	if !IsSort(a) {
		t.Fatal("数组排序失败")
	}
}
