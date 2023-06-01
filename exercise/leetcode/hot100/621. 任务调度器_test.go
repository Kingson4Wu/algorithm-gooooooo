package hot100

import "testing"

/**
给你一个用字符数组 tasks 表示的 CPU 需要执行的任务列表。其中每个字母表示一种不同种类的任务。任务可以以任意顺序执行，并且每个任务都可以在 1 个单位时间内执行完。在任何一个单位时间，CPU 可以完成一个任务，或者处于待命状态。

然而，两个 相同种类 的任务之间必须有长度为整数 n 的冷却时间，因此至少有连续 n 个单位时间内 CPU 在执行不同的任务，或者在待命状态。

你需要计算完成所有任务所需要的 最短时间 。



示例 1：

输入：tasks = ["A","A","A","B","B","B"], n = 2
输出：8
解释：A -> B -> (待命) -> A -> B -> (待命) -> A -> B
     在本示例中，两个相同类型任务之间必须间隔长度为 n = 2 的冷却时间，而执行一个任务只需要一个单位时间，所以中间出现了（待命）状态。
示例 2：

输入：tasks = ["A","A","A","B","B","B"], n = 0
输出：6
解释：在这种情况下，任何大小为 6 的排列都可以满足要求，因为 n = 0
["A","A","A","B","B","B"]
["A","B","A","B","A","B"]
["B","B","B","A","A","A"]
...
诸如此类
示例 3：

输入：tasks = ["A","A","A","A","A","A","B","C","D","E","F","G"], n = 2
输出：16
解释：一种可能的解决方案是：
     A -> B -> C -> A -> D -> E -> A -> F -> G -> A -> (待命) -> (待命) -> A -> (待命) -> (待命) -> A


提示：

1 <= task.length <= 104
tasks[i] 是大写英文字母
n 的取值范围为 [0, 100]
*/

/**
经过思考：
基本原则，每次取允许执行的剩余最多的任务来执行

1、如何判断是否允许执行？使用长度为n的数组来保存？（查找线性）- 改使用 - map， key为任务值，val为最近执行的顺序下标， 通过下标就可以判断是否允许执行 Y
2、如何获取剩余最多的任务？map保存所有任务值对应的剩余数目（还是只能线性查找）- 改使用map， key为剩余的数目，val为任务列表 X

*/

func leastInterval(tasks []byte, n int) int {

	countTasks := map[int][]byte{}
	maxCount := 0
	total := len(tasks)

	taskCount := map[byte]int{}
	for _, task := range tasks {
		taskCount[task]++
	}
	// X
	for task, count := range taskCount {
		if count > maxCount {
			maxCount = count
		}
		if tasks, ok := countTasks[count]; ok {
			tasks = append(tasks, task)
			countTasks[count] = tasks
		} else {
			countTasks[count] = []byte{task}
		}
	}

	result := 0
	// Y
	latestTaskIndex := map[byte]int{}
	for total > 0 {

		for i := maxCount; i >= 1; i-- {
			if tasks, ok := countTasks[i]; ok {
				for _, task := range tasks {
					if index := latestTaskIndex[task]; index == 0 || result-latestTaskIndex[task] >= n {
						result++
						//countTasks[i] =

						//写得太复杂了， 不想写了！！！
					}
				}
			}
		}

	}

	return result
}

func TestLeastInterval(t *testing.T) {
	leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2)
}
