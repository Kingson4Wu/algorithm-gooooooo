package hot100

/**
想太多了，官网就是hash

时间
100 ms
击败
28.37%
内存
8.6 MB
击败
74.42%
*/

func containsDuplicate(nums []int) bool {

	m := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		if m[nums[i]] {
			return true
		}
		m[nums[i]] = true
	}

	return false
}
