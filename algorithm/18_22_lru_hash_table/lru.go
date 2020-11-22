/**
 * Created by lock
 * lru缓存淘汰，hash表，双向链表
    使用双向链表存储数据，链表中的每个结点处理存储数据（data）、前驱指针（prev）、后继指针（next）之外，
	还新增了一个特殊的字段 hnext，hnext 指针是为了将结点串在散列表的拉链中（解决散列冲突），前驱和后继指针是为了将结点串在双向链表中。
	查找：
		O(1)定位，当找到数据之后，我们还需要将它移动到双向链表的尾部。
	删除：
		O(1)时间复杂度里找到要删除的结点。因为我们的链表是双向链表，双向链表可以通过前驱指针O(1)时间复杂度获取前驱结点，所以在双向链表中，删除结点只需要O(1)的时间复杂度。
	添加：
		先看这个数据是否已经在缓存中。如果已经在其中，需要将其移动到双向链表的尾部。
		如果不在其中，还要看缓存有没有满。如果满了，则将双向链表头部的结点删除，然后再将数据放到链表的尾部。
		如果没有满，就直接将数据放到链表的尾部。
*/
package main

import "fmt"

type lruNode struct {
	key      int
	value    interface{}
	hashNext *lruNode //发生hash碰撞，链表法
	prev     *lruNode
	next     *lruNode
}

type lruCache struct {
	node     []lruNode
	head     *lruNode
	tail     *lruNode
	capacity int //容量
	used     int //已经使用
}

func newLruCache(capacity int) *lruCache {
	return &lruCache{
		node:     make([]lruNode, capacity),
		head:     nil, //头指针
		tail:     nil, //尾指针
		capacity: capacity,
		used:     0,
	}
}

func (l *lruCache) getHashKey(key int) int {
	is64Bit := uint64(^uintptr(0)) == ^uint64(0)
	//判断是32bit位 or 64bit位的系统
	if is64Bit {
		//这个hash计算方法与java hashMap的计算hashKey的方法类似
		return (key ^ (key >> 32)) & (l.capacity - 1)
	}
	return (key ^ (key >> 16)) & (l.capacity - 1)
}

func (l *lruCache) addNode(key, val int) {
	newNode := &lruNode{
		key:   key,
		value: val,
	}
	hashKey := l.getHashKey(key)
	hashBucket := &l.node[hashKey]
	//-- 相当于放到链表头部
	newNode.hashNext = hashBucket.hashNext
	hashBucket.hashNext = newNode
	//--
	l.used++
	//如果是空的lru
	if l.tail == nil {
		l.tail, l.head = newNode, newNode
		return
	}
	//链表尾部新增节点
	l.tail.next = newNode
	newNode.prev = l.tail
	//更新 tail 指针指向位置
	l.tail = newNode
}

//移除节点
func (l *lruCache) delNode() {
	if l.head == nil {
		return
	}
	delNodeHashKey := l.getHashKey(l.head.key)
	prev := &l.node[delNodeHashKey]
	tmp := prev.hashNext
	for tmp != nil && tmp.key != l.head.key {
		prev = tmp
		tmp = tmp.hashNext
	}
	if tmp == nil {
		return
	}
	//移除
	prev.hashNext = tmp.hashNext
	l.head = l.head.next
	l.head.prev = nil
	l.used--
}

//移到双向链表尾部
func (l *lruCache) moveToTail(searchNode *lruNode) {
	if l.tail == searchNode {
		return
	}
	if l.head == searchNode {
		l.head = searchNode.next
		l.head.prev = nil
	} else {
		searchNode.next.prev = searchNode.prev
		searchNode.prev.next = searchNode.next
	}
	searchNode.next = nil
	//tail pointer
	l.tail.next = searchNode
	searchNode.prev = l.tail

	l.tail = searchNode
}

func (l *lruCache) searchNode(key int) (searchNode *lruNode) {
	if l.tail == nil {
		return nil
	}
	//根据key算出hash所在的位置
	hashKey := l.getHashKey(key)
	linkedList := l.node[hashKey].hashNext
	for linkedList != nil {
		if linkedList.key == key {
			return linkedList
		}
		linkedList = linkedList.hashNext
	}
	return nil
}

func (l *lruCache) Get(key int) interface{} {
	if l.tail == nil {
		return -1
	}
	if searchNode := l.searchNode(key); searchNode != nil {
		l.moveToTail(searchNode)
		return searchNode.value
	}
	return -1
}

func (l *lruCache) Put(key, val int) {
	//是有已经在lru中
	if searchNode := l.searchNode(key); searchNode != nil {
		searchNode.value = val
		l.moveToTail(searchNode)
		return
	}
	l.addNode(key, val)
	if l.used > l.capacity {
		//超出容量，移除head
		l.delNode()
	}
}

func main() {
	lru := newLruCache(5)
	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Put(3, 3)
	lru.Put(4, 4)
	fmt.Println(lru.Get(2))
	lru.Put(2, 222)
	fmt.Println(lru.Get(2))
	lru.Put(5, 5)
	lru.Put(6, 6)
	fmt.Println(lru.Get(1))
}
