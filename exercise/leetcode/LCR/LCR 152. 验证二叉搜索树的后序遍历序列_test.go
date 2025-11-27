package LCR

import (
	"fmt"
	"testing"
)

// 请实现一个函数来判断整数数组 postorder 是否为二叉搜索树的后序遍历结果。

/**
看完题解自己实现的


如何判断序列合法？
步骤：

取最后一个元素 root

从左往右扫描，找到第一个 > root 的位置 index

index 之前全是左子树

index 到倒数第二个是右子树

右子树部分不能出现小于 root 的元素（必须都 > root）

递归判断左右子数组本身也要是合法后序遍历

🎯关键点一句话总结

找分界点 → 验右侧是否全大于 root → 对左右子数组递归做同样的事。

执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
3.75
MB
击败
95.95%

*/

func verifyTreeOrder(postorder []int) bool {
	if len(postorder) <= 1 {
		return true
	}
	index := 0
	for i := 0; i < len(postorder)-1; i++ {
		if postorder[i] < postorder[len(postorder)-1] {
			index++
		} else {
			break
		}
	}
	// 判断右边是不是都大于root
	for i := index; i < len(postorder)-1; i++ {
		if postorder[i] < postorder[len(postorder)-1] {
			return false
		}
	}
	return verifyTreeOrder(postorder[:index]) && verifyTreeOrder(postorder[index:len(postorder)-1])
}

func TestVerifyTreeOrder(t *testing.T) {
	fmt.Println(verifyTreeOrder([]int{4, 9, 6, 5, 8}))
}
