package topinterview150

import (
	"fmt"
	"testing"
)

/**

给你两个整数 left 和 right ，表示区间 [left, right] ，返回此区间内所有数字 按位与 的结果（包含 left 、right 端点）。



示例 1：

输入：left = 5, right = 7
输出：4
示例 2：

输入：left = 0, right = 0
输出：0
示例 3：

输入：left = 1, right = 2147483647
输出：0


提示：

0 <= left <= right <= 231 - 1
*/
/**
自己想的
因为与完是0的永远为0
从最位开始与，高位为0，则将循环变量值左移，看是否在区间内，以加速遍历效率

能想个大概，但没做出来

*/
/*func rangeBitwiseAnd(left int, right int) int {

	if right == 0 {
		return 0
	}
	val := left
	//低位有多少个0
	n := 1

	for i := left + 1; i <= right; i-- {
		val &= i
		v := val & n
		for {
			if v&1 == 1 {
				break
			}
			n <<= 1
			v >>= 1
		}
		i = n
	}

	return val
}*/

/*
*
方法一：位移
目的是求出两个给定数字的二进制字符串的公共前缀
将两个数字不断向右移动，直到数字相等，即数字被缩减为它们的公共前缀。然后，通过将公共前缀向左移动，将零添加到公共前缀的右边以获得最终结果。

执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
6.55
MB
击败
46.30%
*/
func rangeBitwiseAnd(left int, right int) int {
	count := 0
	for left != right {
		left >>= 1
		right >>= 1
		count++
	}
	return left << count
}
func TestRangeBitwiseAnd(t *testing.T) {
	fmt.Println(rangeBitwiseAnd(5, 7))
}
