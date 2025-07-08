package topinterview145

/**
给你一个嵌套的整数列表 nestedList 。每个元素要么是一个整数，要么是一个列表；该列表的元素也可能是整数或者是其他列表。请你实现一个迭代器将其扁平化，使之能够遍历这个列表中的所有整数。

实现扁平迭代器类 NestedIterator ：

NestedIterator(List<NestedInteger> nestedList) 用嵌套列表 nestedList 初始化迭代器。
int next() 返回嵌套列表的下一个整数。
boolean hasNext() 如果仍然存在待迭代的整数，返回 true ；否则，返回 false 。
你的代码将会用下述伪代码检测：

initialize iterator with nestedList
res = []
while iterator.hasNext()
    append iterator.next() to the end of res
return res
如果 res 与预期的扁平化列表匹配，那么你的代码将会被判为正确。



示例 1：

输入：nestedList = [[1,1],2,[1,1]]
输出：[1,1,2,1,1]
解释：通过重复调用 next 直到 hasNext 返回 false，next 返回的元素的顺序应该是: [1,1,2,1,1]。
示例 2：

输入：nestedList = [1,[4,[6]]]
输出：[1,4,6]
解释：通过重复调用 next 直到 hasNext 返回 false，next 返回的元素的顺序应该是: [1,4,6]。


提示：

1 <= nestedList.length <= 500
嵌套列表中的整数值在范围 [-106, 106] 内
*/

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedInteger struct {
}

func (n NestedInteger) IsInteger() bool {
	return false
}
func (n NestedInteger) GetInteger() int {
	return 0
}
func (n *NestedInteger) SetInteger(value int) {
}
func (n *NestedInteger) Add(elem NestedInteger) {}
func (n NestedInteger) GetList() []*NestedInteger {
	return nil
}

//------

/*
*
执行用时分布
16
ms
击败
6.06%
复杂度分析
消耗内存分布
9.26
MB
击败
6.06%
复杂度分析

竟然自己完成了
不断递归展开
*/
type NestedIterator struct {
	list []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{
		list: nestedList,
	}
}

func (n *NestedIterator) Next() int {

	if len(n.list) == 0 {
		return 0
	}
	v := n.list[0].GetInteger()
	n.list = append(n.list[0].GetList(), n.list[1:]...)
	n.flat()
	return v
}

func (n *NestedIterator) flat() {
	if len(n.list) == 0 {
		return
	}
	if n.list[0].IsInteger() {
		return
	} else {
		n.list = append(n.list[0].GetList(), n.list[1:]...)
	}
	n.flat()
}

func (n *NestedIterator) HasNext() bool {
	n.flat()
	return len(n.list) > 0
}
