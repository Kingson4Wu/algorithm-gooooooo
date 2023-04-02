package hot100

/**

看了题解再实现的

对数字进行计数，相同加1，不同减一
根据题目的前提条件，多数元素大于【n/2】，那么最后存在的那个数一定是多数元素

*/

/*
*
给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

示例 1：

输入：nums = [3,2,3]
输出：3
示例 2：

输入：nums = [2,2,1,1,1,2,2]
输出：2

进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。
*/
func majorityElement(nums []int) int {

	count := 0
	current := 0
	for _, num := range nums {

		if count == 0 {
			current = num
			count++
			continue
		}

		if current == num {
			count++
			continue
		}
		count--
	}
	return current
}
