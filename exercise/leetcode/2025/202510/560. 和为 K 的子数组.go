package _02510

func subarraySum(nums []int, k int) int {
	m := map[int]int{}
	m[0] = 1
	pre := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		count += m[pre-k]
		m[pre]++
	}
	return count
}
