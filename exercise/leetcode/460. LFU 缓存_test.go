package leetcode

import (
	"fmt"
	"testing"
)

/**
请你为 最不经常使用（LFU）缓存算法设计并实现数据结构。

实现 LFUCache 类：

LFUCache(int capacity) - 用数据结构的容量 capacity 初始化对象
int get(int key) - 如果键 key 存在于缓存中，则获取键的值，否则返回 -1 。
void put(int key, int value) - 如果键 key 已存在，则变更其值；如果键不存在，请插入键值对。当缓存达到其容量 capacity 时，则应该在插入新项之前，移除最不经常使用的项。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，应该去除 最近最久未使用 的键。
为了确定最不常使用的键，可以为缓存中的每个键维护一个 使用计数器 。使用计数最小的键是最久未使用的键。

当一个键首次插入到缓存中时，它的使用计数器被设置为 1 (由于 put 操作)。对缓存中的键执行 get 或 put 操作，使用计数器的值将会递增。

函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。



示例：

输入：
["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]
输出：
[null, null, null, 1, null, -1, 3, null, -1, 3, 4]

解释：
// cnt(x) = 键 x 的使用计数
// cache=[] 将显示最后一次使用的顺序（最左边的元素是最近的）
LFUCache lfu = new LFUCache(2);
lfu.put(1, 1);   // cache=[1,_], cnt(1)=1
lfu.put(2, 2);   // cache=[2,1], cnt(2)=1, cnt(1)=1
lfu.get(1);      // 返回 1
                 // cache=[1,2], cnt(2)=1, cnt(1)=2
lfu.put(3, 3);   // 去除键 2 ，因为 cnt(2)=1 ，使用计数最小
                 // cache=[3,1], cnt(3)=1, cnt(1)=2
lfu.get(2);      // 返回 -1（未找到）
lfu.get(3);      // 返回 3
                 // cache=[3,1], cnt(3)=2, cnt(1)=2
lfu.put(4, 4);   // 去除键 1 ，1 和 3 的 cnt 相同，但 1 最久未使用
                 // cache=[4,3], cnt(4)=1, cnt(3)=2
lfu.get(1);      // 返回 -1（未找到）
lfu.get(3);      // 返回 3
                 // cache=[3,4], cnt(4)=1, cnt(3)=3
lfu.get(4);      // 返回 4
                 // cache=[3,4], cnt(4)=2, cnt(3)=3


提示：

0 <= capacity <= 104
0 <= key <= 105
0 <= value <= 109
最多调用 2 * 105 次 get 和 put 方法
*/

/*
*
方法二：双哈希 + 双向链表（时间复杂度 O(1)）
keyToValFreq: map[key] => (value, freq)
freqToKeys: map[freq] => keys 双向链表(按 LRU 顺序存 key)
minFreq: 当前所有 key 中的最小 freq
capacity: 缓存容量

Put:若 key 已存在：- 更新 value !!!

---

### 📦 操作逻辑：

#### 1. `get(key)`

* 如果 key 不存在，返回 -1
* 否则：
  - 获取当前 `val, freq`
  - 从 `freqToKeys[freq]` 中移除该 key
  - freq += 1，并加入 `freqToKeys[freq+1]`
  - 如果 freqToKeys\[freq] 为空且 freq 是 minFreq，则 minFreq++
  - 更新 key 的频率

#### 2. `put(key, value)`
* 如果 `capacity == 0`，直接返回
* 若 key 已存在：
  - 更新 value（同 get 一样更新 freq）!!!!

* 若 key 不存在：
  - 如果当前缓存已满：
    *- 从 freqToKeys\[minFreq] 中 **删除最旧的 key**
    *- 同时从 keyToValFreq 中删掉
  - 插入新 key，freq = 1，更新 minFreq = 1

---

### 📌 核心思路小结：

| 操作          | 时间复杂度 |
| ----------- | ----- |
| get(key)    | O(1)  |
| put(key, v) | O(1)  |

* 利用双哈希表管理 key 到 value/freq 映射，以及 freq 到 LRU key 列表。
* 淘汰策略优先考虑频率，再看访问顺序（LRU）。
*/
type LFUCache struct {
	keyToValFreq   map[int]*LFUNode
	freqToKeysHead map[int]*LFUNode
	freqToKeysTail map[int]*LFUNode
	minFreq        int
	capacity       int
	size           int
}

