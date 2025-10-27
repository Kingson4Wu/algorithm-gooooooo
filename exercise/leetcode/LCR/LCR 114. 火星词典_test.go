package LCR

import (
	"fmt"
	"strings"
	"testing"
)

/*
*
个人想法
有向图
双重循环构建图
有向图的拓扑排序算法

明明自己写对了，还要加上
输入
words =
["abc","ab"]

添加到测试用例
输出
"abc"
预期结果
""

这种恶心测试用例。。。
*/
func alienOrder(words []string) string {

	graph := make(map[string]struct{})

	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			for k := 0; k < len(words[i]) && k < len(words[j]); k++ {
				if words[i][k] == words[j][k] {
					continue
				}
				graph[string(words[i][k])+"_"+string(words[j][k])] = struct{}{}
				break
			}
		}
	}
	inDegree := make(map[string]int)
	for i := 0; i < len(words); i++ {
		for k := 0; k < len(words[i]); k++ {
			inDegree[string(words[i][k])] = 0
		}
	}

	neighborhood := make(map[string][]string)
	for g := range graph {
		m := strings.Split(g, "_")
		a := m[0]
		b := m[1]
		inDegree[b]++
		if arr, ok := neighborhood[a]; ok {
			neighborhood[a] = append(arr, b)
		} else {
			neighborhood[a] = []string{b}
		}
	}
	var queue []string
	for k, v := range inDegree {
		if v == 0 {
			queue = append(queue, k)
		}
	}
	var result []string
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		result = append(result, v)
		for _, neigh := range neighborhood[v] {
			inDegree[neigh]--
			if inDegree[neigh] == 0 {
				queue = append(queue, neigh)
			}
		}
	}
	if len(result) != len(inDegree) {
		return ""
	}
	return strings.Join(result, "")
}

func TestAlienOrder(t *testing.T) {
	fmt.Println(alienOrder([]string{"wrt", "wrf", "er", "ett", "rftt"}))
	fmt.Println(alienOrder([]string{"z", "x"}))
	fmt.Println(alienOrder([]string{"z", "x", "z"}))
	fmt.Println(alienOrder([]string{"z", "z"}))
	fmt.Println(alienOrder([]string{"abc", "ab"}))
	//妈的垃圾题目！！！垃圾测试用例
	fmt.Println(alienOrder([]string{"zy", "zx"}))
}

/**
输入
words =
["abc","ab"]

添加到测试用例
输出
"abc"
预期结果
""


解答错误
81 / 122 个通过的测试用例

官方题解
输入
words =
["zy","zx"]

添加到测试用例
输出
""
预期结果
"yxz"
*/

/**
这测试用例真的有病：
输入
words =
["abc","ab"]

添加到测试用例
输出
"abc"
预期结果
""

输入
words =
["ab","ab","abc"]
输出
""
预期结果
"abc"
*/
