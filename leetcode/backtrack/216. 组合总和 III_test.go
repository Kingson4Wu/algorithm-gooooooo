package backtrack

import (
	"fmt"
	"testing"
)

/*
*
找出所有相加之和为 n 的 k 个数的组合，且满足下列条件：

只使用数字1到9
每个数字 最多使用一次
返回 所有可能的有效组合的列表 。该列表不能包含相同的组合两次，组合可以以任何顺序返回。

示例 1:

输入: k = 3, n = 7
输出: [[1,2,4]]
解释:
1 + 2 + 4 = 7
没有其他符合的组合了。
示例 2:

输入: k = 3, n = 9
输出: [[1,2,6], [1,3,5], [2,3,4]]
解释:
1 + 2 + 6 = 9
1 + 3 + 5 = 9
2 + 3 + 4 = 9
没有其他符合的组合了。
示例 3:

输入: k = 4, n = 1
输出: []
解释: 不存在有效的组合。
在[1,9]范围内使用4个不同的数字，我们可以得到的最小和是1+2+3+4 = 10，因为10 > 1，没有有效的组合。

提示:

2 <= k <= 9
1 <= n <= 60
*/
func combinationSum3(k int, n int) [][]int {

	var ans [][]int
	var path []int

	var backtrack func(start, count, sum int)
	backtrack = func(start, count, sum int) {

		if sum == n && count == k {
			temp := make([]int, k)
			copy(temp, path)
			ans = append(ans, temp)
			return
		}

		for i := start; i <= 9; i++ {
			if sum+i > n {
				break
			}
			if sum+i < n && count >= k {
				break
			}

			path = append(path, i)
			backtrack(i+1, count+1, sum+i)
			path = path[:len(path)-1]
		}
	}
	backtrack(1, 0, 0)
	return ans
}

/**
func combinationSum3(k int, n int) [][]int {
    var res [][]int
    var path []int

    var backtrack func(start, sum int)
    backtrack = func(start, sum int) {
        if len(path) == k {
            if sum == n {
                temp := make([]int, k)
                copy(temp, path)
                res = append(res, temp)
            }
            return
        }

        for i := start; i <= 9; i++ {
            if sum + i > n {
                break // 剪枝
            }

            path = append(path, i)
            backtrack(i + 1, sum + i) // i+1，保证数字不重复
            path = path[:len(path)-1]
        }
    }

    backtrack(1, 0)
    return res
}

*/

func TestCombinationSum3(t *testing.T) {
	fmt.Println(combinationSum3(3, 7))
	fmt.Println(combinationSum3(3, 9))
	fmt.Println(combinationSum3(4, 1))
}

/**
自己模仿前两道题做出来了

执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
3.73
MB
击败
71.68%
复杂度分析
*/
