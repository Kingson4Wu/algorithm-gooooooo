package ugly

import (
	"fmt"
	"math"
	"testing"
)

/*
*
超级丑数 是一个正整数，并满足其所有质因数都出现在质数数组 primes 中。

给你一个整数 n 和一个整数数组 primes ，返回第 n 个 超级丑数 。

题目数据保证第 n 个 超级丑数 在 32-bit 带符号整数范围内。

示例 1：

输入：n = 12, primes = [2,7,13,19]
输出：32
解释：给定长度为 4 的质数数组 primes = [2,7,13,19]，前 12 个超级丑数序列为：[1,2,4,7,8,13,14,16,19,26,28,32] 。
示例 2：

输入：n = 1, primes = [2,3,5]
输出：1
解释：1 不含质因数，因此它的所有质因数都在质数数组 primes = [2,3,5] 中。

提示：

1 <= n <= 105
1 <= primes.length <= 100
2 <= primes[i] <= 1000
题目数据 保证 primes[i] 是一个质数
primes 中的所有值都 互不相同 ，且按 递增顺序 排列
*/
/**
根据回忆写，注意：
1、是indexes[mJ]++
2、结果都是min的，都要++
*/
func nthSuperUglyNumber(n int, primes []int) int {

	result := make([]int, n)
	result[0] = 1
	indexes := make([]int, n)
	for i := 1; i < n; i++ {
		min := math.MaxInt32
		var minJ []int
		for j, prime := range primes {
			val := result[indexes[j]] * prime
			if val < min {
				min = val
				minJ = []int{j}
			} else if val == min {
				minJ = append(minJ, j)
			}
		}
		result[i] = min
		//indexes[minJ] = i
		for _, mJ := range minJ {
			indexes[mJ]++
		}
	}
	return result[n-1]
}

/*
*
panic: runtime error: index out of range [2] with length 2
main.nthSuperUglyNumber(0x2, {0xc000088040, 0x3, 0x4cd580?})
solution.go, line 10
main.__helper__(...)
solution.go, line 30
main.main()
solution.go, line 71
*/
func TestNthSuperUglyNumber(t *testing.T) {
	fmt.Println(nthSuperUglyNumber(12, []int{2, 7, 13, 19}))

	fmt.Println(nthSuperUglyNumber(2, []int{2, 3, 5}))
}

/**
这道题和「264. 丑数 II」相似，区别在于，第 264 题规定丑数是只包含质因数 2、3 和 5 的正整数，这道题规定超级丑数是只包含数组 primes 中的质因数的正整数。

第 264 题的方法包括最小堆和动态规划。由于这道题的数据规模较大，因此最小堆的解法会超时，此处只提供动态规划的解法。

func nthSuperUglyNumber(n int, primes []int) int {
    dp := make([]int, n+1)
    m := len(primes)
    pointers := make([]int, m)
    nums := make([]int, m)
    for i := range nums {
        nums[i] = 1
    }
    for i := 1; i <= n; i++ {
        minNum := math.MaxInt64
        for j := range pointers {
            minNum = min(minNum, nums[j])
        }
        dp[i] = minNum
        for j := range nums {
            if nums[j] == minNum {
                pointers[j]++
                nums[j] = dp[pointers[j]] * primes[j]
            }
        }
    }
    return dp[n]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/super-ugly-number/solutions/924207/chao-ji-chou-shu-by-leetcode-solution-uzff/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
