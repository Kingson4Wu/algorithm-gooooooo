package LCR

import (
	"fmt"
	"testing"
)

// 实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn）。
/**
凭记忆写完

执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
3.84
MB
击败
32.14%
*/
func myPow(x float64, n int) float64 {
	m := make(map[int]float64)
	m[0] = 1
	m[1] = x
	k := n
	if k < 0 {
		k = -k
	}
	t := 1
	for i := 2; i <= k; i *= 2 {
		pre := m[i>>1]
		m[i] = pre * pre
		t = i
	}
	var result float64 = 1
	for k > 0 {
		if v, ok := m[k]; ok {
			result *= v
			break
		}
		if k > t {
			result *= m[t]
			k -= t
		} else {
			t /= 2
		}
	}
	if n > 0 {
		return result
	}
	return 1 / result
}

func TestPow(t *testing.T) {
	//fmt.Println(myPow(2.1, 1))
	fmt.Println(myPow(2, 10))
	//fmt.Println(myPow(2.1, 3))
	//fmt.Println(myPow(2, -2))
	//fmt.Println(myPow(0.97267, 13))

}
