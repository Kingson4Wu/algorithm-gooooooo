package string_match

// BF 算法实现
func BFSearch(haystack, needle string) int {
	m, n := len(haystack), len(needle)
	if n == 0 {
		return 0
	}
	for i := 0; i <= m-n; i++ {
		j := 0
		for ; j < n; j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if j == n {
			return i
		}
	}
	return -1
}

// 简单哈希函数：按字符值相加
func simpleHash(s string) int {
	hash := 0
	for _, ch := range s {
		hash += int(ch - 'a' + 1) // 假设只处理小写字母
	}
	return hash
}

// RK 算法实现
func RKSearch(haystack, needle string) int {
	m, n := len(haystack), len(needle)
	if n == 0 {
		return 0
	}
	if n > m {
		return -1
	}

	needleHash := simpleHash(needle)
	windowHash := simpleHash(haystack[:n])

	for i := 0; i <= m-n; i++ {
		if i > 0 {
			windowHash = windowHash - int(haystack[i-1]-'a'+1) + int(haystack[i+n-1]-'a'+1)
		}
		if windowHash == needleHash {
			// 发生哈希碰撞时逐字符比对
			match := true
			for j := 0; j < n; j++ {
				if haystack[i+j] != needle[j] {
					match = false
					break
				}
			}
			if match {
				return i
			}
		}
	}
	return -1
}
