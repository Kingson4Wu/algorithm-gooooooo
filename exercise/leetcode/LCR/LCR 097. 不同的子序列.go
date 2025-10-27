package LCR

/*
*
感觉跟删除括号那道题很像
超出时间限制

答案要用动态规划
创建二维数组 dp，其中 dp[i][j] 表示在 s[i:] 的子序列中 t[j:] 出现的个数
*/
func numDistinct(s string, t string) int {
	if s == t {
		return 1
	}
	removeLength := len(s) - len(t)
	if removeLength <= 0 {
		return 0
	}
	count := 0
	var dfs func(remain string, start, removeLength int)
	dfs = func(remain string, start, removeLength int) {
		if removeLength == 0 {
			if remain == t {
				count++
			}
			return
		}
		if len(remain)-start < removeLength {
			return
		}

		for i := start; i < len(remain); i++ {
			//if strings.HasPrefix(s, remain[:i]) {
			dfs(remain[:i]+remain[i+1:], i, removeLength-1)
			//}
		}
	}
	dfs(s, 0, removeLength)
	return count
}

/**
超出时间限制
41 / 63 个通过的测试用例
最后执行的输入
添加到测试用例
s =
"dbaaadcddccdddcadacbadbadbabbbcad"
t =
"dadcccbaab"
*/
