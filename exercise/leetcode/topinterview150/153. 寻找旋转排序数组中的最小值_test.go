package topinterview150

import (
	"fmt"
	"testing"
)

/*
*
已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。例如，原数组 nums = [0,1,2,4,5,6,7] 在变化后可能得到：
若旋转 4 次，则可以得到 [4,5,6,7,0,1,2]
若旋转 7 次，则可以得到 [0,1,2,4,5,6,7]
注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。

给你一个元素值 互不相同 的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。

你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。

示例 1：

输入：nums = [3,4,5,1,2]
输出：1
解释：原数组为 [1,2,3,4,5] ，旋转 3 次得到输入数组。
示例 2：

输入：nums = [4,5,6,7,0,1,2]
输出：0
解释：原数组为 [0,1,2,4,5,6,7] ，旋转 4 次得到输入数组。
示例 3：

输入：nums = [11,13,15,17]
输出：11
解释：原数组为 [11,13,15,17] ，旋转 4 次得到输入数组。
*/

/*
*
印象中二分查找，尝试回忆实现
*/
func findMin(nums []int) int {

	var divide func([]int, int, int) int

	divide = func(nums []int, start int, end int) int {
		if nums[start] < nums[end] || start == end {
			return nums[start]
		} else {
			//mid := (end - start) / 2
			// 是 end + start ！！！
			mid := (end + start) / 2
			a := divide(nums, start, mid)
			b := divide(nums, mid+1, end)
			return min(a, b)
		}
	}

	return divide(nums, 0, len(nums)-1)
}

/**
执行出错
10 / 150 个通过的测试用例
panic: runtime error: index out of range [-1]
main.findMin.func1({0xc00008a0c0?, 0xc000012108?, 0xc00004dc30?}, 0x439d07?, 0x440ca6?)
solution.go, line 6
main.findMin.func1({0xc00008a0c0, 0x5, 0x8}, 0x4689ae?, 0x0)
solution.go, line 10
main.findMin.func1({0xc00008a0c0, 0x5, 0x8}, 0x4632cb?, 0x4)
solution.go, line 10
main.findMin.func1({0xc00008a0c0, 0x5, 0x8}, 0xc00004dd00?, 0x4)
solution.go, line 11
main.findMin({0xc00008a0c0?, 0xb?, 0x10?})
查看更多
最后执行的输入
添加到测试用例
nums =
[2,3,4,5,1]
*/

func TestFindMin(t *testing.T) {
	fmt.Println(findMin([]int{2, 3, 4, 5, 1}))
}

/**
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
4.23
MB
击败
24.58%

*/

/**
题解更简洁
func findMin(nums []int) int {
    low, high := 0, len(nums) - 1
    for low < high {
        pivot := low + (high - low) / 2
        if nums[pivot] < nums[high] {
            high = pivot
        } else {
            low = pivot + 1
        }
    }
    return nums[low]
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/solutions/698479/xun-zhao-xuan-zhuan-pai-xu-shu-zu-zhong-5irwp/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
