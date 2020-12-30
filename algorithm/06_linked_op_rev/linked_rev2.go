/**
 * Created by lock
给定一个单链表的头节点 head,实现一个调整单链表的函数，使得每K个节点之间为一组进行逆序，并且从链表的尾部开始组起，
头部剩余节点数量不够一组的不需要逆序。（不能使用队列或者栈作为辅助）

链表:1->2->3->4->5->6->7->8->null, K = 3。
那么 6->7->8，3->4->5，1->2各位一组。
调整后：1->2->5->4->3->8->7->6->null。
其中 1，2不调整，因为不够一组。
*/
package main

import "fmt"

type linkNode struct {
	value        int
	next         *linkNode
	isSplitPoint bool
}

type linkList struct {
	head *linkNode
}

func newLink() *linkList {
	return &linkList{
		head: &linkNode{
			value: 1,
			next:  nil,
		},
	}
}

func (l *linkList) addNode(item linkNode) {
	node := l.head
	for node != nil {
		pre := node
		node = node.next
		if node == nil {
			pre.next = &item
		}
	}
}

func (l *linkList) reverse() {
	node := l.head
	var pre *linkNode
	for node != nil {
		next := node.next
		node.next = pre
		pre = node
		node = next
	}
	l.head = pre
	return
}

func (l *linkList) setRevLinkSplitPoint(k int) {
	header := l.head
	link := l.head
	i := 0
	for link != nil {
		i++
		if i == 0 {
			link.isSplitPoint = false
		}
		if i == k {
			link.isSplitPoint = true
			i = 0
		}
		pre := link
		link = link.next
		if link == nil {
			pre.isSplitPoint = false
		}
	}
	l.head = header
	return
}

func (l *linkList) connectPrev(tmpNode *linkNode, pre *linkNode) *linkNode {
	header := tmpNode
	for tmpNode.next != nil {
		tmpNode = tmpNode.next
	}
	tmpNode.next = pre
	return header
}

func (l *linkList) change(k int) {
	link := l.head
	count := 1
	var pre *linkNode
	start := link
	//var tmpLink linkList
	for link != nil {
		//fmt.Println("x--->", link.value, link.isSplitPoint)
		if link.isSplitPoint && count == k {
			end := link
			fmt.Println("gap", start.value, end.value)
			tmpLink := &linkList{
				head: &linkNode{
					value: start.value,
				},
			}
			//fmt.Println(fmt.Sprintf("addr %p", tmpLink))
			tmpNode := tmpLink.head
			for start != end {
				item := &linkNode{
					value: start.next.value,
					next:  nil,
				}
				tmpNode.next = item
				tmpNode = tmpNode.next
				start = start.next
			}
			tmpNode.next = nil
			pre = l.connectPrev(tmpLink.head, pre)
			start = link.next
			count = 0
		}
		link = link.next
		count++
	}
	var revNewLink *linkNode
	for start != nil {
		next := start.next
		start.next = revNewLink
		revNewLink = start
		start = next
	}
	p := revNewLink
	for revNewLink != nil {
		revNewLink = revNewLink.next
		if revNewLink.next == nil {
			revNewLink.next = pre
			break
		}
	}
	l.head = p
	return
}

func (l *linkList) echo() {
	node := l.head
	str := ""
	for node != nil {
		str += fmt.Sprintf("%d->", node.value)
		node = node.next
	}
	fmt.Println(str)
}

func main() {
	link := newLink()
	for i := 2; i <= 8; i++ {
		item := linkNode{
			value: i,
			next:  nil,
		}
		link.addNode(item)
	}
	link.echo()
	link.reverse()
	link.echo()
	k := 3
	link.setRevLinkSplitPoint(k)
	link.change(k)
	link.echo()
}
