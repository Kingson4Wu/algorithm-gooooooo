package LCR

/*
*
自己想的，
就是一个非递增单调栈
遍历时一直维护和计数（累加栈的长度即可）
这个做法不对

分段的单调栈, 类似分段排序
[7,5,6,4]

[7,6]
[5]

不对，好像是要用分段树之类的。。。
离散化树状数组
*/
func reversePairs(record []int) int {
	var stack []int
	count := 0
	for i := 0; i < len(record); i++ {
		for len(stack) > 0 && stack[len(stack)-1] <= record[i] {
			stack = stack[:len(stack)-1]
		}
		count += len(stack)
		stack = append(stack, record[i])
	}
	return count
}

/**
解答错误
执行用时: 0 ms
Case 1
输入
record =
[7,5,6,4]
输出
4
预期结果
5
*/
