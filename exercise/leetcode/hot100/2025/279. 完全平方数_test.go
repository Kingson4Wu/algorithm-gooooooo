package _025

import (
	"fmt"
	"math"
	"testing"
)

/*
*
个人想法，开平方，找到最大的平方数，逐步求余
12：

	// 9 + 1 + 1 +1 = 4; 4 + 4 + 4 = 3
	//我的方法不对
	// 以下代码错误的
*/
func numSquares(n int) int {

	k := int(math.Sqrt(float64(n)))
	v := k * k
	count := 0
	for n > 0 && k > 0 {
		if n >= v {
			n -= v
			count++
		} else {
			k--
			v = k * k
		}
	}
	return count
}

func TestNumSquares(t *testing.T) {
	fmt.Println(numSquares(12))
	// 9 + 1 + 1 +1 = 4; 4 + 4 + 4 = 3
	//我的方法不对
	fmt.Println(numSquares(13))
	fmt.Println(numSquares(0))
	fmt.Println(numSquares(1))
	fmt.Println(numSquares(2))
	fmt.Println(numSquares(3))
	fmt.Println(numSquares(4))
	fmt.Println(numSquares(5))
}

/**
3
2
0
1
2
3
1
2
*/
