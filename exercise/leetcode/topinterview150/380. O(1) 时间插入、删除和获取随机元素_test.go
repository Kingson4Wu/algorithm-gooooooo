package topinterview150

import (
	"math/rand"
	"time"
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
	s.arr = append(s.arr[0:s.m[val]], s.arr[s.m[val]+1:]...)
	delete(s.m, val)
	return true
}

func (s *RandomizedSet) GetRandom() int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	index := rand.Intn(len(s.arr))
	return s.arr[index]
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
