/**
 * Created by lock
 * 递归反转链表
 */
package main

import "fmt"

type linkNodeItem struct {
	val  int
	next *linkNodeItem
}

type linkTable struct {
	Head *linkNodeItem
}

func newLinkRecursion() *linkTable {
	link := &linkTable{
		Head: &linkNodeItem{
			val: 1,
			next: &linkNodeItem{
				val: 2,
				next: &linkNodeItem{
					val: 3,
					next: &linkNodeItem{
						val: 4,
						next: &linkNodeItem{
							val: 5,
							next: &linkNodeItem{
								val: 6,
								next: &linkNodeItem{
									val:  7,
									next: nil,
								},
							},
						},
					},
				},
			},
		},
	}
	return link
}

func (l *linkTable) print() {
	if l == nil {
		return
	}
	node := l.Head
	str := ""
	for node != nil {
		str += fmt.Sprintf("%d->", node.val)
		node = node.next
	}
	fmt.Println(str)
}

func revLink(link *linkTable) (newLink *linkTable) {
	node := link.Head
	var pre *linkNodeItem
	for node != nil {
		next := node.next
		node.next = pre
		pre = node
		node = next
	}
	link.Head = pre
	return link
}

func recursionRevLink(node *linkNodeItem) (newNode *linkNodeItem) {
	if node.next == nil {
		return node
	}
	last := recursionRevLink(node.next)
	node.next.next = node
	node.next = nil
	return last
}

func (l *linkTable) recursionLink(link *linkTable) (newLink *linkTable) {
	head := link.Head
	revLink := recursionRevLink(head)
	newLink = &linkTable{
		Head: revLink,
	}
	return newLink
}

var successor = &linkNodeItem{}

//反转链表前k个元素
func reverseTopKSub(node *linkNodeItem, k int) (topK *linkNodeItem) {
	if k == 1 {
		// 记录第 n + 1 个节点
		successor = node.next
		return node
	}
	last := reverseTopKSub(node.next, k-1)
	node.next.next = node
	node.next = successor
	return last
}

func (l *linkTable) reverseTopK(link *linkTable, k int) (newLink *linkTable) {
	head := link.Head
	newRevLink := reverseTopKSub(head, k)
	newLink = &linkTable{
		Head: newRevLink,
	}
	return newLink
}

func reverseBetweenSub(node *linkNodeItem, m, n int) (betweenNode *linkNodeItem) {
	if m == 1 {
		return reverseTopKSub(node, n)
	}
	node.next = reverseBetweenSub(node.next, m-1, n-1)
	return node
}

//反转部分链表
func (l *linkTable) reverseBetween(link *linkTable, m, n int) (newLink *linkTable) {
	head := link.Head
	newRevBtLink := reverseBetweenSub(head, m, n)
	newLink = &linkTable{
		Head: newRevBtLink,
	}
	return newLink
}

func main() {
	link := newLinkRecursion()
	fmt.Print("打印原始链表：")
	link.print()
	//第一次反转
	revLink := revLink(link)
	fmt.Print("打印反转后的链表：")
	revLink.print()
	//第二次递归反转
	revCsLink := link.recursionLink(revLink)
	fmt.Print("在反转后的链表基础上，打印递归反转后的链表：")
	revCsLink.print()
	fmt.Print("反转前k个链表：")
	newRevCsLink := link.reverseTopK(revCsLink, 3)
	newRevCsLink.print()
	fmt.Print("反转部分(m=3,n=5)链表：")
	link.reverseBetween(newRevCsLink, 3, 5).print()
}
