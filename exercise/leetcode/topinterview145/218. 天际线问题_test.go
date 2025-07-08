package topinterview145

import (
	"fmt"
	"sort"
	"testing"
)

/**
城市的 天际线 是从远处观看该城市中所有建筑物形成的轮廓的外部轮廓。给你所有建筑物的位置和高度，请返回 由这些建筑物形成的 天际线 。

每个建筑物的几何信息由数组 buildings 表示，其中三元组 buildings[i] = [lefti, righti, heighti] 表示：

lefti 是第 i 座建筑物左边缘的 x 坐标。
righti 是第 i 座建筑物右边缘的 x 坐标。
heighti 是第 i 座建筑物的高度。
你可以假设所有的建筑都是完美的长方形，在高度为 0 的绝对平坦的表面上。

天际线 应该表示为由 “关键点” 组成的列表，格式 [[x1,y1],[x2,y2],...] ，并按 x 坐标 进行 排序 。关键点是水平线段的左端点。列表中最后一个点是最右侧建筑物的终点，y 坐标始终为 0 ，仅用于标记天际线的终点。此外，任何两个相邻建筑物之间的地面都应被视为天际线轮廓的一部分。

注意：输出天际线中不得有连续的相同高度的水平线。例如 [...[2 3], [4 5], [7 5], [11 5], [12 7]...] 是不正确的答案；三条高度为 5 的线应该在最终输出中合并为一个：[...[2 3], [4 5], [12 7], ...]



示例 1：


输入：buildings = [[2,9,10],[3,7,15],[5,12,12],[15,20,10],[19,24,8]]
输出：[[2,10],[3,15],[7,12],[12,0],[15,10],[20,8],[24,0]]
解释：
图 A 显示输入的所有建筑物的位置和高度，
图 B 显示由这些建筑物形成的天际线。图 B 中的红点表示输出列表中的关键点。
示例 2：

输入：buildings = [[0,2,3],[2,5,3]]
输出：[[0,3],[5,0]]


提示：

1 <= buildings.length <= 104
0 <= lefti < righti <= 231 - 1
1 <= heighti <= 231 - 1
buildings 按 lefti 非递减排序
*/
/**
核心思路：扫描线 + 最大堆（优先队列）

步骤解析：
1. 把所有“关键点”提取出来
每个建筑可以转化为两个事件：
	- 左边界：表示建筑开始 left, -height（负号表示进入）
	- 右边界：表示建筑结束 right, height（正号表示离开）
👉 为什么左边高度取负？这样方便在排序时：
	- 相同 x 时，进入建筑优先于离开建筑
	- 进入时高度大的优先（负数更小）
2. 按 x 从小到大排序这些关键点
3. 用一个最大堆维护当前所有还没结束的建筑的高度
	- 进入点：把高度加入堆
	- 离开点：从堆中移除该高度
4. 每次高度变化，就记录一个关键点
	- 当前最大高度 != 上一次最大高度，就说明 skyline 有变化
	- 记录当前 x 坐标和当前堆顶（最大高度）

📌 小结
步骤	内容
1	把建筑转为关键点：进入点（负高）和离开点（正高）
2	排序规则：先按横坐标，再按高度（负数在前）
3	扫描线+最大堆维护当前高度
4	如果当前最大高度变了 → 记录关键点

*/
/**
看完题解自己来实现

1. 按时间排序
2. 最大堆堆顶存当前最大高度
3. 遇到相同高度移除堆内同高度的元素

删除中间结点后， 判断是否还有子结点
和root元素对比，如果大于则向上调整，否则向下调整
			if i < len(h.arr) {
				root := (i - 1) / 2
				if root >= 0 {
					if h.compare(i, root) {
						h.adjustDownToUp(i)
					} else {
						h.adjustUpToDown(i, len(h.arr))
					}
				}
			}
*/
/**
执行用时分布
49
ms
击败
13.89%
复杂度分析
消耗内存分布
9.89
MB
击败
38.89%
复杂度分析

*/

type maxHeap struct {
	arr [][]int
}

func (h *maxHeap) Push(e []int) {
	h.arr = append(h.arr, e)
	h.adjustDownToUp(len(h.arr) - 1)
}

func (h *maxHeap) Remove(height int) {
	for i, event := range h.arr {
		if event[1] == height {
			h.arr[i], h.arr[len(h.arr)-1] = h.arr[len(h.arr)-1], h.arr[i]
			h.arr = h.arr[:len(h.arr)-1]
			if i < len(h.arr) {
				root := (i - 1) / 2
				if root >= 0 {
					if h.compare(i, root) {
						h.adjustDownToUp(i)
					} else {
						h.adjustUpToDown(i, len(h.arr))
					}
				}
			}
			return
		}
	}
}

func (h *maxHeap) adjustUpToDown(root, len int) {

	index := root*2 + 1
	for index < len {
		if index+1 < len && h.compare(index+1, index) {
			index++
		}
		if h.compare(index, root) {
			h.arr[index], h.arr[root] = h.arr[root], h.arr[index]
		} else {
			break
		}
		root = index
		index = root*2 + 1
	}
}

func (h *maxHeap) adjustDownToUp(index int) {
	root := (index - 1) / 2
	for root >= 0 {
		if h.compare(index, root) {
			h.arr[index], h.arr[root] = h.arr[root], h.arr[index]
		} else {
			break
		}
		index = root
		root = (index - 1) / 2
	}
}

func (h *maxHeap) compare(i, j int) bool {
	return h.arr[i][1] > h.arr[j][1]
}

func (h *maxHeap) Top() []int {
	if len(h.arr) == 0 {
		return []int{0, 0}
	}
	return h.arr[0]
}

func getSkyline(buildings [][]int) [][]int {

	//事件
	events := make([][]int, len(buildings)*2)
	for i, b := range buildings {
		events[i*2] = []int{b[0], -b[2]}
		events[i*2+1] = []int{b[1], b[2]}
	}
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0] || (events[i][0] == events[j][0] && events[i][1] < events[j][1])
	})
	var ans [][]int
	h := &maxHeap{}
	preHeight := 0
	for _, event := range events {
		if event[1] < 0 {
			h.Push([]int{event[0], -event[1]})
		} else {
			h.Remove(event[1])
		}
		if h.Top()[1] != preHeight {
			preHeight = h.Top()[1]
			ans = append(ans, []int{event[0], preHeight})
		}
	}
	return ans
}

func TestGetSkyline(t *testing.T) {

	fmt.Println(getSkyline([][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}))
	fmt.Println(getSkyline([][]int{{0, 2, 3}, {2, 5, 3}}))
}
