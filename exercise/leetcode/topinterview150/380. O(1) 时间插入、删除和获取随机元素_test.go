package topinterview150

import (
	"fmt"
	"math/rand"
	"testing"
)

/*
*
实现RandomizedSet 类：

RandomizedSet() 初始化 RandomizedSet 对象
bool insert(int val) 当元素 val 不存在时，向集合中插入该项，并返回 true ；否则，返回 false 。
bool remove(int val) 当元素 val 存在时，从集合中移除该项，并返回 true ；否则，返回 false 。
int getRandom() 随机返回现有集合中的一项（测试用例保证调用此方法时集合中至少存在一个元素）。每个元素应该有 相同的概率 被返回。
你必须实现类的所有函数，并满足每个函数的 平均 时间复杂度为 O(1) 。
*/
type RandomizedSet struct {
	m   map[int]int
	arr []int
}

func Constructor() RandomizedSet {

	return RandomizedSet{
		m: make(map[int]int),
	}
}

func (s *RandomizedSet) Insert(val int) bool {
	if _, ok := s.m[val]; ok {
		return false
	}
	s.m[val] = len(s.arr)
	s.arr = append(s.arr, val)
	return true
}

func (s *RandomizedSet) Remove(val int) bool {
	if _, ok := s.m[val]; !ok {
		return false
	}
	// 这样写第一次删除之后，所有下标的错位了，跟map的对不上了
	//s.arr = append(s.arr[0:s.m[val]], s.arr[s.m[val]+1:]...)

	// 跟最后一个交换位置！！！
	last := len(s.arr) - 1
	s.arr[s.m[val]] = s.arr[last]
	s.m[s.arr[last]] = s.m[val]
	s.arr = s.arr[:last]
	delete(s.m, val)
	return true
}

func (s *RandomizedSet) GetRandom() int {
	/*if len(s.arr) == 0 {
		return 0
	}
	rand.New(rand.NewSource(time.Now().UnixNano()))
	index := rand.Intn(len(s.arr))
	return s.arr[index]*/
	return s.arr[rand.Intn(len(s.arr))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

/**

解答错误
5 / 21 个通过的测试用例

官方题解
输入
["RandomizedSet","insert","insert","remove","insert","remove","getRandom"]
[[],[0],[1],[0],[2],[1],[]]

添加到测试用例
输出
[null,true,true,true,true,true,1]
预期结果
[null,true,true,true,true,true,2]
*/

func TestRandomizedSet(t *testing.T) {
	obj := Constructor()
	obj.Insert(1)
	obj.Insert(3)
	obj.Insert(2)
	obj.Remove(1)
	obj.Remove(3)
	fmt.Println(obj.GetRandom())
}

/**
执行用时分布
56
ms
击败
83.63%
复杂度分析
消耗内存分布
50.07
MB
击败
63.05%
复杂度分析

*/
