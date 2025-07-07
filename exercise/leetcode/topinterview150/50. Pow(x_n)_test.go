package topinterview150

import (
	"fmt"
	"testing"
)

/**
实现 pow(x, n) ，即计算 x 的整数 n 次幂函数（即，xn ）。

示例 1：

输入：x = 2.00000, n = 10
输出：1024.00000
示例 2：

输入：x = 2.10000, n = 3
输出：9.26100
示例 3：

输入：x = 2.00000, n = -2
输出：0.25000
解释：2-2 = 1/22 = 1/4 = 0.25


提示：

-100.0 < x < 100.0
-231 <= n <= 231-1
n 是一个整数
要么 x 不为零，要么 n > 0 。
-104 <= xn <= 104
*/
/**
自己做的
2的指数
从1开始乘2直到大于n，结果保存到map中，方便后续复用
从小于n的最大2指数一直除以2

*/
func myPow(x float64, n int) float64 {

	if x == 0 {
		return 0
	}
	if x == 1 || n == 0 {
		return 1
	}
	sign := true
	if n < 0 {
		sign = false
		n = -n
	}
	m := make(map[int]float64)
	m[1] = x
	i := 2
	v := x
	for i <= n {
		v *= v
		m[i] = v
		i <<= 1
	}
	i >>= 1
	r := n - i

	for r > 0 {
		for i > r {
			i >>= 1
		}
		v *= m[i]
		r -= i
	}
	if sign {
		return v
	}
	return 1 / v
}

/**
解答错误
298 / 307 个通过的测试用例

官方题解
输入
x =
0.97267
n =
13

添加到测试用例
输出
0.67845
预期结果
0.69751

执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
3.83
MB
击败
48.45%
*/

func TestMyPow(t *testing.T) {
	//fmt.Println(myPow(2.1, 1))
	//fmt.Println(myPow(2, 10))
	//fmt.Println(myPow(2.1, 3))
	//fmt.Println(myPow(2, -2))
	fmt.Println(myPow(0.97267, 13))
}
