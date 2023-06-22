package topinterview145

import (
	"fmt"
	"testing"
)

/*
*
如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。

字母和数字都属于字母数字字符。

给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。

示例 1：

输入: s = "A man, a plan, a canal: Panama"
输出：true
解释："amanaplanacanalpanama" 是回文串。
示例 2：

输入：s = "race a car"
输出：false
解释："raceacar" 不是回文串。
示例 3：

输入：s = " "
输出：true
解释：在移除非字母数字字符之后，s 是一个空字符串 "" 。
由于空字符串正着反着读都一样，所以是回文串。

提示：

1 <= s.length <= 2 * 105
s 仅由可打印的 ASCII 字符组成
*/

/*
*
自己写的，感觉写得很啰嗦，有更简洁的写法？


*/
func isPalindrome(s string) bool {

	if len(s) <= 1 {
		return true
	}

	start := 0
	end := len(s) - 1

	for start < end {

		for start < end && !valid(s[start]) {
			start++
		}
		for start < end && !valid(s[end]) {
			end--
		}

		if !same(s[start], s[end]) {
			return false
		}
		start++
		end--
	}

	return true
}

func valid(r byte) bool {
	if r >= 'a' && r <= 'z' {
		return true
	}
	if r >= 'A' && r <= 'Z' {
		return true
	}
	return false
}
func same(a, b byte) bool {
	if toLower(a) == toLower(b) {
		return true
	}
	return false
}

func toLower(a byte) byte {
	if a >= 'A' && a <= 'Z' {
		return a + ('a' - 'A')
	}
	return a
}

func TestIsPalindrome(t *testing.T) {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome(" "))
	fmt.Println(isPalindrome("  "))
	fmt.Println(isPalindrome("-"))
	fmt.Println(isPalindrome("a-a"))
	fmt.Println(isPalindrome("a- b-=a"))
	fmt.Println(isPalindrome("a- g b-=a"))
}
