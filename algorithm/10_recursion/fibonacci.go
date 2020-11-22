/**
 * Created by lock
 * 斐波那契数列
 */
package main

import "fmt"

type fib struct {
	val map[int]int
}

func newFib(n int) *fib {
	return &fib{
		val: make(map[int]int, n),
	}
}

func (f *fib) getFib(n int) int {
	if f.val[n] != 0 {
		return f.val[n]
	}
	if n == 1 || n == 2 {
		f.val[n] = 1
		return 1
	}
	res := f.getFib(n-1) + f.getFib(n-2)
	f.val[n] = res
	return res
}

func (fib *fib) Print(n int) {
	fmt.Println(fib.val[n])
}

func main() {
	fib := newFib(10)
	for i := 1; i < 15; i++ {
		fib.getFib(i)
		fib.Print(i)
	}
}
