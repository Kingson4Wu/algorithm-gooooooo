package _02510

/*
*
根据回忆完成
快慢指针第一次相遇之后，把其中一个指针置为起始点，第二次相遇时就是环的入口
*/
func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	fast = 0
	for {
		slow = nums[slow]
		fast = nums[fast]
		if slow == fast {
			break
		}
	}
	return slow
}

/**
执行用时分布
12
ms
击败
13.54%
复杂度分析
消耗内存分布
9.59
MB
击败
98.28%

*/
