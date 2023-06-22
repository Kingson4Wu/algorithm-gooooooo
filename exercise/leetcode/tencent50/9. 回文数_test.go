package tencent50

import (
	"fmt"
	"testing"
)

/**
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

例如，121 是回文，而 123 不是。


示例 1：

输入：x = 121
输出：true
示例 2：

输入：x = -121
输出：false
解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3：

输入：x = 10
输出：false
解释：从右向左读, 为 01 。因此它不是一个回文数。
*/
/**
简单的题，还是写不出题解的效果。。。。

反转一半数字
*/

/**
时间
12 ms
击败
68.65%
内存
4.1 MB
击败
91.68%
*/

func isPalindrome(x int) bool {

	//小于0 或者 个位数为0
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	reserveNum := 0
	for x > reserveNum {
		reserveNum = reserveNum*10 + x%10
		x /= 10
	}
	return reserveNum == x || reserveNum/10 == x
}

func TestIsPalindrome(t *testing.T) {

	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
}
