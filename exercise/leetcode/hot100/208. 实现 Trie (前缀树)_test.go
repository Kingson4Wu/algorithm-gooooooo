package hot100

import (
	"fmt"
	"testing"
)

/*
*
Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补完和拼写检查。

请你实现 Trie 类：

Trie() 初始化前缀树对象。
void insert(String word) 向前缀树中插入字符串 word 。
boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回 false 。
boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否则，返回 false 。

示例：

输入
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
输出
[null, null, true, false, true, null, true]

解释
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // 返回 True
trie.search("app");     // 返回 False
trie.startsWith("app"); // 返回 True
trie.insert("app");
trie.search("app");     // 返回 True
*/

/**
看过题解，理解完比较简单
1、Trie 使用一个数组记录全部字母对应的下一个Trie， 还有一个标志位记录是否结尾；
（其实就是一个树，根节点是一个52个字母(大小写)组成的数组 ！！）
2、利用ch - 'A' < 26 的特性
*/

type Trie struct {
	words []*Trie
	isEnd bool
}

func Constructor() Trie {

	return Trie{
		words: make([]*Trie, 52),
		isEnd: false,
	}
}

func (t *Trie) Insert(word string) {

	root := t
	for i, ch := range word {
		index := ch - 'A'
		trie := root.words[index]
		if trie == nil {
			node := Constructor()
			trie = &node
			root.words[index] = trie
		}
		/** 最后一个字符标志位结尾 */
		if i == len(word)-1 {
			trie.isEnd = true
		}
		root = trie
	}
}

func (t *Trie) Search(word string) bool {

	root := t
	for _, ch := range word {
		index := ch - 'A'
		trie := root.words[index]
		if trie == nil {
			return false
		}
		root = trie
	}
	return root.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {

	root := t
	for _, ch := range prefix {
		index := ch - 'A'
		trie := root.words[index]
		if trie == nil {
			return false
		}
		root = trie
	}
	return true
}

func TestTrie(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))
	fmt.Println(trie.Search("app"))
	fmt.Println(trie.StartsWith("app"))
	trie.Insert("app")
	fmt.Println(trie.Search("app"))

}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
