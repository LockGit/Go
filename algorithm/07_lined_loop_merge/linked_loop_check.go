/**
 * Created by lock
 * 链表中环的检测
 */
package main

import "fmt"

type nodeLoopItem struct {
	value interface{}
	next  *nodeLoopItem
}

type linkedLoop struct {
	header *nodeLoopItem
}

func (l *linkedLoop) add(n int) {
	l.header = &nodeLoopItem{
		value: "header",
		next:  nil,
	}
	headerNode := l.header
	for i := 0; i < n; i++ {
		node := &nodeLoopItem{
			value: i,
			next:  nil,
		}
		headerNode.next = node
		headerNode = headerNode.next
	}
}

func (l *linkedLoop) checkHasCycle() bool {
	if l.header != nil {
		fast := l.header
		slow := l.header
		for fast != nil && fast.next != nil {
			slow = slow.next
			fast = fast.next.next
			if fast == slow {
				return true
			}
		}
	}
	return false
}

func main() {
	llc := new(linkedLoop)
	llc.add(8)
	node := llc.header
	for node != nil {
		fmt.Println(node.value)
		node = node.next
	}

}
