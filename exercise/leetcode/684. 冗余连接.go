package leetcode

/*
*
树可以看成是一个连通且 无环 的 无向 图。

给定往一棵 n 个节点 (节点值 1～n) 的树中添加一条边后的图。添加的边的两个顶点包含在 1 到 n 中间，且这条附加的边不属于树中已存在的边。图的信息记录于长度为 n 的二维数组 edges ，edges[i] = [ai, bi] 表示图中在 ai 和 bi 之间存在一条边。

请找出一条可以删去的边，删除后可使得剩余部分是一个有着 n 个节点的树。如果有多个答案，则返回数组 edges 中最后出现的边。
*/
func findRedundantConnection(edges [][]int) []int {

	parent := make(map[int]int)
	/** 先记录初始连通关系， 自己和自己连通*/
	for i := 1; i <= len(edges); i++ {
		parent[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			//递归找到最终的parent
			//find不会有死循环问题，因为形成环之前，使用union连接时会判断是否已经连通，没机会形成环
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	/** 将两个点连在一起 */
	union := func(from, to int) bool {
		x, y := find(from), find(to)
		if x == y {
			return false
		}
		parent[x] = find(parent[y])
		return true
	}
	for _, e := range edges {
		if !union(e[0], e[1]) {
			return e
		}
	}
	return nil
}

/**
执行用时分布
2
ms
击败
84.21%
复杂度分析
消耗内存分布
4.78
MB
击败
15.79%
*/
