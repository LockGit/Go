/**
 * Created by lock
 * 查找链表的中间节点，删除链表的第n个节点操作，删除链表的倒数第n个节点操作
 */
package main

import "fmt"

type node struct {
	value interface{}
	next  *node
}

type link struct {
	header *node
}

func (l *link) add(n int) {
	l.header = &node{
		value: "head",
		next:  nil,
	}
	lnk := l.header
	for i := 0; i < n; i++ {
		node := &node{
			value: i,
			next:  nil,
		}
		lnk.next = node
		lnk = lnk.next
	}
}

func (l *link) findCenterNode() *node {
	if l.header == nil || l.header.next == nil {
		return nil
	}
	if l.header.next.next == nil {
		return l.header.next
	}
	//找到中间节点
	fast, slow := l.header, l.header
	for nil != fast && nil != fast.next {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

//删除第n个节点 可以使用快慢指针
func (l *link) delNodeByNum(num int) *link {
	if num <= 0 {
		num = 1
	}
	delNodeLink := l.header
	i := 1
	for delNodeLink != nil {
		if i == num {
			delNodeLink.next = delNodeLink.next.next
			break
		}
		delNodeLink = delNodeLink.next
		i++
	}
	return l
}

//删除倒数第n个节点
func (l *link) delNodeByRevNum(num int) {
	if nil == l.header || nil == l.header.next {
		return
	}
	if num <= 0 {
		num = 1
	}
	fast := l.header
	for i := 1; i <= num && fast != nil; i++ {
		fast = fast.next
	}
	if fast == nil {
		return
	}
	slow := l.header
	//这个地方用fast.next 针对1个节点的情况
	for fast.next != nil {
		slow = slow.next
		fast = fast.next
	}
	slow.next = slow.next.next
	return
}

func main() {
	l := new(link)
	//增加6个节点
	l.add(6)
	node := l.header
	for node != nil {
		fmt.Println(node.value)
		node = node.next
	}
	centerNode := l.findCenterNode()
	fmt.Println("center node", centerNode)

	fmt.Println("del node...")
	delNodeLink := l.delNodeByNum(2).header
	for delNodeLink != nil {
		fmt.Println(delNodeLink.value)
		delNodeLink = delNodeLink.next
	}

	fmt.Println("dev rev node...")
	l.delNodeByRevNum(1)
	delRevNodeLink := l.header
	for delRevNodeLink != nil {
		fmt.Println(delRevNodeLink.value)
		delRevNodeLink = delRevNodeLink.next
	}
}
