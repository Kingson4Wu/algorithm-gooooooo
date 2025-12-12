package LCR

import (
	"fmt"
	"testing"
)

/*
*
按以往的经验自己做，不过要靠调试调整,而且写得很乱。。。
最后还是做错了。。。
*/
func findCircleNum(isConnected [][]int) int {
	parents := make([]int, len(isConnected))
	union := func(i, j int) {
		if isConnected[i][j] == 1 {
			if parents[j] > i {
				parents[j] = i
				for parents[i] != i {
					parents[j] = parents[i]
					i = parents[i]
				}
			} else {
				k := parents[i]
				parents[i] = parents[j]
				for parents[k] != parents[j] {
					parents[k], k = parents[j], parents[k]
				}
			}

		}
	}
	for i := 0; i < len(isConnected); i++ {
		parents[i] = i
	}
	for i := 0; i < len(isConnected); i++ {
		for j := i; j < len(isConnected); j++ {
			if i == j {
				continue
			}
			union(i, j)
		}
	}
	m := make(map[int]bool)
	count := 0
	for _, p := range parents {
		if !m[p] {
			count++
			m[p] = true
		}
	}
	return count
}

func TestFindCircleNum(t *testing.T) {
	fmt.Println(findCircleNum([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}))
	fmt.Println(findCircleNum([][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}))
	fmt.Println(findCircleNum([][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 1, 1}}))
}

/**
解答错误
66 / 113 个通过的测试用例

官方题解
输入
isConnected =
[[1,0,0,1],[0,1,1,0],[0,1,1,1],[1,0,1,1]]

添加到测试用例
输出
2
预期结果
1
*/
