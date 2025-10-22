package _02510

import (
	"fmt"
	"testing"
)

/*
*
定义：f(t, k) 表示在最多尝试 t 次、拥有 k 个鸡蛋的情况下，最多能测试的楼层数

转移方程：f(t, k) = 1 + f(t-1, k-1) + f(t-1, k)
最终寻找最小的 t，使得 f(t, k) >= n

执行用时分布
15
ms
击败
45.12%
复杂度分析
消耗内存分布
11.97
MB
击败
18.29%

双重循环应该从2开始，k和t为1的场景一开始已经初始化！！
*/
func superEggDrop(k int, n int) int {

	f := make([][]int, n+1)
	for t := 0; t <= n; t++ {
		f[t] = make([]int, k+1)
	}
	for t := 0; t <= n; t++ {
		f[t][1] = t
	}
	for kk := 0; kk <= k; kk++ {
		f[1][kk] = 1
	}
	ans := n
	if f[1][k] >= n {
		ans = 1
	}
	for t := 2; t <= n; t++ {
		for kk := 2; kk <= k; kk++ {
			f[t][kk] = 1 + f[t-1][kk] + f[t-1][kk-1]
		}
		if f[t][k] >= n {
			ans = t
			break
		}
	}
	return ans
}

/**
解答错误
90 / 121 个通过的测试用例

官方题解
输入
k =
1
n =
3

添加到测试用例
输出
2
预期结果
3
*/

func TestSuperEggDrop(t *testing.T) {
	fmt.Println(superEggDrop(1, 3))
}
