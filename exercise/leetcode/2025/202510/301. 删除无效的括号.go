package _02510

func removeInvalidParentheses(s string) []string {

	isValid := func(s string) bool {
		count := 0
		for i := 0; i < len(s); i++ {
			if s[i] == '(' {
				count++
			} else {
				if count == 0 {
					return false
				}
				count--
			}
		}
		return true
	}

	findRemoveCount := func(s string) (int, int) {
		lRemove, rRemove := 0, 0
		for i := 0; i < len(s); i++ {
			if s[i] == '(' {
				lRemove++
			} else {
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
		if start == len(s) && lRemove == 0 && rRemove == 0 && isValid(s) {
			ans = append(ans, s)
			return
		}
		for i := start; i < len(s); i++ {
			if i > start && s[i-1] == s[i] {
				continue
			}
			if lRemove+rRemove > len(s)-start {
				break
			}
			if s[i] == '(' {
				dfs(s[:start]+s[start+1:], start+1, lRemove-1, rRemove)
			} else if s[i] == ')' {
				dfs(s[:start]+s[start+1:], start+1, lRemove, rRemove-1)
			}
		}
	}
	lRemove, rRemove := findRemoveCount(s)
	dfs(s, 0, lRemove, rRemove)
	return ans
}
