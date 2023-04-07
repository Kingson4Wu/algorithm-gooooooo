package tencent50

/*
*
时间
28 ms
击败
82.25%
内存
6.3 MB
击败
53.18%
*/
func reverseString(s []byte) {

	if len(s) <= 1 {
		return
	}

	start, end := 0, len(s)-1

	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}

}
