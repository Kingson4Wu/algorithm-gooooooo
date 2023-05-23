package hot100

import (
	"fmt"
	"testing"
)

/*
*

给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符

示例 1：

输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
示例 2：

输入：word1 = "intention", word2 = "execution"
输出：5
解释：
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')
*/

/**
看了题解：

对A插入 等同于 对B删除
对A删除 等同于 对B插入
对A替换 等同于 对B替换

所以只需要关注对A的操作即可

设dp[i,j] ,表示， 长度为i的A子串 转换到 长度为j的B子串 的距离
显然，可以算出dp[0,y], dp[x,0]的距离分别是y和x

如果s[i] == s[j], 则 dp[i,j] = dp[i-1, j-1]
否则
1. 插入 dp[i,j]= dp[i-1, j] + 1
2. 删除 dp[i,j]= dp[i, j-1] + 1
3. 替换 dp[i,j]= dp[i-1, j-1] + 1

动态规划
本质上就是避免重复计算的穷举

时间
0 ms
击败
100%
内存
5.3 MB
击败
28.55%
*/

func minDistance(word1 string, word2 string) int {

	x := len(word1)
	y := len(word2)
	dp := make([][]int, x+1)
	for i := 0; i <= x; i++ {
		dp[i] = make([]int, y+1)
	}
	/** 初始化 */
	for i := 0; i <= x; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= y; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= x; i++ {
		for j := 1; j <= y; j++ {
			// 下标为0代表第一个，长度为0，所以要减1
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(min(dp[i-1][j]+1, dp[i][j-1]+1), dp[i-1][j-1]+1)
			}
		}
	}

	return dp[x][y]
}

func TestMinDistance(t *testing.T) {
	fmt.Println(minDistance("horse", "ros"))
	fmt.Println(minDistance("intention", "execution"))
}
