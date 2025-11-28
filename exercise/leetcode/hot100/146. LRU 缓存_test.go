package hot100

import (
	"fmt"
	"testing"
)

/**
请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。



示例：

输入
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
输出
[null, null, null, 1, null, -1, null, -1, 3, 4]

解释
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // 缓存是 {1=1}
lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
lRUCache.get(1);    // 返回 1
lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
lRUCache.get(2);    // 返回 -1 (未找到)
lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
lRUCache.get(1);    // 返回 -1 (未找到)
lRUCache.get(3);    // 返回 3
lRUCache.get(4);    // 返回 4


提示：

1 <= capacity <= 3000
0 <= key <= 10000
0 <= value <= 105
最多调用 2 * 105 次 get 和 put
*/

/*
*
执行用时分布
441
ms
击败
53.26%
复杂度分析
消耗内存分布
74.29
MB
击败
84.78%

凭回忆能做出来，不过要debug才能全部把细节调好
*/
type LruNode struct {
	k    int
	v    int
	pre  *LruNode
	next *LruNode
}

type LRUCache struct {
	m        map[int]*LruNode
	head     *LruNode
	tail     *LruNode
	capacity int
	size     int
}

func Constructor4(capacity int) LRUCache {

	return LRUCache{
		capacity: capacity,
		m:        map[int]*LruNode{},
	}
}

func (this *LRUCache) moveToHead(node *LruNode) {
	//新节点
	if node.next == nil && node.pre == nil {
		if this.size == 0 {
			this.head = node
			this.tail = node
			return
		}
		this.head.pre = node
		node.next = this.head
		this.head = node
		return
	}
	//队首
	if node == this.head {
		return
	}
	//队尾
	if node == this.tail {
		this.tail = this.tail.pre
		this.tail.next = nil
		node.pre = nil
		node.next = this.head
		this.head.pre = node
		this.head = node
		return
	}
	node.pre.next = node.next
	node.next.pre = node.pre
	node.next = this.head
	this.head.pre = node
	this.head = node
}

func (this *LRUCache) removeTail() {
	delete(this.m, this.tail.k)
	if this.tail.pre != nil {
		this.tail.pre.next = nil
	}
	this.tail = this.tail.pre
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.m[key]; ok {
		this.moveToHead(node)
		return node.v
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {

	if node, ok := this.m[key]; ok {
		node.v = value // //这句忘记了！！！！
		this.moveToHead(node)
		return
	}
	this.m[key] = &LruNode{
		k: key,
		v: value,
	}
	this.moveToHead(this.m[key])

	if this.size == this.capacity {
		this.removeTail()
	} else {
		this.size++
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func TestLRUCache(t *testing.T) {
	obj := Constructor4(2)
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
