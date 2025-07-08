package topinterview145

/**
给你四个整数数组 nums1、nums2、nums3 和 nums4 ，数组长度都是 n ，请你计算有多少个元组 (i, j, k, l) 能满足：

0 <= i, j, k, l < n
nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0


示例 1：

输入：nums1 = [1,2], nums2 = [-2,-1], nums3 = [-1,2], nums4 = [0,2]
输出：2
解释：
两个元组如下：
1. (0, 0, 0, 1) -> nums1[0] + nums2[0] + nums3[0] + nums4[1] = 1 + (-2) + (-1) + 2 = 0
2. (1, 1, 0, 0) -> nums1[1] + nums2[1] + nums3[0] + nums4[0] = 2 + (-1) + (-1) + 0 = 0
示例 2：

输入：nums1 = [0], nums2 = [0], nums3 = [0], nums4 = [0]
输出：1


  提示：

n == nums1.length
n == nums2.length
n == nums3.length
n == nums4.length
1 <= n <= 200
-228 <= nums1[i], nums2[i], nums3[i], nums4[i] <= 228
*/

/*
看答案
方法一：分组 + 哈希表
将四个数组分成两部分，A 和 B 为一组，C 和 D 为另外一组。
对于 A 和 B，我们使用二重循环对它们进行遍历，得到所有 A[i]+B[j] 的值并存入哈希映射中
对于 C 和 D，我们同样使用二重循环对它们进行遍历。当遍历到 C[k]+D[l] 时，如果 −(C[k]+D[l]) 出现在哈希映射中，那么将 −(C[k]+D[l]) 对应的值累加进答案中。

*/

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {

	ans := 0
	m := make(map[int]int)
	for _, u := range nums1 {
		for _, v := range nums2 {
			m[u+v]++
		}
	}
	for _, u := range nums3 {
		for _, v := range nums4 {
			ans += m[-u-v]
		}
	}
	return ans
}

/**
执行用时分布
105
ms
击败
55.06%
复杂度分析
消耗内存分布
8.25
MB
击败
36.54%
*/
