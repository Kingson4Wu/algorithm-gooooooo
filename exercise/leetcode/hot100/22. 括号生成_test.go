package hot100

import (
	"fmt"
	"testing"
)

/**
看了示例代码自己重写的

1、递归，子问题
2、'()' 添加在子串的各个位置
3、hash排重

时间
0 ms
击败
100%
内存
3.1 MB
击败
10.27%
*/

func generateParenthesis(n int) []string {

	if n == 1 {
		return []string{"()"}
	}
	resultMap := make(map[string]bool)
	var result []string
	for _, sub := range generateParenthesis(n - 1) {
		for i := 0; i < len(sub); i++ {
			s := sub[:i] + "()" + sub[i:]
			if ok, _ := resultMap[s]; !ok {
				resultMap[s] = true
				result = append(result, s)
			}
		}
	}
	return result
}

func TestGenerateParenthesis(t *testing.T) {
	fmt.Println(generateParenthesis(1))
	fmt.Println(generateParenthesis(2))
	fmt.Println(generateParenthesis(3))
}
