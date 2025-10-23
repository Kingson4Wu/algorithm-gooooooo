package _02510

/*
*
执行用时分布
3
ms
击败
26.32%
复杂度分析
消耗内存分布
4.79
MB
击败
10.53%
*/
func generateParenthesis(n int) []string {

	if n == 1 {
		return []string{"()"}
	}
	var ans []string
	duplicated := make(map[string]bool)
	sub := generateParenthesis(n - 1)
	for _, s := range sub {
		for i := 0; i < len(s); i++ {
			str := s[:i] + "()" + s[i:]
			if _, ok := duplicated[str]; !ok {
				ans = append(ans, str)
				duplicated[str] = true
			}
		}
	}
	return ans
}
