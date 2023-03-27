package hot100

import (
	"fmt"
	"math"
	"testing"
)

/*
自己做的

+1 和溢出的情况没考虑到

官方题解不知道是啥
*/
func myAtoi(s string) int {

	result := 0
	start := false
	negative := false
	for _, c := range s {
		if c >= '0' && c <= '9' {
			if !start {
				start = true
			}
			origin := result
			result = result*10 + int(c-'0')

			if result < 0 || result < origin {

				if !negative {

					return math.MaxInt32
				}

				if negative {
					return math.MinInt32
				}

			}

		} else if c == '-' {
			if !start {
				start = true
				negative = true
			} else {
				break
			}
		} else if c == '+' {
			if !start {
				start = true
				negative = false
			} else {
				break
			}
		} else if c == ' ' && !start {
			continue
		} else {
			break
		}
	}

	if negative {
		result = -result
	}

	if result > math.MaxInt32 {
		return math.MaxInt32
	} else if result < math.MinInt32 {
		return math.MinInt32
	}

	return result

}

/*
*
时间
0 ms
击败
100%
内存
2 MB
击败
78.60%
*/
func TestMyAtoi(t *testing.T) {
	fmt.Println(myAtoi("18446744073709551617"))
}
