package topinterview145

import (
	"sort"
	"strconv"
)

/**
给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。

注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。



示例 1：

输入：nums = [10,2]
输出："210"
示例 2：

输入：nums = [3,30,34,5,9]
输出："9534330"


提示：

1 <= nums.length <= 100
0 <= nums[i] <= 109
*/
/**
直接排序
个位数9 前面

*/
/*func largestNumber(nums []int) string {

	getNumArr := func(num int) []int {
		var result []int
		for num > 0 {
			v := num % 10
			result = append(result, v)
			num /= 10
		}
		return result
	}

	sort.Slice(nums, func(i, j int) bool {

		x := getNumArr(i)
		y := getNumArr(j)
		if len(x) == 0 {
			return false
		} else if len(y) == 0 {
			return true
		} else {
			for len(x) > 0 && len(y) > 0 {
				if x[len(x)-1] > y[len(y)-1] {
					return true
				} else if x[len(x)-1] < y[len(y)-1] {
					return false
				}
				x = x[:len(x)-1]
				y = y[:len(y)-1]
			}

		}

		return true
	})

	return ""
}*/

func largestNumber(nums []int) string {

	//取巧，答案是总结另外的规律
	sort.Slice(nums, func(i, j int) bool {
		a, b := strconv.Itoa(nums[i]), strconv.Itoa(nums[j])
		return a+b > b+a
	})
	if nums[0] == 0 {
		return "0"
	}
	var ans []byte
	for _, x := range nums {
		ans = append(ans, strconv.Itoa(x)...)
	}
	return string(ans)
}
