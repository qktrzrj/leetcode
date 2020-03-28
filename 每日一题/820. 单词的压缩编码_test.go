package test17_16

import (
	"fmt"
	"sort"
	"testing"
)

// 给定一个单词列表，我们将这个列表编码成一个索引字符串 S 与一个索引列表 A。
//
//例如，如果这个列表是 ["time", "me", "bell"]，我们就可以将其表示为 S = "time#bell#" 和 indexes = [0, 2, 5]。
//
//对于每一个索引，我们可以通过从字符串 S 中索引的位置开始读取字符串，直到 "#" 结束，来恢复我们之前的单词列表。
//
//那么成功对给定单词列表进行编码的最小字符串长度是多少呢？
//
//
//
//示例：
//
//输入: words = ["time", "me", "bell"]
//输出: 10
//说明: S = "time#bell#" ， indexes = [0, 2, 5] 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/short-encoding-of-words
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
//
// 解题思路：将索引列表中，字符串的公共部分拼接，得到最短的索引字符串
// 构建一个trie树，最后遍历trie树得到各字符高度总和

type Trie struct {
	data     rune
	children [26]*Trie
	level    int
}

var RootTrie = &Trie{data: '/'}

func NewTrie(data rune, level int) *Trie {
	return &Trie{data: data, level: level}
}

func insert(text []rune) int {
	p := RootTrie
	high := 0
	for i := len(text) - 1; i >= 0; i-- {
		t := text[i]
		index := t - 'a'
		if p.children[index] == nil {
			high = p.level + 1
			node := NewTrie(t, high)
			p.children[index] = node
		}
		p = p.children[index]
	}
	if high > 0 {
		high++
	}
	return high
}

type sortStrings []string

func (s sortStrings) Len() int {
	return len(s)
}

func (s sortStrings) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func (s sortStrings) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func minimumLengthEncoding(words []string) int {
	min := 0
	var w sortStrings = words
	sort.Sort(w)
	for i := len(w) - 1; i >= 0; i-- {
		min += insert([]rune(w[i]))
	}
	return min
}

func Test_minimumLengthEncoding(t *testing.T) {
	fmt.Println(minimumLengthEncoding([]string{"me", "time"}))
}
