package topinterview150

import (
	"fmt"
	"testing"
)

/**
自己做出来的，很多小细节低级错误，需要通过debug修正
遍历记录将要输出行的下标和所有单词长度（或剩余数量）

执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
4.03
MB
击败
13.84%
*/

/*
*
给定一个单词数组 words 和一个长度 maxWidth ，重新排版单词，使其成为每行恰好有 maxWidth 个字符，且左右两端对齐的文本。

你应该使用 “贪心算法” 来放置给定的单词；也就是说，尽可能多地往每行中放置单词。必要时可用空格 ' ' 填充，使得每行恰好有 maxWidth 个字符。

要求尽可能均匀分配单词间的空格数量。如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数。

文本的最后一行应为左对齐，且单词之间不插入额外的空格。

注意:

单词是指由非空格字符组成的字符序列。
每个单词的长度大于 0，小于等于 maxWidth。
输入单词数组 words 至少包含一个单词。

示例 1:

输入: words = ["This", "is", "an", "example", "of", "text", "justification."], maxWidth = 16
输出:
[

	"This    is    an",
	"example  of text",
	"justification.  "

]
示例 2:

输入:words = ["What","must","be","acknowledgment","shall","be"], maxWidth = 16
输出:
[

	"What   must   be",
	"acknowledgment  ",
	"shall be        "

]
解释: 注意最后一行的格式应为 "shall be    " 而不是 "shall     be",

	因为最后一行应为左对齐，而不是左右两端对齐。
	第二行同样为左对齐，这是因为这行只包含一个单词。

示例 3:

输入:words = ["Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"]，maxWidth = 20
输出:
[

	"Science  is  what we",
	"understand      well",
	"enough to explain to",
	"a  computer.  Art is",
	"everything  else  we",
	"do                  "

]
*/
func fullJustify(words []string, maxWidth int) []string {

	var results []string
	var indexes []int
	remainLength := maxWidth

	for i := 0; i < len(words); i++ {

		if len(words[i]) <= remainLength-len(indexes) {
			indexes = append(indexes, i)
			remainLength -= len(words[i])
		} else {
			// 只有一个单词的场景
			if len(indexes) == 1 {
				var suffix string
				for j := 0; j < remainLength; j++ {
					suffix += " "
				}
				results = append(results, words[indexes[0]]+suffix)
			} else {
				avgSpace := remainLength / (len(indexes) - 1)
				leftSpace := remainLength - (avgSpace * (len(indexes) - 1))
				var result string
				for i, index := range indexes {
					result += words[index]
					if i < len(indexes)-1 {
						for j := 0; j < avgSpace; j++ {
							result += " "
						}
						if leftSpace > 0 {
							result += " "
							leftSpace--
						}
					}
				}
				results = append(results, result)
			}
			remainLength = maxWidth
			indexes = []int{}
			i--
			continue
		}

		// 最后一行的场景
		if i == len(words)-1 {
			var result string
			for j, index := range indexes {
				result += words[index]
				if j < len(indexes)-1 {
					result += " "
				}
			}
			remain := maxWidth - len(result)
			for j := 0; j < remain; j++ {
				result += " "
			}
			results = append(results, result)
		}
	}

	return results
}

func TestFullJustify(t *testing.T) {

	/*for _, line := range fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16) {
		fmt.Println(line)
	}*/
	fmt.Println("====")
	for _, line := range fullJustify([]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16) {
		fmt.Println(line)
	}
	fmt.Println("====")
	for _, line := range fullJustify([]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20) {
		fmt.Println(line)
	}
}
