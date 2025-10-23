package _02510

/*
*
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

示例 1：

输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
示例 2：

输入：nums = [0,3,7,2,5,8,4,6,0,1]
输出：9
*/
/**
根据回忆写
全部生成hashmap，遍历时在由更小那个数来通过hashmap计算最大连续序列

执行用时分布
406
ms
击败
16.67%
复杂度分析
消耗内存分布
11.08
MB
击败
27.78%
*/
func longestConsecutive(nums []int) int {

	m := make(map[int]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = true
	}
	maxCnt := 0
	for i := 0; i < len(nums); i++ {
		if m[nums[i]-1] {
			continue
		}
		cnt := 1
		for m[nums[i]+cnt] {
			cnt++
		}
		if cnt > maxCnt {
			maxCnt = cnt
		}
	}
	return maxCnt
}
