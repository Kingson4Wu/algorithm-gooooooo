package _02510

/*
*
根据回忆
建立一个结果数组，先算该位置左边的所有乘积
在用一个变量从右边累计算起得出最后结果

执行用时分布
4
ms
击败
21.82%
复杂度分析
消耗内存分布
10.54
MB
击败
26.85%
*/
func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	pre := 1
	for i := 0; i < len(nums); i++ {
		ans[i] = pre
		pre *= nums[i]
	}
	pre = 1
	for i := len(nums) - 1; i >= 0; i-- {
		ans[i] *= pre
		pre *= nums[i]
	}
	return ans
}
