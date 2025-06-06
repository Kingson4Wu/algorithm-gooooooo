package tencent50

/*
*
n 位格雷码序列 是一个由 2n 个整数组成的序列，其中：
每个整数都在范围 [0, 2n - 1] 内（含 0 和 2n - 1）
第一个整数是 0
一个整数在序列中出现 不超过一次
每对 相邻 整数的二进制表示 恰好一位不同 ，且
第一个 和 最后一个 整数的二进制表示 恰好一位不同
给你一个整数 n ，返回任一有效的 n 位格雷码序列 。

示例 1：

输入：n = 2
输出：[0,1,3,2]
解释：
[0,1,3,2] 的二进制表示是 [00,01,11,10] 。
- 00 和 01 有一位不同
- 01 和 11 有一位不同
- 11 和 10 有一位不同
- 10 和 00 有一位不同
[0,2,3,1] 也是一个有效的格雷码序列，其二进制表示是 [00,10,11,01] 。
- 00 和 10 有一位不同
- 10 和 11 有一位不同
- 11 和 01 有一位不同
- 01 和 00 有一位不同
示例 2：

输入：n = 1
输出：[0,1]

提示：

1 <= n <= 16
*/

/*
*
题意重点：
每对 相邻 整数的二进制表示 恰好一位不同 ，且
第一个 和 最后一个 整数的二进制表示 恰好一位不同
*/

/*
*
方法一：归纳法

	func grayCode(n int) []int {
	    ans := make([]int, 1, 1<<n)
	    for i := 1; i <= n; i++ {
	        for j := len(ans) - 1; j >= 0; j-- {
	            ans = append(ans, ans[j]|1<<(i-1))
	        }
	    }
	    return ans
	}

作者：力扣官方题解
链接：https://leetcode.cn/problems/gray-code/solutions/1196538/ge-lei-bian-ma-by-leetcode-solution-cqi7/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

方法二：公式法

	func grayCode(n int) []int {
	    ans := make([]int, 1<<n)
	    for i := range ans {
	        ans[i] = i>>1 ^ i
	    }
	    return ans
	}

作者：力扣官方题解
链接：https://leetcode.cn/problems/gray-code/solutions/1196538/ge-lei-bian-ma-by-leetcode-solution-cqi7/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func grayCode(n int) []int {

	return []int{}
}
