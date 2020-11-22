/**
 * Created by lock
 * 模拟数组的插入、删除、按照下标随机访问操作
 */
package main

import (
	"fmt"
	"github.com/pkg/errors"
)

type arr struct {
	data   []int
	length int
}

func newArr(size int) *arr {
	if size <= 0 {
		return nil
	}
	return &arr{
		data:   make([]int, size, size),
		length: 0,
	}
}

func (a *arr) len() int {
	return a.length
}

func (a *arr) checkOutRange(index int) bool {
	if index >= cap(a.data) {
		return true
	}
	return false
}

//插入数值到索引index上
func (a *arr) insert(index int, val int) error {
	if a.len() == cap(a.data) {
		return errors.New("arr full")
	}
	//越界check
	if (index != a.length && a.checkOutRange(index)) || index < 0 {
		return errors.New("out of index range")
	}
	for i := a.length; i > index; i-- {
		//元素都后移了
		a.data[i] = a.data[i-1]
	}
	a.data[index] = val
	a.length++
	return nil
}

func (a *arr) insertToTail(v int) error {
	return a.insert(a.len(), v)
}

//查找
func (a *arr) find(index int) (int, error) {
	if a.checkOutRange(index) {
		return 0, errors.New("find out of range")
	}
	return a.data[index], nil
}

//删除
func (a *arr) delete(index int) (v int, err error) {
	if a.checkOutRange(index) {
		return 0, errors.New("delete out of range")
	}
	v = a.data[index]

	//1 2 3 4 5 6 0 0 0 0 0
	for i := a.len() - 1; i > index; i-- {
		a.data[i-1] = a.data[i]
	}
	//for i := index; i < a.len()-1; i++ {
	//	a.data[i] = a.data[i+1]
	//}
	if index < a.length {
		a.length--
	}
	return v, nil
}

func (a *arr) print() {
	for i := 0; i < a.len(); i++ {
		fmt.Println("key:", i, ",val:", a.data[i])
	}
}

func main() {
	a := newArr(10)
	a.insert(0, 8)
	a.insert(1, 8)
	a.insert(2, 8)
	a.insert(1, 18)
	fmt.Println("find index eq 1:")
	fmt.Println(a.find(1))
	fmt.Println("delete index eq 1:")
	//fmt.Println(a.delete(1))
	//fmt.Println(a.delete(0))
	fmt.Println("---")
	a.print()
	fmt.Println("now arr length", a.length)
	fmt.Println(a.delete(8))
	fmt.Println(a.delete(9))
	fmt.Println(a.delete(1))
	//fmt.Println(a.delete(3))
	fmt.Println("end")
	a.print()
}
