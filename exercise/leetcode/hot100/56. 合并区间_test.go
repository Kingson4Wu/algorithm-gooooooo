package hot100

import (
	"fmt"
	"sort"
	"testing"
)

/**
自己做的
1、先按左边排序
2、在遍历比对区间是否有交集做处理

时间
20 ms
击败
65.50%
内存
6.6 MB
击败
93.9%
*/

/*
*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

示例 1：

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2：

输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。

提示：

1 <= intervals.length <= 104
intervals[i].length == 2
0 <= starti <= endi <= 104
*/

type rangeArr [][]int

func (a rangeArr) Len() int           { return len(a) }
func (a rangeArr) Less(i, j int) bool { return a[i][0] < a[j][0] }
func (a rangeArr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func merge(intervals [][]int) [][]int {

	if len(intervals) <= 0 {
		return intervals
	}

	sort.Sort(rangeArr(intervals))

	var result [][]int
	start, end := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][1] <= end {
			continue
		}
		if intervals[i][0] > end {
			result = append(result, []int{start, end})
			start, end = intervals[i][0], intervals[i][1]
			continue
		}
		end = intervals[i][1]
	}
	result = append(result, []int{start, end})

	return result
}

func TestMerge(t *testing.T) {
	fmt.Println(merge([][]int{{1, 3}, {15, 18}, {2, 6}, {8, 10}}))
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))
	fmt.Println(merge([][]int{{1, 1}, {1, 5}}))
	fmt.Println(merge([][]int{{4, 5}, {6, 8}, {3, 9}}))
}
