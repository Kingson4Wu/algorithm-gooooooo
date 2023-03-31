package hot100

import (
	"fmt"
	"testing"
)

/**
自己做的

分析规律，未填入数据的数据列表一定是非递增的，所以列表的最近一个即最小值

时间
132 ms
击败
57.23%
内存
9.9 MB
击败
7.87%

本人的做法跟官方题解单调栈基本一样
方法二：单调栈
可以维护一个存储下标的单调栈，从栈底到栈顶的下标对应的温度列表中的温度依次递减。如果一个下标在单调栈里，则表示尚未找到下一次温度更高的下标。

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
