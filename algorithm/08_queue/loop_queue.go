/**
 * Created by lock
 * 基于数组的环形队列
 */
package main

import "fmt"

type loopArrQueue struct {
	data []interface{}
	head int
	tail int
	size int
}

func newLoopArrQueue(size int) *loopArrQueue {
	return &loopArrQueue{
		data: make([]interface{}, size),
		head: 0,
		tail: 0,
		size: size,
	}
}

func (l *loopArrQueue) isEmpty() bool {
	if l.head == l.tail {
		return true
	}
	return false
}

func (l *loopArrQueue) isFull() bool {
	if (l.tail+1)%l.size == l.head {
		return true
	}
	return false
}

func (l *loopArrQueue) push(item interface{}) bool {
	if l.isFull() {
		return false
	}
	l.data[l.tail] = item
	l.tail = (l.tail + 1) % l.size
	return true
}

func (l *loopArrQueue) pop() interface{} {
	if l.isEmpty() {
		return nil
	}
	v := l.data[l.head]
	l.head = (l.head + 1) % l.size
	return v
}

func (l *loopArrQueue) String() string {

	if l.isEmpty() {
		return "empty queue"
	}
	result := "head"
	var i = l.head
	for true {
		result += fmt.Sprintf("<-%+v", l.data[i])
		i = (i + 1) % l.size
		if i == l.tail {
			break
		}
	}
	result += "<-tail"
	return result
}

func main() {
	q := newLoopArrQueue(9)
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
