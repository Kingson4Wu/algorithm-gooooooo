package _02510

/*
*
根据回忆写

执行用时分布
18
ms
击败
76.47%
复杂度分析
消耗内存分布
10.45
MB
击败
27.94%
*/
func minCapability(nums []int, k int) int {
	upper, lower := nums[0], nums[0]
	for _, num := range nums {
		if upper < num {
			upper = num
		} else if lower > num {
			lower = num
		}
	}
	for lower <= upper {
		middle := (lower + upper) / 2
		visit := false
		cnt := 0
		for i := 0; i < len(nums); i++ {
			if !visit && nums[i] <= middle {
				cnt++
				visit = true
				if cnt >= k {
					break
				}
			} else {
				visit = false
			}
		}
		if cnt >= k {
			upper = middle - 1
		} else {
			lower = middle + 1
		}
	}
	return lower
}