type LFUNode struct {
	key  int
	val  int
	freq int
	pre  *LFUNode
	next *LFUNode
}

func Constructor(capacity int) LFUCache {

	return LFUCache{
		capacity:       capacity,
		keyToValFreq:   make(map[int]*LFUNode),
		freqToKeysHead: make(map[int]*LFUNode),
		freqToKeysTail: make(map[int]*LFUNode),
	}
}

func (l *LFUCache) access(node *LFUNode) {
	// 删除
	if node.pre == nil && node.next == nil {
		l.freqToKeysHead[node.freq] = nil
		l.freqToKeysTail[node.freq] = nil
		if node.freq == l.minFreq {
			l.minFreq = node.freq + 1
		}
	} else if node.pre != nil && node.next != nil {
		node.pre.next = node.next
		node.next.pre = node.pre
	} else if node.pre == nil {
		l.freqToKeysHead[node.freq] = node.next
		node.next.pre = nil
	} else {
		l.freqToKeysTail[node.freq] = node.pre
		node.pre.next = nil
	}
	node.freq++
	// 新增
	if l.freqToKeysTail[node.freq] == nil {
		l.freqToKeysHead[node.freq] = node
		l.freqToKeysTail[node.freq] = node
		node.next = nil
		node.pre = nil
	} else {
		node.pre = l.freqToKeysTail[node.freq]
		l.freqToKeysTail[node.freq].next = node
		l.freqToKeysTail[node.freq] = node
		node.next = nil
	}
}

func (l *LFUCache) Get(key int) int {

	if node, ok := l.keyToValFreq[key]; ok {
		l.access(node)
		return node.val
	}
	return -1
}

func (l *LFUCache) Put(key int, value int) {

	if l.capacity == 0 {
		return
	}

	if node, ok := l.keyToValFreq[key]; ok {
		// 别漏了更新值
		if node.val != value {
			node.val = value
		}
		l.access(node)
	} else {

		if l.size == l.capacity {
			oldNode := l.freqToKeysHead[l.minFreq]
			if oldNode == l.freqToKeysTail[l.minFreq] {
				l.freqToKeysHead[l.minFreq] = nil
				l.freqToKeysTail[l.minFreq] = nil
			} else {
				l.freqToKeysHead[l.minFreq] = oldNode.next
				oldNode.next.pre = nil
			}
			delete(l.keyToValFreq, oldNode.key)
			l.size--
		}

		node := &LFUNode{
			key:  key,
			freq: 1,
			val:  value,
		}
		l.keyToValFreq[key] = node
		l.minFreq = 1
		l.size++

		// 新增
		if l.freqToKeysTail[node.freq] == nil {
			l.freqToKeysHead[node.freq] = node
			l.freqToKeysTail[node.freq] = node
			node.next = nil
			node.pre = nil
		} else {
			node.pre = l.freqToKeysTail[node.freq]
			l.freqToKeysTail[node.freq].next = node
			l.freqToKeysTail[node.freq] = node
			node.next = nil
		}
	}
}

/**
解答错误
13 / 25 个通过的测试用例

官方题解
输入
["LFUCache","put","put","put","put","get"]
[[2],[3,1],[2,1],[2,2],[4,4],[2]]

添加到测试用例
输出
[null,null,null,null,null,1]
预期结果
[null,null,null,null,null,2]
*/

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

/**
执行用时分布
120
ms
击败
28.87%
复杂度分析
消耗内存分布
74.89
MB
击败
65.27%

*/

func TestLFUCache(t *testing.T) {
	obj := Constructor(2)
	obj.Put(3, 1)
	obj.Put(2, 1)
	obj.Put(2, 2)
	obj.Put(4, 4)
	fmt.Println(obj.Get(2))
}
