package topinterview150

import (
	"fmt"
	"testing"
)

/*
*
给定三个字符串 s1、s2、s3，请你帮忙验证 s3 是否是由 s1 和 s2 交错 组成的。

两个字符串 s 和 t 交错 的定义与过程如下，其中每个字符串都会被分割成若干 非空 子字符串：

s = s1 + s2 + ... + sn
t = t1 + t2 + ... + tm
|n - m| <= 1
交错 是 s1 + t1 + s2 + t2 + s3 + t3 + ... 或者 t1 + s1 + t2 + s2 + t3 + s3 + ...
注意：a + b 意味着字符串 a 和 b 连接。

示例 1：

输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
输出：true
示例 2：

输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
输出：false
示例 3：

输入：s1 = "", s2 = "", s3 = ""
输出：true

提示：

0 <= s1.length, s2.length <= 100
0 <= s3.length <= 200
s1、s2、和 s3 都由小写英文字母组成

进阶：您能否仅使用 O(s2.length) 额外的内存空间来解决它?
*/
/**
怎么感觉好像挺简单，两个指针就可以搞定？？
理解错了，不能先到先得
"aabcc", "dbbca", "aadbbcbcac"
最后会剩下ca 和 ac 无法匹配

----

动态规划
 dp(i,j) 表示 s1的前 i 个元素和 s2 的前 j 个元素是否能交错组成 s3的前 i+j 个元素
 dp(i,j) = (dp[i-1,j] && s1[i-1] == s3[i+j-1]) ||  (dp[i,j-1] && s2[j-1] == s3[i+j-1])
 s1的第i个下标为i-1
 dp[0][0] =true

可滚动数组优化空间复杂度， 空间降为 O(min(m, n))
dp 的第 i 行只和第 i−1 行相关

*/
/*func isInterleave(s1 string, s2 string, s3 string) bool {

	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 1; i <= m; i++ {
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
		if !dp[i][0] {
			break
		}
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = dp[0][j-1] && s2[j-1] == s3[j-1]
		if !dp[0][j] {
			break
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
		}
	}
	return dp[m][n]
}*/
func isInterleave(s1 string, s2 string, s3 string) bool {

	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 1; i <= m; i++ {
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
		if !dp[i][0] {
			break
		}
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = dp[0][j-1] && s2[j-1] == s3[j-1]
		if !dp[0][j] {
			break
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
		}
	}
	return dp[m][n]
}

/**
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
4.01
MB
击败
51.31%

*/

func TestIsInterleave(t *testing.T) {
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbcac"))
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbbaccc"))
	fmt.Println(isInterleave("", "", ""))
}

/**
官方答案：
func isInterleave(s1 string, s2 string, s3 string) bool {
    n, m, t := len(s1), len(s2), len(s3)
    if (n + m) != t {
        return false
    }
    f := make([][]bool, n + 1)
    for i := 0; i <= n; i++ {
        f[i] = make([]bool, m + 1)
    }
    f[0][0] = true
    for i := 0; i <= n; i++ {
        for j := 0; j <= m; j++ {
            p := i + j - 1
            if i > 0 {
                f[i][j] = f[i][j] || (f[i-1][j] && s1[i-1] == s3[p])
            }
            if j > 0 {
                f[i][j] = f[i][j] || (f[i][j-1] && s2[j-1] == s3[p])
            }
        }
    }
    return f[n][m]
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/interleaving-string/solutions/335373/jiao-cuo-zi-fu-chuan-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

一维

func isInterleave(s1 string, s2 string, s3 string) bool {
    n, m, t := len(s1), len(s2), len(s3)
    if (n + m) != t {
        return false
    }
    f := make([]bool, m + 1)
    f[0] = true
    for i := 0; i <= n; i++ {
        for j := 0; j <= m; j++ {
            p := i + j - 1
            if i > 0 {
                f[j] = f[j] && s1[i-1] == s3[p]
            }
            if j > 0 {
                f[j] = f[j] || f[j-1] && s2[j-1] == s3[p]
            }
        }
    }
    return f[m]
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/interleaving-string/solutions/335373/jiao-cuo-zi-fu-chuan-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
