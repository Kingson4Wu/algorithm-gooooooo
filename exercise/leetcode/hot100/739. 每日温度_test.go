package hot100

import (
	"fmt"
	"testing"
)

/**
自己做的

分析规律，未填入数据的数据列表一定是非递增的，所以列表的最近一个即最小值

*/
/*
*
给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。

示例 1:

输入: temperatures = [73,74,75,71,69,72,76,73]
输出: [1,1,4,2,1,1,0,0]
示例 2:

输入: temperatures = [30,40,50,60]
输出: [1,1,1,0]
示例 3:

输入: temperatures = [30,60,90]
输出: [1,1,0]
*/
func dailyTemperatures(temperatures []int) []int {

	if len(temperatures) == 0 {
		return []int{}
	}
	if len(temperatures) == 1 {
		return []int{0}
	}

	var unhandled []int
	result := make([]int, len(temperatures))

	unhandled = append(unhandled, 0)

	for i := 1; i < len(temperatures); i++ {
		if temperatures[i] <= temperatures[unhandled[len(unhandled)-1]] {
			unhandled = append(unhandled, i)
			continue
		}
		for j := len(unhandled) - 1; j >= 0; j-- {
			if temperatures[unhandled[j]] < temperatures[i] {
				result[unhandled[j]] = i - unhandled[j]
				unhandled = unhandled[:len(unhandled)-1]
			} else {
				break
			}
		}
		unhandled = append(unhandled, i)
	}
	return result
}

func TestDailyTemperatures(t *testing.T) {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	fmt.Println(dailyTemperatures([]int{30, 40, 50, 60}))
	fmt.Println(dailyTemperatures([]int{30, 60, 90}))
}
