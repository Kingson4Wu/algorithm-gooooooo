package interview

import (
	"fmt"
	"testing"
)

func TestNewTrie(t *testing.T) {
	// 初始化 WordStore
	store := NewWordStore()

	// 添加一些单词
	store.AddWord("hello")
	store.AddWord("hell")
	store.AddWord("heaven")
	store.AddWord("heavy")

	// 检查单词是否存在
	fmt.Println("Contains 'hello':", store.Contains("hello"))     // 输出: true
	fmt.Println("Contains 'goodbye':", store.Contains("goodbye")) // 输出: false

	// 前缀查询
	fmt.Println("Count words with prefix 'he':", store.CountWordsWithPrefix("he"))     // 输出: 4
	fmt.Println("Count words with prefix 'hell':", store.CountWordsWithPrefix("hell")) // 输出: 2
	fmt.Println("Count words with prefix 'hea':", store.CountWordsWithPrefix("hea"))   // 输出: 2
}
