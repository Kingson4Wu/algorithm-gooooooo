package LCR

import (
	"fmt"
	"testing"
)

/*
*
自己想的
写得有点丑
还是做错了
*/
func digitOneInNumber(num int) int {
	count := 0
	pre := 0
	preNum := 0
	for pre <= num {
		if pre <= num {
			count += preNum
		} else {
			break
		}
		if pre+1 <= num {
			count += preNum + 1
		} else {
			break
		}
		for i := 2; i <= 9; i++ {
			v := pre + i
			if v <= num {
				count += preNum
			} else {
				break
			}
		}
		preNum++
		pre = pre*10 + 10
	}
	return count
}

func TestDigitOneInNumber(t *testing.T) {
	fmt.Println(digitOneInNumber(0))
	fmt.Println(digitOneInNumber(13))
	fmt.Println(digitOneInNumber(21))
}
