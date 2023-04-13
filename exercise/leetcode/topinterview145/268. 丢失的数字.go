package topinterview145

/**

仿照之前的题解做的

可以和第41题做对比
*/

func missingNumber(nums []int) int {

	for i := 0; i < len(nums); i++ {

		index := nums[i]
		for {
			if index >= len(nums) || nums[index] == index {
				break
			}
			nums[index], index = index, nums[index]
		}

	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}
	return len(nums)
}
