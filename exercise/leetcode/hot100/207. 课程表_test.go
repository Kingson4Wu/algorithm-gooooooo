package hot100

import (
	"fmt"
	"testing"
)

/*
*
leetcode/graph/207. 课程表.go
*/
/**
有向图检查是否有环
拓扑排序，检查
*/
func canFinish(numCourses int, prerequisites [][]int) bool {

	inDegree := make(map[int]int)
	neighbourhood := make(map[int][]int)
	for i := 0; i < len(prerequisites); i++ {
		if _, ok := inDegree[prerequisites[i][0]]; !ok {
			inDegree[prerequisites[i][0]] = 0
		}
		inDegree[prerequisites[i][1]]++
		if _, ok := neighbourhood[prerequisites[i][0]]; !ok {
			neighbourhood[prerequisites[i][0]] = []int{prerequisites[i][1]}
		} else {
			neighbourhood[prerequisites[i][0]] = append(neighbourhood[prerequisites[i][0]], prerequisites[i][1])
		}
	}
	var ans []int
	var queue []int
	for k, v := range inDegree {
		if v == 0 {
			queue = append(queue, k)
		}
	}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		ans = append(ans, node)
		for _, neighbour := range neighbourhood[node] {
			inDegree[neighbour]--
			if inDegree[neighbour] == 0 {
				queue = append(queue, neighbour)
			}
		}
	}
	return len(ans) == numCourses
}

func TestCanFinish(t *testing.T) {
	fmt.Println(canFinish(2, [][]int{{1, 0}}))
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
	fmt.Println(canFinish(5, [][]int{{1, 0}, {0, 2}, {0, 3}, {3, 4}}))
	fmt.Println(canFinish(3, [][]int{{1, 2}, {2, 3}}))
	fmt.Println(canFinish(5, [][]int{{1, 0}, {0, 2}, {2, 4}, {0, 3}, {3, 4}}))
	fmt.Println(canFinish(4, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}}))
	fmt.Println(canFinish(1, [][]int{}))
}

/**
输入
numCourses =
1
prerequisites =
[]

添加到测试用例
输出
false
预期结果
true
*/
