package interview

// TrieNode 定义前缀树（Trie）的节点
type TrieNode struct {
	children map[rune]*TrieNode
	count    int // 记录以该节点为前缀的单词数量
	endCount int // 记录以该节点为结束的单词数量
}

// Trie 定义前缀树结构
type Trie struct {
	root *TrieNode
}

// WordStore 定义存储结构 WordStore，包含哈希集合和前缀树
type WordStore struct {
	words map[string]bool // 哈希集合用于快速查找单词是否存在
	trie  *Trie           // 前缀树用于统计前缀匹配数量
}

// NewTrieNode 初始化一个 TrieNode
func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		count:    0,
		endCount: 0,
	}
}

// NewTrie 初始化一个 Trie
func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

// Insert 向 Trie 中插入一个单词
func (t *Trie) Insert(word string) {
	node := t.root
	for _, char := range word {
		if _, ok := node.children[char]; !ok {
			node.children[char] = NewTrieNode()
		}
		node = node.children[char]
		node.count++
	}
	node.endCount++ // 最后一个字符节点，表示有单词以此结尾
}

// CountWordsWithPrefix 返回 Trie 中指定前缀的单词数量
func (t *Trie) CountWordsWithPrefix(prefix string) int {
	node := t.root
	for _, char := range prefix {
		if _, ok := node.children[char]; !ok {
			return 0
		}
		node = node.children[char]
	}
	return node.count
}

// CountWordsExactly 获取指定单词出现次数
func (t *Trie) CountWordsExactly(word string) int {
	node := t.root
	for _, char := range word {
		if _, ok := node.children[char]; !ok {
			return 0
		}
		node = node.children[char]
	}
	return node.endCount
}

// NewWordStore 初始化 WordStore
func NewWordStore() *WordStore {
	return &WordStore{
		words: make(map[string]bool),
		trie:  NewTrie(),
	}
}

// AddWord 向 WordStore 中添加一个单词
func (ws *WordStore) AddWord(word string) {
	if !ws.words[word] {
		ws.words[word] = true
		ws.trie.Insert(word)
	}
}

// Contains 判断单词是否存在于 WordStore 中
func (ws *WordStore) Contains(word string) bool {
	return ws.words[word]
}

// CountWordsWithPrefix 返回 WordStore 中指定前缀的单词数量
func (ws *WordStore) CountWordsWithPrefix(prefix string) int {
	return ws.trie.CountWordsWithPrefix(prefix)
}
