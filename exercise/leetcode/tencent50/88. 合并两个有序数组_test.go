package tencent50

import (
	"fmt"
	"testing"
)

/*
*
从后面开始遍历， 注意m==0的情况

时间
0 ms
击败
100%
内存
2.2 MB
击败
43.54%
*/

/*
*
示例 1：

输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
输出：[1,2,2,3,5,6]
解释：需要合并 [1,2,3] 和 [2,5,6] 。
合并结果是 [1,2,2,3,5,6] ，其中斜体加粗标注的为 nums1 中的元素。
示例 2：

输入：nums1 = [1], m = 1, nums2 = [], n = 0
输出：[1]
解释：需要合并 [1] 和 [] 。
合并结果是 [1] 。
示例 3：

输入：nums1 = [0], m = 0, nums2 = [1], n = 1
输出：[1]
解释：需要合并的数组是 [] 和 [1] 。
合并结果是 [1] 。
注意，因为 m = 0 ，所以 nums1 中没有元素。nums1 中仅存的 0 仅仅是为了确保合并结果可以顺利存放到 nums1 中。
*/
func merge(nums1 []int, m int, nums2 []int, n int) {

	//j<0 说明数组1为空
	j := m - 1
	index := m + n - 1
	for i := n - 1; i >= 0; i-- {
		if j < 0 || nums2[i] > nums1[j] {
			nums1[index] = nums2[i]
		} else {
			nums1[index] = nums1[j]
			j--
			i++
		}
		index--
	}
}

func TestMerge(t *testing.T) {

	arr1 := []int{1, 2, 3, 0, 0, 0}
	merge(arr1, 3, []int{2, 5, 6}, 3)
	fmt.Println(arr1)
	arr2 := []int{1}
	merge(arr2, 1, []int{}, 0)
	fmt.Println(arr2)
	arr3 := []int{0}
	merge(arr3, 0, []int{1}, 1)
	fmt.Println(arr3)
}
