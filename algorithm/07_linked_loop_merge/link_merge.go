/**
 * Created by lock
 * 两个有序的链表合并
	//1->5->10
	//2->3->4->8->10
	//1->2->3->4->5->8->10
*/
package main

import "fmt"

type nodeItem struct {
	value interface{}
	next  *nodeItem
}

type linkMerge struct {
	header *nodeItem
}

func (l linkMerge) add(n int) *linkMerge {
	l.header = &nodeItem{
		value: "header",
		next:  nil,
	}
	headerNode := l.header
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			//留点缝隙
			continue
		}
		item := &nodeItem{
			value: i,
			next:  nil,
		}
		headerNode.next = item
		headerNode = headerNode.next
	}
	return &l
}

func (l linkMerge) add2(n int) *linkMerge {
	l.header = &nodeItem{
		value: "header",
		next:  nil,
	}
	headerNode := l.header
	for i := 0; i < n; i++ {
		if i%2 != 0 {
			continue
		}
		item := &nodeItem{
			value: i,
			next:  nil,
		}
		headerNode.next = item
		headerNode = headerNode.next
	}
	return &l
}

func (l *linkMerge) mergeLink(l1, l2 *linkMerge) *linkMerge {
	if nil == l1 || nil == l1.header || nil == l1.header.next {
		return l2
	}
	if nil == l2 || nil == l2.header || nil == l2.header.next {
		return l1
	}
	newLink := &linkMerge{
		header: &nodeItem{
			value: "newHeader",
			next:  nil,
		},
	}
	newLinkCurrentItem := newLink.header
	link1 := l1.header.next
	link2 := l2.header.next
	for link1 != nil && link2 != nil {
		if link1.value.(int) > link2.value.(int) {
			newLinkCurrentItem.next = link2
			link2 = link2.next
		} else {
			newLinkCurrentItem.next = link1
			link1 = link1.next
		}
		newLinkCurrentItem = newLinkCurrentItem.next
	}
	//剩余部分
	if link1 != nil {
		newLinkCurrentItem.next = link1
	} else if link2 != nil {
		newLinkCurrentItem.next = link2
	}
	return newLink
}

func (l *linkMerge) merge2(link1, link2 *nodeItem) *linkMerge {

	return l
}

func main() {
	lm := new(linkMerge)
	link1 := lm.add(8)
	node := link1.header
	fmt.Println("first link")
	for node != nil {
		fmt.Println(node.value)
		node = node.next
	}
	//第二条链表
	link2 := lm.add2(8)
	node2 := link2.header
	fmt.Println("second link")
	for node2 != nil {
		fmt.Println(node2.value)
		node2 = node2.next
	}

	cLink1 := link1 // header忽略
	cLink2 := link2 // header忽略
	fmt.Println("--merge--")
	newLink := lm.mergeLink(cLink1, cLink2)
	mergeLink := newLink.header
	for mergeLink != nil {
		fmt.Println(mergeLink.value)
		mergeLink = mergeLink.next
	}
}
