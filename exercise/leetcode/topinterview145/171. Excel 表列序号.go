package topinterview145

/*
*
比较简单
*/

/*
*
给你一个字符串 columnTitle ，表示 Excel 表格中的列名称。返回 该列名称对应的列序号 。

例如：

A -> 1
B -> 2
C -> 3
...
Z -> 26
AA -> 27
AB -> 28
...

示例 1:

输入: columnTitle = "A"
输出: 1
示例 2:

输入: columnTitle = "AB"
输出: 28
示例 3:

输入: columnTitle = "ZY"
输出: 701
*/
func titleToNumber(columnTitle string) int {

	res := 0

	for i := 0; i < len(columnTitle); i++ {

		num := (int)(columnTitle[i] - 'A' + 1)
		res = 26*res + num

	}

	return res

}
