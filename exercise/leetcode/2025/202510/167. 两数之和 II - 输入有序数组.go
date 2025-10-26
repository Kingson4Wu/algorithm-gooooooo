package _02510

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			break
		}
		if sum > target {
			right--
		} else {
			left++
		}
	}
	return []int{left, right}
}

/**
执行用时分布
4
ms
击败
74.51%
复杂度分析
消耗内存分布
4.93
MB
击败
25.49%
*/
