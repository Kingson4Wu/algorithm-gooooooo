package _02510

/*
*
根据回忆应该是队列, 肯定是递减的

执行用时分布
120
ms
击败
67.21%
复杂度分析
消耗内存分布
10.36
MB
击败
24.59%
*/
func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	var queue []int
	queue = append(queue, 0)
	for i := 1; i < len(temperatures); i++ {
		for len(queue) > 0 && temperatures[i] > temperatures[queue[len(queue)-1]] {
			ans[queue[len(queue)-1]] = i - (queue[len(queue)-1])
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
	}
	return ans
}
