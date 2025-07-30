package graph

import (
	"fmt"
	"testing"
)

/*
*
自己写
*/
func canFinish(numCourses int, prerequisites [][]int) bool {
	inDegrees := make([]int, numCourses, numCourses)
	neighbors := make([][]int, numCourses, numCourses)
	for _, pre := range prerequisites {
		inDegrees[pre[0]]++
		if len(neighbors[pre[1]]) == 0 {
			neighbors[pre[1]] = []int{pre[0]}
		} else {
			neighbors[pre[1]] = append(neighbors[pre[1]], pre[0])
		}
	}
	var queue []int
	for i, inDegree := range inDegrees {
		if inDegree == 0 {
			queue = append(queue, i)
		}
	}
	count := 0
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		count++
		for _, neighbor := range neighbors[course] {
			inDegrees[neighbor]--
			if inDegrees[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}
	return count == numCourses
}

func TestCanFinish(t *testing.T) {
	fmt.Println(canFinish(2, [][]int{{1, 0}}))
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
	fmt.Println(canFinish(5, [][]int{{1, 0}, {0, 2}, {0, 3}, {3, 4}}))
	fmt.Println(canFinish(3, [][]int{{0, 1}, {1, 2}}))
	fmt.Println(canFinish(5, [][]int{{1, 0}, {0, 2}, {2, 4}, {0, 3}, {3, 4}}))
	fmt.Println(canFinish(4, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}}))

	/**
	numCourses = 1
	prerequisites = []
	含义是：你有一门课程要上，没有任何先修课要求。
	*/
	fmt.Println(canFinish(1, [][]int{}))
}

/**
true
false
true
true
true
false
true
*/
