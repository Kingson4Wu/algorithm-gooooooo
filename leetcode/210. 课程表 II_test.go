package leetcode

import (
	"fmt"
	"testing"
)

/*
*
现在你总共有 numCourses 门课需要选，记为 0 到 numCourses - 1。给你一个数组 prerequisites ，其中 prerequisites[i] = [ai, bi] ，表示在选修课程 ai 前 必须 先选修 bi 。

例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示：[0,1] 。
返回你为了学完所有课程所安排的学习顺序。可能会有多个正确的顺序，你只要返回 任意一种 就可以了。如果不可能完成所有课程，返回 一个空数组 。

示例 1：

输入：numCourses = 2, prerequisites = [[1,0]]
输出：[0,1]
解释：总共有 2 门课程。要学习课程 1，你需要先完成课程 0。因此，正确的课程顺序为 [0,1] 。
示例 2：

输入：numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
输出：[0,2,1,3]
解释：总共有 4 门课程。要学习课程 3，你应该先完成课程 1 和课程 2。并且课程 1 和课程 2 都应该排在课程 0 之后。
因此，一个正确的课程顺序是 [0,1,2,3] 。另一个正确的排序是 [0,2,1,3] 。
示例 3：

输入：numCourses = 1, prerequisites = []
输出：[0]

提示：
1 <= numCourses <= 2000
0 <= prerequisites.length <= numCourses * (numCourses - 1)
prerequisites[i].length == 2
0 <= ai, bi < numCourses
ai != bi
所有[ai, bi] 互不相同
*/

/*
*
拓扑排序，跟207. 课程表 基本一样
卡恩算法，判断是否有环
凭几天前的回忆完成

执行用时分布
4
ms
击败
25.84%
复杂度分析
消耗内存分布
7.97
MB
击败
19.89%
复杂度分析
*/
func findOrder(numCourses int, prerequisites [][]int) []int {
	inDegree := make(map[int]int)
	neighbourhoods := make(map[int][]int)
	for i := 0; i < numCourses; i++ {
		inDegree[i] = 0
	}
	var ans []int
	for i := 0; i < len(prerequisites); i++ {
		inDegree[prerequisites[i][0]]++
		if list, ok := neighbourhoods[prerequisites[i][1]]; ok {
			neighbourhoods[prerequisites[i][1]] = append(list, prerequisites[i][0])
		} else {
			neighbourhoods[prerequisites[i][1]] = []int{prerequisites[i][0]}
		}
	}
	var queue []int
	for course, num := range inDegree {
		if num == 0 {
			queue = append(queue, course)
		}
	}
	for len(queue) > 0 {
		course := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		ans = append(ans, course)
		for _, neighbour := range neighbourhoods[course] {
			inDegree[neighbour]--
			if inDegree[neighbour] == 0 {
				queue = append(queue, neighbour)
			}
		}
	}
	if len(ans) < numCourses {
		return []int{}
	}
	return ans
}

func TestFindOrder(t *testing.T) {
	fmt.Println(findOrder(2, [][]int{{1, 0}}))
	fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
	fmt.Println(findOrder(1, [][]int{}))
}
