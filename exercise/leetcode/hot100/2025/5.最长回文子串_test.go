package _025

import (
	"fmt"
	"testing"
)

func longestPalindrome(s string) string {

	//maxI, maxJ := 0, 0
	maxI := 0
	maxLength := 1

	p := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		p[i] = make([]bool, len(s))
		//长度为1的初始化
		p[i][i] = true
	}

	for length := 2; length <= len(s); length++ {
		for i := 0; i+length-1 < len(s); i++ {
			j := i + length - 1

			if s[i] != s[j] {
				p[i][j] = false
			} else {
				if length == 2 {
					p[i][j] = true
				} else {
					p[i][j] = p[i+1][j-1]
				}
				if maxLength < length && p[i][j] {
					maxLength = length
					maxI = i
					//maxJ = j
				}
			}
		}
	}

	//return s[maxI : maxJ+1]
	return s[maxI : maxI+maxLength]
}

func TestLongestPalindrome(t *testing.T) {

	fmt.Println(longestPalindrome("abcabcbb"))
	fmt.Println(longestPalindrome("bbbbb"))
	fmt.Println(longestPalindrome("pwwkew"))
	fmt.Println(longestPalindrome(" "))

}
