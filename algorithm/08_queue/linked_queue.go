/**
 * Created by lock
 * 基于链表的链式队列
 */
package main

import "fmt"

type queueNode struct {
	val  interface{}
	next *queueNode
}

type linkedQueue struct {
	head       *queueNode
	tail       *queueNode
	realLength int
	size       int
}

func newLinkedQueue(size int) *linkedQueue {
	return &linkedQueue{
		head:       nil,
		tail:       nil,
		realLength: 0,
		size:       size,
	}
}

func (l *linkedQueue) push(val interface{}) {
	if l.realLength == l.size {
		return
	}
	node := &queueNode{
		val:  val,
		next: nil,
	}
	if l.tail == nil {
		l.tail = node
		l.head = node
	} else {
		l.tail.next = node
		l.tail = node
	}
	l.realLength++
}

func (l *linkedQueue) pop() interface{} {
	if l.head == nil {
		return nil
	}
	val := l.head.val
	l.head = l.head.next
	l.realLength--
	return val
}

func (l *linkedQueue) String() string {
	if l.head == nil {
		return "empty queue"
	}
	result := "head<-"
	for cur := l.head; cur != nil; cur = cur.next {
		result += fmt.Sprintf("<-%+v", cur.val)
	}
	result += "<-tail"
	return result
}

func main() {
	q := newLinkedQueue(2)
	for i := 0; i <= 6; i++ {
		q.push(i)
	}
	fmt.Println(q)

	for i := 0; i <= 3; i++ {
		if v := q.pop(); v != nil {
			fmt.Println("pop:", v)
		}
	}
	fmt.Println(q)
}
