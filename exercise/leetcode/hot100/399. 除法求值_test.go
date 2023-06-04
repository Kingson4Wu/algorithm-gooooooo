package hot100

import (
	"fmt"
	"testing"
)

/*
*
给你一个变量对数组 equations 和一个实数值数组 values 作为已知条件，其中 equations[i] = [Ai, Bi] 和 values[i] 共同表示等式 Ai / Bi = values[i] 。每个 Ai 或 Bi 是一个表示单个变量的字符串。

另有一些以数组 queries 表示的问题，其中 queries[j] = [Cj, Dj] 表示第 j 个问题，请你根据已知条件找出 Cj / Dj = ? 的结果作为答案。

返回 所有问题的答案 。如果存在某个无法确定的答案，则用 -1.0 替代这个答案。如果问题中出现了给定的已知条件中没有出现的字符串，也需要用 -1.0 替代这个答案。

注意：输入总是有效的。你可以假设除法运算中不会出现除数为 0 的情况，且不存在任何矛盾的结果。

示例 1：

输入：equations = [["a","b"],["b","c"]], values = [2.0,3.0], queries = [["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]
输出：[6.00000,0.50000,-1.00000,1.00000,-1.00000]
解释：
条件：a / b = 2.0, b / c = 3.0
问题：a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ?
结果：[6.0, 0.5, -1.0, 1.0, -1.0 ]
示例 2：

输入：equations = [["a","b"],["b","c"],["bc","cd"]], values = [1.5,2.5,5.0], queries = [["a","c"],["c","b"],["bc","cd"],["cd","bc"]]
输出：[3.75000,0.40000,5.00000,0.20000]
示例 3：

输入：equations = [["a","b"]], values = [0.5], queries = [["a","b"],["b","a"],["a","c"],["x","y"]]
输出：[0.50000,2.00000,-1.00000,-1.00000]
*/
/**
想了一下，只能穷举，为了避免重复计算，保存计算过的结果，即动态规划
寻找的过程个人写得好像很复杂. 写了一堆shi, 明明有思路，却写不对
*/
/**
涉及图
*/
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {

	dp := map[string]bool{}
	directions := map[string]map[string]float64{}

	/** 初始化 */
	for i := 0; i < len(equations); i++ {
		a := equations[i][0]
		b := equations[i][1]

		dp[a+"-"+b] = true
		dp[b+"-"+a] = true

		if direction, ok := directions[a]; ok {
			direction[b] = values[i]
		} else {
			direction := map[string]float64{}
			direction[b] = values[i]
			directions[a] = direction
		}

		if direction, ok := directions[b]; ok {
			direction[a] = 1 / values[i]
		} else {
			direction := map[string]float64{}
			direction[a] = 1 / values[i]
			directions[b] = direction
		}

	}

	// a-b b-f f-d
	// a-d

	var seek func(direction map[string]float64, sourceV float64, source, target string)
	seek = func(direction map[string]float64, sourceV float64, source, target string) {
		if len(direction) == 0 {
			return
		}
		if _, ok := dp[source+"-"+target]; ok {
			return
		}

		if v, ok := direction[target]; ok {
			dp[source+"-"+target] = true
			dp[target+"-"+source] = true
			directions[source][target] = sourceV * v
			directions[target][source] = 1 / (sourceV * v)
			return
		}
		for k, v := range direction {
			if d, ok := directions[k]; ok {
				seek(d, v, k, target)
			}
		}
		if v, ok := direction[target]; ok {
			dp[source+"-"+target] = true
			dp[target+"-"+source] = true
			directions[source][target] = sourceV * v
			directions[target][source] = 1 / (sourceV * v)
			return
		}
	}

	results := make([]float64, len(queries))
	for i := 0; i < len(queries); i++ {
		a := queries[i][0]
		b := queries[i][1]

		if _, ok := directions[a]; !ok {
			results[i] = -1
			continue
		}
		if _, ok := directions[b]; !ok {
			results[i] = -1
			continue
		}
		if a == b {
			results[i] = 1
			continue
		}

		if _, ok := dp[a+"-"+b]; ok {
			results[i] = directions[a][b]
			continue
		}

		direction := directions[a]

		for k, v := range direction {
			if d, ok := directions[k]; ok {
				seek(d, v, a, b)
			}
		}

		if _, ok := dp[a+"-"+b]; ok {
			results[i] = directions[a][b]
			continue
		}
		results[i] = -1
	}

	return results
}

func TestCalcEquation(t *testing.T) {
	//fmt.Println(calcEquation([][]string{{"a", "b"}, {"b", "c"}}, []float64{2.0, 3.0}, [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}))
	//fmt.Println(calcEquation([][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}}, []float64{1.5, 2.5, 5.0}, [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}}))
	//fmt.Println(calcEquation([][]string{{"a", "b"}}, []float64{0.5}, [][]string{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"x", "y"}}))
	fmt.Println(calcEquation([][]string{{"x1", "x2"}, {"x2", "x3"}, {"x3", "x4"}, {"x4", "x5"}}, []float64{3.0, 4.0, 5.0, 6.0}, [][]string{{"x1", "x5"}, {"x5", "x2"}, {"x2", "x4"}, {"x2", "x2"}, {"x2", "x9"}, {"x9", "x9"}}))
}

/**
runtime: out of memory: cannot allocate 268435456-byte block (138149888 in use)
fatal error: out of memory

equations =
[["x1","x2"],["x2","x3"],["x3","x4"],["x4","x5"]]
values =
[3.0,4.0,5.0,6.0]
queries =
[["x1","x5"],["x5","x2"],["x2","x4"],["x2","x2"],["x2","x9"],["x9","x9"]]
*/
