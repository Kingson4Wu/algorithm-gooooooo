package hot100

/**
竟然是困难。。。
自己做的。。
基本思路，记录数组的最大最小下标，最小的之间进行对比，最大的之间进行对比，然后开始下标++，对终止下标--
看了题解，还能优化，因为是中位数，只需要对比开始下标，直到总长度为一半即可，减少一半时间！
*/
/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数

 4. 寻找两个正序数组的中位数
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

算法的时间复杂度应该为 O(log (m+n)) 。

示例 1：

输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
示例 2：

输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5

执行结果：
通过
显示详情
添加备注

执行用时：
16 ms
, 在所有 Go 提交中击败了
44.88%
的用户
内存消耗：
4.9 MB
, 在所有 Go 提交中击败了
67.06%
的用户
通过测试用例：
2094 / 2094
*/

/**
方法一：二分查找
给定两个有序数组，要求找到两个有序数组的中位数，最直观的思路有以下两种：

使用归并的方式，合并两个有序数组，得到一个大的有序数组。大的有序数组的中间位置的元素，即为中位数。

不需要合并两个有序数组，只要找到中位数的位置即可。由于两个数组的长度已知，因此中位数对应的两个数组的下标之和也是已知的。维护两个指针，初始时分别指向两个数组的下标 00 的位置，每次将指向较小值的指针后移一位（如果一个指针已经到达数组末尾，则只需要移动另一个数组的指针），直到到达中位数的位置。

方法二：划分数组
说明

方法一的时间复杂度已经很优秀了，但本题存在时间复杂度更低的一种方法。这里给出推导过程，勇于挑战自己的读者可以进行尝试。

思路与算法

为了使用划分的方法解决这个问题，需要理解「中位数的作用是什么」。在统计中，中位数被用来：

将一个集合划分为两个长度相等的子集，其中一个子集中的元素总是大于另一个子集中的元素。

！！将两个数组，在一半划分！判断小的那两半最大的那个数即可（可能需要对奇偶数情况做特殊处理！）

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays/solution/xun-zhao-liang-ge-you-xu-shu-zu-de-zhong-wei-s-114/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

/**
digits := []int{2, 1, 3, 0}
	leetcode.Exexute(digits)
*/

func Exexute(nums1 []int, nums2 []int) float64 {
	return findMedianSortedArrays(nums1, nums2)
}

// @lc code=start
/*func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	nums1_length := len(nums1)
	nums2_length := len(nums2)

	total_length := nums1_length + nums2_length

	nums1_start_index, nums2_start_index := 0, 0
	nums1_end_index, nums2_end_index := nums1_length-1, nums2_length-1

	for total_length > 2 {

		if nums1_length == 0 {
			nums2_start_index++
			nums2_end_index--
			nums2_length -= 2
			total_length -= 2
			continue
		}
		if nums2_length == 0 {
			nums1_start_index++
			nums1_end_index--
			nums1_length -= 2
			total_length -= 2
			continue
		}

		//去掉最小的数
		if nums1[nums1_start_index] <= nums2[nums2_start_index] {
			nums1_start_index++
			nums1_length--
			total_length--
		} else {
			nums2_start_index++
			nums2_length--
			total_length--
		}

		//去掉最大的数
		if nums1_length == 0 {
			nums2_end_index--
			nums2_length--
			total_length--
			continue
		}
		if nums2_length == 0 {
			nums1_end_index--
			nums1_length--
			total_length--
			continue
		}

		if nums1[nums1_end_index] >= nums2[nums2_end_index] {
			nums1_end_index--
			nums1_length--
			total_length--
		} else {
			nums2_end_index--
			nums2_length--
			total_length--
		}

	}

	if total_length == 2 {
		if nums1_length == 0 {
			return float64(nums2[nums2_start_index]+nums2[nums2_end_index]) / 2
		}
		if nums2_length == 0 {
			return float64(nums1[nums1_start_index]+nums1[nums1_end_index]) / 2
		}
		return float64(nums1[nums1_start_index]+nums2[nums2_start_index]) / 2
	}

	if nums1_length == 1 {
		return float64(nums1[nums1_start_index])
	}
	return float64(nums2[nums2_start_index])

}*/

// @lc code=end

// 20250623
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0
	}
	if len(nums1) == 0 {
		if len(nums2)%2 == 0 {
			return float64(nums2[len(nums2)/2]+nums2[len(nums2)/2-1]) / 2
		} else {
			return float64(nums2[len(nums2)/2])
		}
	}
	if len(nums2) == 0 {
		if len(nums1)%2 == 0 {
			return float64(nums1[len(nums1)/2]+nums1[len(nums1)/2-1]) / 2
		} else {
			return float64(nums1[len(nums1)/2])
		}
	}
	totalLength := len(nums1) + len(nums2)
	nums1Index := 0
	nums2Index := 0
	endLength := totalLength/2 + 1
	count := 0
	num1, num2 := 0, 0
	for nums1Index < len(nums1) && nums2Index < len(nums2) && count < endLength {
		count++
		num1 = num2
		if nums1[nums1Index] < nums2[nums2Index] {
			num2 = nums1[nums1Index]
			nums1Index++
			continue
		}
		num2 = nums2[nums2Index]
		nums2Index++
	}
	if count < endLength {
		if nums1Index == len(nums1) {
			index := nums2Index + (endLength - count) - 1
			num1 = num2
			num2 = nums2[index]
			if index-1 >= 0 {
				if nums2[index-1] > num1 {
					num1 = nums2[index-1]
				}
			}
		}
		if nums2Index == len(nums2) {
			index := nums1Index + (endLength - count) - 1
			num1 = num2
			num2 = nums1[index]
			if index-1 >= 0 {
				if nums1[index-1] > num1 {
					num1 = nums1[index-1]
				}
			}
		}
	}

	if totalLength%2 == 0 {
		return float64(num1+num2) / 2
	} else {
		return float64(num2)
	}
}

/**
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
6.30
MB
击败
68.13%

*/
