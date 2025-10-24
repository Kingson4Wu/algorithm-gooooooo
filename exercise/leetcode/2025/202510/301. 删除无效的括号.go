package _02510

/*
*
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
4.22
MB
击败
67.07%
*/
func removeInvalidParentheses(s string) []string {

	isValid := func(s string) bool {
		count := 0
		for i := 0; i < len(s); i++ {
			if s[i] == '(' {
				count++
			} else if s[i] == ')' {
				if count == 0 {
					return false
				}
				count--
			}
		}
		//漏了这句！！！！
		//return true
		return count == 0
	}

	findRemoveCount := func(s string) (int, int) {
		lRemove, rRemove := 0, 0
		for i := 0; i < len(s); i++ {
			if s[i] == '(' {
				lRemove++
			} else if s[i] == ')' {
				if lRemove > 0 {
					lRemove--
				} else {
					rRemove++
				}
			}
		}
		return lRemove, rRemove
	}

	var ans []string
	var dfs func(s string, start int, lRemove, rRemove int)
	dfs = func(s string, start int, lRemove, rRemove int) {
		//当字符串还没遍历完但括号已删光的情况会被跳过；
		//isValid 的检查应该在所有删完时进行；
		//不应该要求 start == len(s) 才算完，因为删除字符时 len(s) 在变。
		if lRemove == 0 && rRemove == 0 {
			if isValid(s) {
				ans = append(ans, s)
			}
			return
		}
		for i := start; i < len(s); i++ {
			if i > start && s[i-1] == s[i] {
				continue
			}
			if lRemove+rRemove > len(s)-start {
				break
			}
			if lRemove > 0 && s[i] == '(' {
				//dfs(s[:start]+s[start+1:], start+1, lRemove-1, rRemove)
				dfs(s[:i]+s[i+1:], i, lRemove-1, rRemove)
				//元素已经被删了，应该用i而不是i+1
			} else if rRemove > 0 && s[i] == ')' {
				//dfs(s[:start]+s[start+1:], start+1, lRemove, rRemove-1)
				dfs(s[:i]+s[i+1:], i, lRemove, rRemove-1)
			}
		}
	}
	lRemove, rRemove := findRemoveCount(s)
	dfs(s, 0, lRemove, rRemove)
	return ans
}
