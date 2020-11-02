/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/21
 * Time: 00:03
 * 将获得两个非空链接列表，表示两个非负整数。 数字以相反的顺序存储，并且它们的每个节点包含单个数字。 添加两个数字并将其作为链接列表返回。
 * https://www.jianshu.com/p/1a4c349ae859
 * Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
 * Output: 7 -> 0 -> 8
 */

package main

import (
	"fmt"
)

type ElementType interface{}

type Node struct {
	Data ElementType
	next *Node
}

type LinkedList struct {
	head *Node
}

func CreateLinkedList() *LinkedList {
	head := new(Node)
	return &LinkedList{head}
}

func getLinkedListOne() *LinkedList {
	node1 := CreateLinkedList()
	node2 := CreateLinkedList()
	node3 := CreateLinkedList()
	node4 := CreateLinkedList()
	node4.head.Data = nil
	node4.head.next = nil
	node1.head.Data = 2
	node1.head.next = node2.head
	node2.head.Data = 4
	node2.head.next = node3.head
	node3.head.Data = 3
	node3.head.next = node4.head
	return node1
}

func getLinkedListTwo() *LinkedList {
	node1 := CreateLinkedList()
	node2 := CreateLinkedList()
	node3 := CreateLinkedList()
	node4 := CreateLinkedList()
	node4.head.Data = nil
	node4.head.next = nil
	node1.head.Data = 5
	node1.head.next = node2.head
	node2.head.Data = 6
	node2.head.next = node3.head
	node3.head.Data = 4
	node3.head.next = node4.head
	return node1
}

func CreateNewLinkedList(num int) *LinkedList {
	node := CreateLinkedList()
	node.head.Data = num
	return node
}

func main() {
	linkedListOne := getLinkedListOne().head
	linkedListTwo := getLinkedListTwo().head
	carry := 0
	var res LinkedList
	var resEnd LinkedList
	for linkedListOne.next != nil || linkedListTwo.next != nil || carry == 1 {
		tmp := 0
		if linkedListOne.next != nil && linkedListOne.Data != nil {
			tmp = tmp + linkedListOne.Data.(int)
			linkedListOne = linkedListOne.next
		}

		if linkedListTwo.next != nil && linkedListTwo.Data != nil {
			tmp = tmp + linkedListTwo.Data.(int)
			linkedListTwo = linkedListTwo.next
		}
		tmp = tmp + carry
		digit := tmp % 10
		carry = tmp / 10
		if res.head == nil {
			res.head = CreateNewLinkedList(digit).head
			resEnd.head = res.head
		} else {
			resEnd.head.next = CreateNewLinkedList(digit).head
			resEnd.head = resEnd.head.next
		}
	}
	//创建一个空链表链接到尾部，后面遍历链表for循环有!=nil判断
	endNode := CreateLinkedList().head
	resEnd.head.next = endNode

	fmt.Println("打印链表数据:")
	resNode := res.head
	for resNode.next != nil {
		fmt.Println(resNode.Data)
		resNode = resNode.next
	}
}
