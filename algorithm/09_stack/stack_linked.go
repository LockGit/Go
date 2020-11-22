/**
 * Created by lock
 * 链式栈，用链表实现的栈
 */
package main

import "fmt"

type node struct {
	next *node
	val  interface{}
}

type linkStack struct {
	topNode *node
}

func newLinkStack() *linkStack {
	return &linkStack{nil}
}

func (l *linkStack) isEmpty() bool {
	return l.topNode == nil
}

func (l *linkStack) pop() interface{} {
	if l.isEmpty() {
		return nil
	}
	v := l.topNode.val
	l.topNode = l.topNode.next
	return v
}

func (l *linkStack) push(v interface{}) {
	l.topNode = &node{
		next: l.topNode,
		val:  v,
	}
}

func main() {
	sl := newLinkStack()
	for i := 0; i <= 3; i++ {
		sl.push(i)
		fmt.Println("push", i)
	}
	for i := 0; i <= 3; i++ {
		sl.pop()
		fmt.Println("pop", i)
	}
}
