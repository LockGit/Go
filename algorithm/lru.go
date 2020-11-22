/**
 * Created by Vim
 * User: lock
 * Date: 2019/2/18
 * Time: 18:33
 * 如果数据最近被访问过，那么将来被访问的几率也更高（访问后，把当前节点调整到链表头，新加入的Cache直接加到链表头中）。
 * 对应的还有另一种淘汰策略（如果访问的频率越高，将来被访问的几率更高，需要一个计数器，计数器排序后，调整链表位置，淘汰无用Cache）
 */
package main

import "fmt"

type LruLinkNode struct {
	Key   int
	Value int
	pre   *LruLinkNode
	next  *LruLinkNode
}

type LRUCache struct {
	limit   int
	HashMap map[int]*LruLinkNode
	head    *LruLinkNode
	end     *LruLinkNode
}

func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{limit: capacity}
	lruCache.HashMap = make(map[int]*LruLinkNode, capacity)
	return lruCache
}

func (l *LRUCache) Get(key int) int {
	if v, ok := l.HashMap[key]; ok {
		l.refreshNode(v)
		return v.Value
	} else {
		return -1
	}
}

func (l *LRUCache) Put(key int, value int) {
	if v, ok := l.HashMap[key]; !ok {
		if len(l.HashMap) >= l.limit {
			oldKey := l.removeNode(l.head)
			delete(l.HashMap, oldKey)
		}
		node := LruLinkNode{Key: key, Value: value}
		l.addNode(&node)
		l.HashMap[key] = &node
	} else {
		v.Value = value
		l.refreshNode(v)
	}
}

func (l *LRUCache) refreshNode(node *LruLinkNode) {
	if node == l.end {
		return
	}
	l.removeNode(node)
	l.addNode(node)
}

func (l *LRUCache) removeNode(node *LruLinkNode) int {
	if node == l.end {
		l.end = l.end.pre
	} else if node == l.head {
		l.head = l.head.next
	} else {
		node.pre.next = node.next
		node.next.pre = node.pre
	}
	return node.Key
}

func (l *LRUCache) addNode(node *LruLinkNode) {
	if l.end != nil {
		l.end.next = node
		node.pre = l.end
		node.next = nil
	}
	l.end = node
	if l.head == nil {
		l.head = node
	}
}

func main() {
	LRUCache := Constructor(3)
	LRUCache.Put(1, 1)
	LRUCache.Put(2, 2)
	LRUCache.Put(3, 3)
	LRUCache.Put(4, 4)
	LRUCache.Put(5, 5)
	fmt.Println(LRUCache.Get(3))
}
