/**
 * Created by lock
 * 基于数组实现的顺序队列
 */
package main

import "fmt"

type arrayQueue struct {
	Items []interface{}
	Size  int
	Head  int
	Tail  int
}

func newArrayQueue(size int) *arrayQueue {
	return &arrayQueue{
		Items: make([]interface{}, size),
		Size:  size,
		Head:  0,
		Tail:  0,
	}
}

func (a *arrayQueue) push(item interface{}) bool {
	if a.Tail == a.Size {
		return false
	}
	a.Items[a.Tail] = item
	a.Tail++
	return true
}

func (a *arrayQueue) pop() interface{} {
	if a.Tail == a.Head {
		return nil
	}
	v := a.Items[a.Head]
	a.Head++
	return v
}

func (a *arrayQueue) String() string {
	if a.Head == a.Tail {
		return "empty queue"
	}
	result := "head"
	for i := a.Head; i <= a.Tail-1; i++ {
		result += fmt.Sprintf("<-%+v", a.Items[i])
	}
	result += "<-tail"
	return result
}

func main() {
	q := newArrayQueue(10)
	for i := 0; i < 11; i++ {
		if b := q.push(i); b == false {
			fmt.Println("push fail:", i)
		}
	}
	fmt.Println(q.String())

	for i := 0; i <= 11; i++ {
		if v := q.pop(); v != nil {
			fmt.Println("pop val:", v)
		}
	}

	fmt.Println(q.String())
}
