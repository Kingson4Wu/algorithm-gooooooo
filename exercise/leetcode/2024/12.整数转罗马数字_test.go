package _024

import (
	"fmt"
	"testing"
)

// 自己做的，真的中等，可能是自己的写法太丑了？

var romanCode = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
}

func intToRoman(num int) string {

	var result string
	index := 1
	for num > 0 {
		v := num % 10
		num /= 10

		if v <= 3 {
			code := romanCode[index]
			for i := 0; i < v; i++ {
				result = code + result
			}
		} else if v == 4 {
			code := romanCode[index]
			code5 := romanCode[index*5]
			result = code + code5 + result
		} else if v >= 5 && v <= 8 {
			code := romanCode[index]
			code5 := romanCode[index*5]
			for i := 0; i < v-5; i++ {
				result = code + result
			}
			result = code5 + result
		} else {
			code := romanCode[index]
			code5 := romanCode[index*10]
			result = code + code5 + result
		}
		index *= 10
	}

	return result
}

func TestName(*testing.T) {
	fmt.Println(intToRoman(16))
}
