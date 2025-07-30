package backtrack

/*
*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
*/
func letterCombinations(digits string) []string {

	var phoneMap = map[rune]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	var ans []string
	var path []rune
	// 不是组合，每个位置上不存在选跟不选的问题，只是要选哪个的问题
	var dfs func(cur int)
	dfs = func(cur int) {
		if len(path) == len(digits) {
			ans = append(ans, string(path))
			return
		}
		for _, letter := range phoneMap[rune(digits[cur])] {
			path = append(path, letter)
			dfs(cur + 1)
			path = path[:len(path)-1]
		}
	}

	if digits != "" {
		dfs(0)
	}
	return ans
}

/**
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
3.85
MB
击败
68.23%
复杂度分析

*/
