package hot100

import (
	"sort"
)

/**
个人想法，只会使用hashmap来解题。
(官方答案也确实这么狗)

时间
20 ms
击败
82.47%
内存
9.2 MB
击败
11.82%

*/

/**

给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。



示例 1:

输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
示例 2:

输入: strs = [""]
输出: [[""]]
示例 3:

输入: strs = ["a"]
输出: [["a"]]
*/

func groupAnagrams(strs []string) [][]string {

	group := make(map[string][]string)
	for _, str := range strs {
		key := sortString(str)
		if v, ok := group[key]; ok {
			v = append(v, str)
			group[key] = v
		} else {
			group[key] = []string{str}
		}
	}

	var result [][]string
	//result = append(result, []string{})
	for _, v := range group {
		result = append(result, v)
	}

	return result
}

func sortString(str string) string {
	b := []byte(str)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	return string(b)
}
