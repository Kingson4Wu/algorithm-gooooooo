package hot100

/*
*
自己做的，算出每一行的间隔规律
解题方式属于官方第三种，不过官方规律更精简

	func convert(s string, numRows int) string {
	    n, r := len(s), numRows
	    if r == 1 || r >= n {
	        return s
	    }
	    t := r*2 - 2
	    ans := make([]byte, 0, n)
	    for i := 0; i < r; i++ { // 枚举矩阵的行
	        for j := 0; j+i < n; j += t { // 枚举每个周期的起始下标
	            ans = append(ans, s[j+i]) // 当前周期的第一个字符
	            if 0 < i && i < r-1 && j+t-i < n {
	                ans = append(ans, s[j+t-i]) // 当前周期的第二个字符
	            }
	        }
	    }
	    return string(ans)
	}

作者：力扣官方题解
链接：https://leetcode.cn/problems/zigzag-conversion/solutions/1298127/z-zi-xing-bian-huan-by-leetcode-solution-4n3u/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func convert(s string, numRows int) string {

	if numRows == 1 {
		return s
	}

	result := make([]byte, len(s))
	index := 0
	for i := 1; i <= numRows; i++ {
		if i == 1 || i == numRows {
			gap := numRows*2 - 2
			j := i - 1
			for j < len(s) {
				result[index] = s[j]
				index++
				j += gap
			}
			continue
		}
		gap1 := (numRows - i) * 2
		gap2 := (i - 1) * 2
		j := i - 1
		gap1Add := true
		for j < len(s) {
			result[index] = s[j]
			index++
			if gap1Add {
				j += gap1
				gap1Add = false
			} else {
				j += gap2
				gap1Add = true
			}
		}
	}
	return string(result)

}
