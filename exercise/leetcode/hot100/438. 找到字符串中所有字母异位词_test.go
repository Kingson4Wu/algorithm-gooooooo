package hot100

import (
	"fmt"
	"testing"
)

/*
*

438. 找到字符串中所有字母异位词
中等
1.2K
相关企业
给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。

异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。

示例 1:

输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。

	示例 2:

输入: s = "abab", p = "ab"
输出: [0,1,2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。
*/

/**
竟然好像有思路

滑动窗口+hash表

感觉自己写得有点乱

优化，
1.记录map的数量，而不是每次遍历比对
2.

时间
28 ms
击败
21.15%
内存
4.9 MB
击败
12.98%

*/

func findAnagrams(s string, p string) []int {

	targetM := map[byte]int{}

	for i := 0; i < len(p); i++ {
		targetM[p[i]]++
	}

	match := func(a, b map[byte]int) bool {
		if len(a) != len(b) {
			return false
		}
		for k, v := range a {
			if vv, ok := b[k]; !ok || vv != v {
				return false
			}
		}
		return true
	}

	var result []int
	currentM := map[byte]int{}
	currentLen := 0

	for i := 0; i < len(s); i++ {

		if count, ok := targetM[s[i]]; !ok {
			for k := range currentM {
				delete(currentM, k)
			}
			currentLen = 0
			continue
		} else if currentM[s[i]] < count {
			currentM[s[i]]++
			currentLen++
		} else {
			/** 将超出 */
			start := i - currentLen
			if s[start] == s[i] {
				continue
			} else {

				length := currentLen
				// 低级错误，这里写错了！
				for j := start; j < start+length; j++ {
					if s[i] == s[j] {
						break
					}
					currentLen--
					currentM[s[j]]--
				}
			}
		}

		if match(currentM, targetM) {
			start := i - len(p) + 1
			result = append(result, start)
			currentM[s[start]]--
			currentLen--
		}

	}

	return result
}

func TestFindAnagrams(t *testing.T) {

	//fmt.Println(findAnagrams("cbaebabacd", "abc"))
	//fmt.Println(findAnagrams("abab", "ab"))
	fmt.Println(findAnagrams("cccccccbbbbbbbbbaaaaa", "abc"))
	//输出
	//[14]
	//预期结果
	//[]
}

/**

方法一：滑动窗口
思路

根据题目要求，我们需要在字符串 s 寻找字符串 p 的异位词。因为字符串 p 的异位词的长度一定与字符串 p 的长度相同，所以我们可以在字符串 s 中构造一个长度为与字符串 p 的长度相同的滑动窗口，并在滑动中维护窗口中每种字母的数量；当窗口中每种字母的数量与字符串 p 中每种字母的数量相同时，则说明当前窗口为字符串 p 的异位词。

算法

在算法的实现中，我们可以使用数组来存储字符串 p 和滑动窗口中每种字母的数量。

细节

当字符串 s 的长度小于字符串 p 的长度时，字符串 s 中一定不存在字符串 p 的异位词。但是因为字符串 s 中无法构造长度与字符串 p 的长度相同的窗口，所以这种情况需要单独处理。

func findAnagrams(s, p string) (ans []int) {
    sLen, pLen := len(s), len(p)
    if sLen < pLen {
        return
    }

    var sCount, pCount [26]int
    for i, ch := range p {
        sCount[s[i]-'a']++
        pCount[ch-'a']++
    }
    if sCount == pCount {
        ans = append(ans, 0)
    }

    for i, ch := range s[:sLen-pLen] {
        sCount[ch-'a']--
        sCount[s[i+pLen]-'a']++
        if sCount == pCount {
            ans = append(ans, i+1)
        }
    }
    return
}

方法二：优化的滑动窗口
思路和算法

在方法一的基础上，我们不再分别统计滑动窗口和字符串 p 中每种字母的数量，而是统计滑动窗口和字符串 p 中每种字母数量的差；并引入变量 differ 来记录当前窗口与字符串 p 中数量不同的字母的个数，并在滑动窗口的过程中维护它。

在判断滑动窗口中每种字母的数量与字符串 p 中每种字母的数量是否相同时，只需要判断 differ 是否为零即可。

class Solution:
    def findAnagrams(self, s: str, p: str) -> List[int]:
        s_len, p_len = len(s), len(p)

        if s_len < p_len:
            return []

        ans = []
        count = [0] * 26
        for i in range(p_len):
            count[ord(s[i]) - 97] += 1
            count[ord(p[i]) - 97] -= 1

        differ = [c != 0 for c in count].count(True)

        if differ == 0:
            ans.append(0)

        for i in range(s_len - p_len):
            if count[ord(s[i]) - 97] == 1:  # 窗口中字母 s[i] 的数量与字符串 p 中的数量从不同变得相同
                differ -= 1
            elif count[ord(s[i]) - 97] == 0:  # 窗口中字母 s[i] 的数量与字符串 p 中的数量从相同变得不同
                differ += 1
            count[ord(s[i]) - 97] -= 1

            if count[ord(s[i + p_len]) - 97] == -1:  # 窗口中字母 s[i+p_len] 的数量与字符串 p 中的数量从不同变得相同
                differ -= 1
            elif count[ord(s[i + p_len]) - 97] == 0:  # 窗口中字母 s[i+p_len] 的数量与字符串 p 中的数量从相同变得不同
                differ += 1
            count[ord(s[i + p_len]) - 97] += 1

            if differ == 0:
                ans.append(i + 1)

        return ans

作者：力扣官方题解
链接：https://leetcode.cn/problems/find-all-anagrams-in-a-string/solutions/1123971/zhao-dao-zi-fu-chuan-zhong-suo-you-zi-mu-xzin/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
