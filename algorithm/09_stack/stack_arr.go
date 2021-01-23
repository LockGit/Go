/**
 * Created by lock
 * 顺序栈，用数组实现的栈。
 */
package main

import (
	"errors"
	"fmt"
)

type stackArr struct {
	Items []interface{} //数组
	Size  int           //栈的大小
	Count int           //栈中元素个数
}

//压栈
func (s *stackArr) push(item interface{}) bool {
	if s.Size == s.Count {
		return false
	}
	s.Items = append(s.Items, item)
	s.Count++
	return true
}

//出栈
func (s *stackArr) pop() (item interface{}, err error) {
	if s.Count == 0 {
		return nil, errors.New("no item")
	}
	item = s.Items[s.Count-1]
	s.Count = s.Count - 1
	return
}

//tips:demo,非线程安全
func main() {
	stackSize := 100
	s := &stackArr{
		Items: make([]interface{}, stackSize),
		Size:  stackSize,
		Count: 0,
	}
	for i := 0; i < 8; i++ {
		ret := s.push(i)
		fmt.Println("push:", i, ",result:", ret)
	}
	for i := 0; i < 8; i++ {
		if ret, err := s.pop(); err == nil {
			fmt.Println("pop:", i, ",result:", ret)
		}
	}

}
