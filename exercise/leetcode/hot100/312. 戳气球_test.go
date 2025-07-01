package hot100

import (
	"fmt"
	"testing"
)

/**
有 n 个气球，编号为0 到 n - 1，每个气球上都标有一个数字，这些数字存在数组 nums 中。

现在要求你戳破所有的气球。戳破第 i 个气球，你可以获得 nums[i - 1] * nums[i] * nums[i + 1] 枚硬币。 这里的 i - 1 和 i + 1 代表和 i 相邻的两个气球的序号。如果 i - 1或 i + 1 超出了数组的边界，那么就当它是一个数字为 1 的气球。

求所能获得硬币的最大数量。



示例 1：
输入：nums = [3,1,5,8]
输出：167
解释：
nums = [3,1,5,8] --> [3,5,8] --> [3,8] --> [8] --> []
coins =  3*1*5    +   3*5*8   +  1*3*8  + 1*8*1 = 167
示例 2：

输入：nums = [1,5]
输出：10


提示：

n == nums.length
1 <= n <= 300
0 <= nums[i] <= 100
*/

/**

看题解：https://zhuanlan.zhihu.com/p/144384951

先来顺一下解决这种问题的套路：
我们前文多次强调过，很显然只要涉及求最值，没有任何奇技淫巧，一定是穷举所有可能的结果，然后对比得出最值。
所以说，只要遇到求最值的算法问题，首先要思考的就是：如何穷举出所有可能的结果？
穷举主要有两种算法，就是回溯算法和动态规划，前者就是暴力穷举，而后者是根据状态转移方程推导「状态」。

定义 dp 数组的含义：

dp[i][j] = x 表示，戳破气球 i 和气球 j 之间（开区间，不包括 i 和 j）的所有气球，可以获得的最高分数为 x。

那么根据这个定义，题目要求的结果就是 dp[0][n+1] 的值，而 base case 就是 dp[i][j] = 0，其中 0 <= i <= n+1, j <= i+1，因为这种情况下，开区间 (i, j) 中间根本没有气球可以戳。

// base case 已经都被初始化为 0
int[][] dp = new int[n + 2][n + 2];

*/

/*
*
动态规划
首尾各添加一个元素1，方便计算，创建n+2的dp
dp[i][j] = x 表示，戳破气球 i 和气球 j 之间（开区间，不包括 i 和 j）的所有气球，可以获得的最高分数为 x。
dp[i][j]= max(val[i]×val[k]×val[j] + dp[i][k] + dp[k][j]), （i+1 < k < j-1） (i < j-1)
dp[i][j]= 0, （i+1 < k < j-1） (i >= j-1)
最终答案即为 dp[0][n+1]

(n-1, n) 是一个空区间（n-1+1 == n，没有气球可以戳），所以其值一定是 0，即默认值。
*/
func maxCoins(nums []int) int {

	n := len(nums)
	dp := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		dp[i] = make([]int, n+2)
	}

	// 冗余这份数据，方便后面计算
	val := make([]int, n+2)
	val[0], val[n+1] = 1, 1
	for i := 1; i <= n; i++ {
		val[i] = nums[i-1]
	}

	//dp[i][j] j-i<=1 的都等于0，没有气球可以戳
	for step := 2; step <= n+1; step++ {
		for i := 0; i <= n-step+1; i++ {
			j := i + step
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], val[i]*val[k]*val[j]+dp[i][k]+dp[k][j])
			}
		}
	}
	return dp[0][n+1]
}

/**
执行用时分布
39
ms
击败
60.32%
复杂度分析
消耗内存分布
7.32
MB
击败
57.94%
*/

func TestMaxCoins(t *testing.T) {

	fmt.Println(maxCoins([]int{3, 1, 5, 8}))
	fmt.Println(maxCoins([]int{1, 5}))
}
