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

/**
A = 65
a = 97

'a' - 'A' = 32 !!!!
不是26 ！！！
*/

/**
时间
60 ms
击败
20.74%
内存
47.4 MB
击败
5.5%
*/

/**
提示：

1 <= word.length, prefix.length <= 2000
word 和 prefix 仅由小写英文字母组成 !!!!
insert、search 和 startsWith 调用次数 总计 不超过 3 * 104 次
*/

type Trie struct {
	words [26]*Trie
	isEnd bool
}

func Constructor() Trie {

	return Trie{}
}

func (t *Trie) Insert(word string) {

	root := t
	for _, ch := range word {
		index := ch - 'a'
		trie := root.words[index]
		if trie == nil {
			/*node := Constructor()
			trie = &node
			root.words[index] = trie*/
			root.words[index] = &Trie{}
		}
		/** 最后一个字符标志位结尾 */
		/*if i == len(word)-1 {
			trie.isEnd = true
		}*/
		//root = trie
		root = root.words[index]
	}
	root.isEnd = true
}

func (t *Trie) Search(word string) bool {

	root := t
	for _, ch := range word {
		index := ch - 'a'
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
		index := ch - 'a'
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
	/*trie.Insert("apple")
	fmt.Println(trie.Search("apple"))
	fmt.Println(trie.Search("app"))
	fmt.Println(trie.StartsWith("app"))
	trie.Insert("app")
	fmt.Println(trie.Search("app"))*/

	trie.Insert("p")
	fmt.Println(trie.StartsWith("pr"))
	fmt.Println(trie.Search("p"))
	trie.Insert("pr")
	fmt.Println(trie.StartsWith("pre"))
	fmt.Println(trie.Search("pr"))
	trie.Insert("pre")
	fmt.Println(trie.StartsWith("pre"))
	fmt.Println(trie.Search("pre"))
	trie.Insert("pref")
	fmt.Println(trie.StartsWith("pref"))
	fmt.Println(trie.Search("pref"))
	trie.Insert("prefi")
	fmt.Println(trie.StartsWith("pref"))
	fmt.Println(trie.Search("prefi"))
	trie.Insert("prefix")
	fmt.Println(trie.StartsWith("prefi"))
	fmt.Println(trie.Search("prefix"))

}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

/**
panic: runtime error: index out of range [55] with length 52
main.(*Trie).Insert(...)
solution.go, line 19
main.(*DriverSolution).Helper_Select_Method(0x4b6460?, {0xc000014cb6?, 0xa?}, {0xc000011620?, 0x49ecb8?, 0xc00007e100?}, 0xc00005cbe0)
solution.go, line 66
main.(*DriverSolution).Helper(0xc0000920d0?, {0xc00007adc0, 0x13, 0x4b2d60?}, {0xc000098000, 0x13, 0xc000014038?})
solution.go, line 129
main.main()
solution.go, line 174
最后执行的输入
添加到测试用例
["Trie","insert","startsWith","search","insert","startsWith","search","insert","startsWith","search","insert","startsWith","search","insert","startsWith","search","insert","startsWith","search"]
[[],["p"],["pr"],["p"],["pr"],["pre"],["pr"],["pre"],["pre"],["pre"],["pref"],["pref"],["pref"],["prefi"],["pref"],["prefi"],["prefix"],["prefi"],["prefix"]]
*/
