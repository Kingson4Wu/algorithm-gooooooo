package _02510

/*
*
根据回忆和自己的思考 写对了
先计算出最大最小值
使用一个数组作为map，计算对应的下标，判断是否存在，如果重复或算出不在合理下标的范围，则不符合等差数列

执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
4.16
MB
击败
75.00%
*/
func canMakeArithmeticProgression(arr []int) bool {
	minVal, maxVal := arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > maxVal {
			maxVal = arr[i]
		} else if arr[i] < minVal {
			minVal = arr[i]
		}
	}
	if minVal == maxVal {
		return true
	}
	if (maxVal-minVal)%(len(arr)-1) != 0 {
		return false
	}
	step := (maxVal - minVal) / (len(arr) - 1)
	m := make([]bool, len(arr))
	for i := 0; i < len(arr); i++ {
		if (arr[i]-minVal)%step != 0 {
			return false
		}
		index := (arr[i] - minVal) / step
		if m[index] {
			return false
		}
		m[index] = true
	}
	return true
}
