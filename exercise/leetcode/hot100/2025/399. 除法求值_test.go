package _025

import (
	"fmt"
	"testing"
)

/*
*
个人多次尝试完成
有向图遍历
双向保存除法的分母分子映射以及对应的数据， val、 1/val， 保存在同个map， 另外一个对应的val map作为计算辅助
递归有向图遍历，visited标记， 回溯维护val列表值（全局变量）

	vals = append(vals, chValueChain[a][j])
	if ch == b {
		var r float64 = 1
		for i := 0; i < len(vals); i++ {
			r *= vals[i]
		}
		result = r
	} else {
		dfs(ch, b)
	}
	vals = vals[:len(vals)-1]
*/
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {

	chSet := make(map[string]bool)
	chChain := make(map[string][]string)
	chValueChain := make(map[string][]float64)

	for i := 0; i < len(equations); i++ {
		chSet[equations[i][0]] = true
		chSet[equations[i][1]] = true
		if l, ok := chChain[equations[i][0]]; ok {
			chChain[equations[i][0]] = append(l, equations[i][1])
		} else {
			chChain[equations[i][0]] = []string{equations[i][1]}
		}
		if l, ok := chValueChain[equations[i][0]]; ok {
			chValueChain[equations[i][0]] = append(l, values[i])
		} else {
			chValueChain[equations[i][0]] = []float64{values[i]}
		}

		if l, ok := chChain[equations[i][1]]; ok {
			chChain[equations[i][1]] = append(l, equations[i][0])
		} else {
			chChain[equations[i][1]] = []string{equations[i][0]}
		}
		if l, ok := chValueChain[equations[i][1]]; ok {
			chValueChain[equations[i][1]] = append(l, 1/values[i])
		} else {
			chValueChain[equations[i][1]] = []float64{1 / values[i]}
		}
	}

	getValue := func(k, v string) float64 {
		visited := make(map[string]bool)
		var vals []float64
		var result float64 = -1
		var dfs func(a, b string)
		dfs = func(a, b string) {
			if l, ok := chChain[a]; ok {
				for j, ch := range l {
					if visited[a+"_"+ch] {
						continue
					}
					visited[a+"_"+ch] = true
					vals = append(vals, chValueChain[a][j])
					if ch == b {
						var r float64 = 1
						for i := 0; i < len(vals); i++ {
							r *= vals[i]
						}
						result = r
					} else {
						dfs(ch, b)
					}
					vals = vals[:len(vals)-1]
				}
			}
		}
		dfs(k, v)
		return result
	}

	ans := make([]float64, len(queries))
	for i := 0; i < len(queries); i++ {
		if !chSet[queries[i][0]] || !chSet[queries[i][1]] {
			ans[i] = -1
			continue
		}
		if queries[i][0] == queries[i][1] {
			ans[i] = 1
			continue
		}
		val := getValue(queries[i][0], queries[i][1])
		ans[i] = val
	}
	return ans
}

/**
解答错误
15 / 29 个通过的测试用例

官方题解
输入
equations =
[["a","e"],["b","e"]]
values =
[4.0,3.0]
queries =
[["a","b"],["e","e"],["x","x"]]

添加到测试用例
输出
[-1.00000,1.00000,-1.00000]
预期结果
[1.33333,1.0,-1.0]


解答错误
26 / 29 个通过的测试用例

官方题解
输入
equations =
[["a","b"],["b","c"],["a","c"]]
values =
[2.0,3.0,6.0]
queries =
[["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]

添加到测试用例
输出
[3.00000,0.50000,-1.00000,1.00000,-1.00000]
预期结果
[6.0,0.5,-1.0,1.0,-1.0]

执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
4.14
MB
击败
5.73%
*/

func TestCalcEquation(t *testing.T) {

	fmt.Println(calcEquation([][]string{{"a", "b"}, {"b", "c"}, {"a", "c"}}, []float64{2.0, 3.0, 6.0}, [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}))

	fmt.Println(calcEquation([][]string{{"a", "e"}, {"b", "e"}}, []float64{4.0, 3.0}, [][]string{{"a", "b"}, {"e", "e"}, {"x", "x"}}))

	fmt.Println(calcEquation([][]string{{"a", "b"}, {"b", "c"}}, []float64{2.0, 3.0}, [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}))
	fmt.Println(calcEquation([][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}}, []float64{1.5, 2.5, 5.0}, [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}}))
	fmt.Println(calcEquation([][]string{{"a", "b"}}, []float64{0.5}, [][]string{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"x", "y"}}))
	fmt.Println(calcEquation([][]string{{"x1", "x2"}, {"x2", "x3"}, {"x3", "x4"}, {"x4", "x5"}}, []float64{3.0, 4.0, 5.0, 6.0}, [][]string{{"x1", "x5"}, {"x5", "x2"}, {"x2", "x4"}, {"x2", "x2"}, {"x2", "x9"}, {"x9", "x9"}}))
}

/**
[6 0.5 -1 1 -1]
[1.3333333333333333 1 -1]
[6 0.5 -1 1 -1]
[3.75 0.4 5 0.2]
[0.5 2 -1 -1]
[360 0.008333333333333333 20 1 -1 -1]


*/
