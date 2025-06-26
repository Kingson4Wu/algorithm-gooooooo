package topinterview150

import (
	"fmt"
	"strconv"
	"testing"
)

/**
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
3.71
MB
击败
82.81%

*/
/*
*
给定一个  无重复元素 的 有序 整数数组 nums 。

区间 [a,b] 是从 a 到 b（包含）的所有整数的集合。

返回 恰好覆盖数组中所有数字 的 最小有序 区间范围列表 。也就是说，nums 的每个元素都恰好被某个区间范围所覆盖，并且不存在属于某个区间但不属于 nums 的数字 x 。

列表中的每个区间范围 [a,b] 应该按如下格式输出：

"a->b" ，如果 a != b
"a" ，如果 a == b

示例 1：

输入：nums = [0,1,2,4,5,7]
输出：["0->2","4->5","7"]
解释：区间范围是：
[0,2] --> "0->2"
[4,5] --> "4->5"
[7,7] --> "7"
示例 2：

输入：nums = [0,2,3,4,6,8,9]
输出：["0","2->4","6","8->9"]
解释：区间范围是：
[0,0] --> "0"
[2,4] --> "2->4"
[6,6] --> "6"
[8,9] --> "8->9"

提示：

0 <= nums.length <= 20
-231 <= nums[i] <= 231 - 1
nums 中的所有值都 互不相同
nums 按升序排列
*/
func summaryRanges(nums []int) []string {

	if len(nums) == 0 {
		return []string{}
	}

	var result []string
	start, end := 0, 0
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[end]+1 {
			end++
		} else {
			if nums[start] == nums[end] {
				result = append(result, strconv.Itoa(nums[start]))
			} else {
				result = append(result, fmt.Sprintf("%v->%v", nums[start], nums[end]))
			}
			start, end = i, i
		}
	}

	if nums[start] == nums[end] {
		result = append(result, strconv.Itoa(nums[start]))
	} else {
		result = append(result, fmt.Sprintf("%v->%v", nums[start], nums[end]))
	}

	return result
}

func TestSummaryRanges(t *testing.T) {
	fmt.Println(summaryRanges([]int{-1}))
}

/**
解答错误
31 / 34 个通过的测试用例

官方题解
输入
nums =
[-1]

添加到测试用例
输出
[]
预期结果
["-1"]
*/
