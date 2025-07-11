package topinterview150

import (
	"fmt"
	"testing"
)

/*
*
给定一个字符串 s 和一个字符串数组 words。 words 中所有字符串 长度相同。

	s 中的 串联子串 是指一个包含  words 中所有字符串以任意顺序排列连接起来的子串。

例如，如果 words = ["ab","cd","ef"]， 那么 "abcdef"， "abefcd"，"cdabef"， "cdefab"，"efabcd"， 和 "efcdab" 都是串联子串。 "acdbef" 不是串联子串，因为他不是任何 words 排列的连接。
返回所有串联子串在 s 中的开始索引。你可以以 任意顺序 返回答案。

示例 1：

输入：s = "barfoothefoobarman", words = ["foo","bar"]
输出：[0,9]
解释：因为 words.length == 2 同时 words[i].length == 3，连接的子字符串的长度必须为 6。
子串 "barfoo" 开始位置是 0。它是 words 中以 ["bar","foo"] 顺序排列的连接。
子串 "foobar" 开始位置是 9。它是 words 中以 ["foo","bar"] 顺序排列的连接。
输出顺序无关紧要。返回 [9,0] 也是可以的。
示例 2：

输入：s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
输出：[]
解释：因为 words.length == 4 并且 words[i].length == 4，所以串联子串的长度必须为 16。
s 中没有子串长度为 16 并且等于 words 的任何顺序排列的连接。
所以我们返回一个空数组。
示例 3：

输入：s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
输出：[6,9,12]
解释：因为 words.length == 3 并且 words[i].length == 3，所以串联子串的长度必须为 9。
子串 "foobarthe" 开始位置是 6。它是 words 中以 ["foo","bar","the"] 顺序排列的连接。
子串 "barthefoo" 开始位置是 9。它是 words 中以 ["bar","the","foo"] 顺序排列的连接。
子串 "thefoobar" 开始位置是 12。它是 words 中以 ["the","foo","bar"] 顺序排列的连接。

提示：

1 <= s.length <= 104
1 <= words.length <= 5000
1 <= words[i].length <= 30
words[i] 和 s 由小写英文字母组成

自己想的
hashmap + 滑动窗口
hashmap 记录单词库以及个数， 记录当前窗口的单词是否存在
单词出现超过1次时，将左窗口右移到第一次出现的位置的下一个单词，同时窗口对应的hash做调整
单词不存在时，清空窗口的hashmap
注意单词库的单词可以出现多次！

按单词长度分组， 比如 单词长度为2，总共5个单词，则总长度为10， 而s的长度为23
则分组成：
0,2,4,6,8,10,12,14,16,18,20
1,3,5,7,9,11,13,15,17,18,19,21
总共两组

。
*/
func findSubstring(s string, words []string) []int {

	var ans []int
	wordLength := len(words[0])
	wordCount := len(words)
	strLength := wordLength * wordCount
	if len(s) < strLength {
		return ans
	}
	//单词库
	wordsLib := make(map[string]int, wordCount)
	// 子串的单词下标, 可以有多个，判断是否存在
	strWordsIndex := make(map[string][]int, wordCount)
	strWordsCount := 0
	for i := 0; i < wordCount; i++ {
		wordsLib[words[i]]++
	}

	for i := 0; i < wordLength; i++ {
		if len(s)-i < strLength {
			break
		}
		//清空
		strWordsCount = 0
		for k := range strWordsIndex {
			delete(strWordsIndex, k)
		}
		for j := i; j <= len(s)-wordLength; j += wordLength {
			w := s[j : j+wordLength]
			count := wordsLib[w]
			if count > 0 {
				if indexes, ok := strWordsIndex[w]; !ok || len(indexes) < count {
					if arr, ok := strWordsIndex[w]; !ok {
						strWordsIndex[w] = []int{j}
					} else {
						strWordsIndex[w] = append(arr, j)
					}
					strWordsCount++
					if strWordsCount == wordCount {
						firstIndex := j + wordLength - strLength
						ans = append(ans, firstIndex)
						firstWord := s[firstIndex : firstIndex+wordLength]
						strWordsIndex[firstWord] = strWordsIndex[firstWord][1:]
						strWordsCount--
					}
				} else {
					//左窗口右移到当前
					index := indexes[0]
					for k, wi := range strWordsIndex {
						newWi := wi[:0] // 复用底层数组，避免频繁分配
						//这个 range 会在循环开始时 复制 wi 整个切片结构（包括指向底层数组的指针 + len + cap）。
						//也就是说，就算你在循环里改 wi[:0]，不会影响到 range 所用的 wi 的副本，因为底层数组还在、长度已复制。
						for _, ix := range wi {
							if ix > index {
								newWi = append(newWi, ix)
							} else {
								strWordsCount--
							}
						}
						if len(newWi) == 0 {
							delete(strWordsIndex, k)
						} else {
							strWordsIndex[k] = newWi
						}
					}

					if arr, ok := strWordsIndex[w]; !ok {
						strWordsIndex[w] = []int{j}
					} else {
						strWordsIndex[w] = append(arr, j)
					}
					strWordsCount++
				}
			} else {
				//清空
				strWordsCount = 0
				for k := range strWordsIndex {
					delete(strWordsIndex, k)
				}
			}
		}
	}
	return ans
}

/**
解答错误
6 / 182 个通过的测试用例

官方题解
输入
s =
"wordgoodgoodgoodbestword"
words =
["word","good","best","good"]

添加到测试用例
输出
[]
预期结果
[8]

执行用时分布
17
ms
击败
46.07%
复杂度分析
消耗内存分布
7.41
MB
击败
92.50%
复杂度分析

*/

func TestFindSubstring(t *testing.T) {
	fmt.Println(findSubstring("bcabbcaabbccacacbabccacaababcbb", []string{"c", "b", "a", "c", "a", "a", "a", "b", "c"}))
	fmt.Println(findSubstring("ababababab", []string{"ababa", "babab"}))
	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}))
	fmt.Println(findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}))
	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}))
}
