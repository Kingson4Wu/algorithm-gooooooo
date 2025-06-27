package topinterview150

import (
	"fmt"
	"testing"
)

/**
自己总结规律完成的
*/
/**
执行用时分布
12
ms
击败
6.23%
复杂度分析
消耗内存分布
3.87
MB
击败
27.87%
*/
/*
*
给定一个整数 n ，返回 n! 结果中尾随零的数量。

提示 n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1

示例 1：

输入：n = 3
输出：0
解释：3! = 6 ，不含尾随 0
示例 2：

输入：n = 5
输出：1
解释：5! = 120 ，有一个尾随 0
示例 3：

输入：n = 0
输出：0
*/
func trailingZeroes(n int) int {

	count := 0
	two := 0
	five := 0

	for i := 0; i <= n; i++ {
		num := i
		for num%10 == 0 && num/10 > 0 {
			count++
			num /= 10
		}
		for num%2 == 0 && num/2 > 0 {
			two++
			num /= 2
		}
		for num%5 == 0 && num/5 > 0 {
			five++
			num /= 5
		}
		for two > 0 && five > 0 {
			count++
			two--
			five--
		}
	}

	return count
}

func TestTrailingZeroes(t *testing.T) {
	//fmt.Println(trailingZeroes(0))
	//fmt.Println(trailingZeroes(3))
	fmt.Println(trailingZeroes(5))
}

/**
func trailingZeroes(n int) (ans int) {
    for n > 0 {
        n /= 5
        ans += n
    }
    return
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/factorial-trailing-zeroes/solutions/1360892/jie-cheng-hou-de-ling-by-leetcode-soluti-1egk/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
