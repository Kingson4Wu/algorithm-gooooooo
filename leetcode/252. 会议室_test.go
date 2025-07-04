package leetcode

import "sort"

/**
给你一个会议时间安排的数组 intervals，每个会议时间都用一个长度为 2 的数组表示 intervals[i] = [start, end]，表示会议开始于 start，结束于 end。
请你判断一个人是否能够参加所有会议（即，任意两个会议时间不能重叠）。

*/

func canAttendMeetings(intervals [][]int) bool {
	// 按开始时间排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 遍历检测是否有重叠
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}
	return true
}
