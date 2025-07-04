package leetcode

/*
*
给你两个整数，被除数 dividend 和除数 divisor。将两数相除，要求 不使用 乘法、除法和取余运算。

整数除法应该向零截断，也就是截去（truncate）其小数部分。例如，8.345 将被截断为 8 ，-2.7335 将被截断至 -2 。

返回被除数 dividend 除以除数 divisor 得到的 商 。

注意：假设我们的环境只能存储 32 位 有符号整数，其数值范围是 [−231,  231 − 1] 。本题中，如果商 严格大于 231 − 1 ，则返回 231 − 1 ；如果商 严格小于 -231 ，则返回 -231 。

示例 1:

输入: dividend = 10, divisor = 3
输出: 3
解释: 10/3 = 3.33333.. ，向零截断后得到 3 。
示例 2:

输入: dividend = 7, divisor = -3
输出: -2
解释: 7/-3 = -2.33333.. ，向零截断后得到 -2 。

提示：

-231 <= dividend, divisor <= 231 - 1
divisor != 0
*/
func divide(dividend int, divisor int) int {

	const IntMax = 1<<31 - 1 // 2147483647
	const IntMin = -1 << 31  // -2147483648

	// 特殊溢出处理
	if dividend == IntMin && divisor == -1 {
		return IntMax
	}

	reverse := false
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		reverse = true
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	count := 0
	for dividend >= divisor {
		count++
		dividend -= divisor
	}
	if reverse {
		return -count
	}
	return count
}

/**
解答错误
11 / 994 个通过的测试用例

官方题解
输入
dividend =
-2147483648
divisor =
-1

添加到测试用例
输出
2147483648
预期结果
2147483647

如果商 严格大于 231 − 1 ，则返回 231 − 1 ；如果商 严格小于 -231 ，则返回 -231
*/

/**
超出时间限制
13 / 994 个通过的测试用例
最后执行的输入
添加到测试用例
dividend =
2147483647
divisor =
2
*/
