package _025

type MinStack struct {
	stack  []int
	mStack []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(val int) {

	s.stack = append(s.stack, val)
	if len(s.mStack) == 0 {
		s.mStack = append(s.mStack, val)
	} else {
		s.mStack = append(s.mStack, min(val, s.mStack[len(s.mStack)-1]))
	}

}

func (s *MinStack) Pop() {
	s.stack = s.stack[:len(s.stack)-1]
	s.mStack = s.mStack[:len(s.mStack)-1]

}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]

}

func (s *MinStack) GetMin() int {
	return s.mStack[len(s.mStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

/**
执行用时分布
6
ms
击败
26.41%
复杂度分析
消耗内存分布
9.92
MB
击败
10.41%
*/

/**
在栈的基础上，增加同样大小的辅助栈，用于记录当前的最小值
相当于每次push的时候，都记录当前快照下的最小值
*/
