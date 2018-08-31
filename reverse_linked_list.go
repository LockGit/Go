/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/30
 * Time: 17:56
 * golang反转链表
 */
package main

import "fmt"

type reverseLink struct {
	data int
	next *reverseLink
}

func reverseLinked(node *reverseLink) *reverseLink {
	newNode := new(reverseLink)
	for node.next != nil {
		next := node.next
		node.next = newNode
		newNode = node
		node = next
	}
	return newNode
}

func main() {
	node1 := new(reverseLink)
	node2 := new(reverseLink)
	node3 := new(reverseLink)
	node4 := new(reverseLink)
	node1.data = 1
	node1.next = node2
	node2.data = 2
	node2.next = node3
	node3.data = 3
	node3.next = node4
	node4.data = 4
	node4.next = new(reverseLink)

	node := reverseLinked(node1)
	for node.next != nil {
		fmt.Println(node.data)
		node = node.next
	}
}
