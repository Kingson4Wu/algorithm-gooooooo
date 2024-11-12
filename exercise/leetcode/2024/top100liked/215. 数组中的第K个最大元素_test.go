package top100liked

func findKthLargest(nums []int, k int) int {

	return 0
}

/**
调整方法（寻找左右节点最大的那个跟根节点交换），自上而下

1.初始化堆
从最后一个根节点开始遍历，自下而上调调整方法
2.遍历排序（TopK）
第一个根节点开始遍历到根节点，和最后一个节点交换，再调调整方法
*/
