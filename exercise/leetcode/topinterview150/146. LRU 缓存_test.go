package topinterview150

import (
	"fmt"
	"testing"
)

/**
看了题解，再自己写

hashmap保存keyval
双向链表保存数量，队头为最近访问，用于后续淘汰

执行用时分布
95
ms
击败
31.73%
复杂度分析
消耗内存分布
94.67
MB
击败
6.28%
复杂度分析


*/

type LRUCache struct {
	data     map[int]*LinkedNode
	head     *LinkedNode
	tail     *LinkedNode
	count    int
	capacity int
}

type LinkedNode struct {
	key  int
	val  int
	pre  *LinkedNode
	next *LinkedNode
}

func Constructor1(capacity int) LRUCache {
	return LRUCache{
		data:     make(map[int]*LinkedNode),
		capacity: capacity,
	}
}

func (lru *LRUCache) Get(key int) int {
	if node, ok := lru.data[key]; ok {
		lru.putNodeToHead(node)
		return node.val
	}
	return -1
}

func (lru *LRUCache) putNodeToHead(node *LinkedNode) {
	if node.next == nil && node.pre != nil {
		lru.tail = node.pre
	}
	if node.pre != nil {
		node.pre.next = node.next
		if node.next != nil {
			node.next.pre = node.pre
		}
		lru.head.pre = node
		node.next = lru.head
		lru.head = node
		lru.head.pre = nil
	}
}

func (lru *LRUCache) Put(key int, value int) {

	if lru.capacity == 0 {
		return
	}

	if node, ok := lru.data[key]; ok {
		node.val = value
		lru.putNodeToHead(node)
	} else {
		newNode := &LinkedNode{key: key, val: value}
		lru.data[key] = newNode
		if lru.head == nil {
			lru.head = newNode
			lru.tail = newNode
			lru.count = 1
		} else {
			lru.head.pre = newNode
			newNode.next = lru.head
			lru.head = newNode
			lru.count++
			if lru.count > lru.capacity {
				delete(lru.data, lru.tail.key)
				lru.count--
				if lru.tail.pre != nil {
					lru.tail.pre.next = nil
				}
				lru.tail = lru.tail.pre
			}
		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func TestLRUCache(t *testing.T) {
	obj := Constructor1(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	obj.Put(3, 3)
	fmt.Println(obj.Get(2))
	obj.Put(4, 4)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(4))

}
