package LCR

type CQueue struct {
	arr []int
}

func Constructor() CQueue {

	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {

	this.arr = append(this.arr, value)
}

func (this *CQueue) DeleteHead() int {

	if len(this.arr) == 0 {
		return -1
	}
	v := this.arr[0]
	this.arr = this.arr[1:]
	return v
}

/**
很无聊的题目
就是栈

执行用时分布
46
ms
击败
96.77%
复杂度分析
消耗内存分布
10.23
MB
击败
62.90%

*/

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
