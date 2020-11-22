/**
 * Created by lock
 * 阶乘的递归实现
 */
package main

import (
	"fmt"
)

type factorial struct {
	val map[int]int
}

func newFac(n int) *factorial {
	return &factorial{
		val: make(map[int]int, n),
	}
}

func (f *factorial) facMul(n int) int {
	if f.val[n] != 0 {
		return f.val[n]
	}
	if n <= 1 {
		f.val[n] = 1
		return 1
	} else {
		res := n * f.facMul(n-1)
		f.val[n] = res
		return res
	}
}

func (f *factorial) print(n int) {
	println(f.val[n])
}

//递归
func (f *factorial) Mul(n int) int {
	if n <= 1 {
		return 1
	}
	return n * f.Mul(n-1)
}

//非递归
func (f *factorial) Mul2(n int) int {
	res := 1
	for n > 0 {
		res = res * n
		n--
	}
	return res
}

func main() {
	f := new(factorial)
	res := f.Mul(3)
	fmt.Println("res", res)
	res2 := f.Mul2(3)
	fmt.Println("res2", res2)

	fac := newFac(10)
	for i := 1; i < 15; i++ {
		//对多个数进行阶乘，这个地方循环的时候，对于已经计算过的不用重新计算
		fac.facMul(i)
		fac.print(i)
	}
}
