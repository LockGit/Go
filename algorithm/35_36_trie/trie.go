/**
* Created by lock
* trie树，字典数
这个是简版的demo，练习使用，在实际场景中用的是加强版
*/
package main

import (
	"fmt"
	"sync"
)

type Bytes []byte

type trieNode struct {
	children map[byte]*trieNode
	symbol   byte
	value    []byte
	root     bool
}

type Trie struct {
	rw   sync.RWMutex
	root *trieNode
	size int
}

func NewTrie() *Trie {
	return &Trie{
		root: &trieNode{root: true, children: make(map[byte]*trieNode)},
		size: 1,
	}
}

func newNode(symbol byte) *trieNode {
	return &trieNode{children: make(map[byte]*trieNode), symbol: symbol}
}

func (t *Trie) Insert(key, value Bytes) {
	t.rw.Lock()
	defer t.rw.Unlock()
	currNode := t.root
	for _, symbol := range key {
		if currNode.children[symbol] == nil {
			currNode.children[symbol] = newNode(symbol)
		}
		currNode = currNode.children[symbol]
	}
	if currNode.value == nil {
		t.size++
	}
	currNode.value = value
}

func (t *Trie) Search(key Bytes) (Bytes, bool) {
	t.rw.RLock()
	defer t.rw.RUnlock()
	currNode := t.root
	for _, symbol := range key {
		if currNode.children[symbol] == nil {
			return nil, false
		}
		currNode = currNode.children[symbol]
	}
	return currNode.value, true
}

func main() {
	tree := NewTrie()
	tree.Insert([]byte("a"), []byte("1"))
	tree.Insert([]byte("ab"), []byte("12"))
	res, _ := tree.Search([]byte("ab"))
	fmt.Println(string(res))
}
