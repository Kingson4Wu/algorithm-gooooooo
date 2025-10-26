package _02510

/**
根据回忆写
应该是前两个双重循环先构造map
后两个双重循环判断map是否存在同时计数

执行用时分布
103
ms
击败
63.89%
复杂度分析
消耗内存分布
8.26
MB
击败
35.83%
*/

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m := map[int]int{}
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			m[nums1[i]+nums2[j]]++
		}
	}
	count := 0
	for i := 0; i < len(nums3); i++ {
		for j := 0; j < len(nums4); j++ {
			count += m[-(nums3[i] + nums4[j])]
		}
	}
	return count
}
