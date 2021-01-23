/**
 * Created by lock
 * 判断链表是否有环
 * 已知链表有环，找到环的开始位置
 */
package main

import (
	"fmt"
	"strings"
)

type linkLoopNode struct {
	val  interface{}
	next *linkLoopNode
}

type linkLoop struct {
	Head *linkLoopNode
}

func newLinkLoop() *linkLoop {
	return &linkLoop{
		Head: &linkLoopNode{},
	}
}

func (l *linkLoop) add(n int) {
	l.Head = &linkLoopNode{
		val:  "head",
		next: nil,
	}
	node := l.Head
	for i := 0; i <= n; i++ {
		nodeItem := &linkLoopNode{
			val:  i,
			next: nil,
		}
		node.next = nodeItem
		node = node.next
	}
}

func (l *linkLoop) Print() {
	node := l.Head
	var arr []string
	for node != nil {
		arr = append(arr, fmt.Sprintf("%v", node.val))
		node = node.next
	}
	fmt.Println(strings.Join(arr, "-->"))
}

//添加一个环
func (l *linkLoop) addRingNode() {
	node := l.Head
	i := 0
	var startNode *linkLoopNode
	for node != nil {
		if i == 2 {
			startNode = node
		}
		if i == 6 {
			//add a ring , break
			node.next = startNode
			break
		}
		node = node.next
		i++
	}
}

//判断链表是否有环
func (l *linkLoop) checkHaveLinkLoop() (b bool) {
	link := l.Head
	slow, fast := link, link
	for slow != nil && fast != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			return true
		}
	}
	return false
}

//已知链表有环，找到环开始的节点
func (l *linkLoop) findLinkLoopPosition() (node *linkLoopNode) {
	link := l.Head
	slow, fast := link, link
	for slow != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			//第一次相遇后，break
			break
		}
	}
	//慢指针重新指向head,一次步进一个
	slow = link
	for slow != fast && fast != nil {
		slow = slow.next
		fast = fast.next
	}
	return slow
}

func main() {
	l := newLinkLoop()
	l.add(8)
	l.Print()
	l.addRingNode()
	fmt.Println("check loop ret:", l.checkHaveLinkLoop())
	fmt.Println("find loop position:", l.findLinkLoopPosition())
}
