/**
 * Created by lock
Fibonacci 严格来说不属于动态规划，不涉及求最值
可以用动态规划的思路来优化斐波那契数列
*/
package main

import "fmt"

//没有优化的fib
func fib(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

var fibMemo = map[int]int{}

//增加备忘录的fib
func fibAddMemo(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	if memoVal, ok := fibMemo[n]; ok {
		return memoVal
	}
	memo := fibAddMemo(n-1) + fibAddMemo(n-2)
	fibMemo[n] = memo
	return fibMemo[n]
}

func main() {
	fmt.Println("raw fib ...")
	for i := 0; i <= 4; i++ {
		fmt.Println(fib(i))
	}

	fmt.Println("add memo fib ...")
	for i := 0; i <= 4; i++ {
		fmt.Println(fibAddMemo(i))
	}

	fmt.Println("add equation fib ...(使用状态转移方程的fib（自底向上）)")
	n := 5
	fibEquation := make([]int, n+1)
	fibEquation[1] = 1
	for i := 2; i <= n; i++ {
		fibEquation[i] = fibEquation[i-1] + fibEquation[i-2]
	}
	fmt.Println(fibEquation)

	fmt.Println("只涉及两个状态的变更，状态转移方程的空间复杂度可以由O(n)转到O(1)，dp table的空间复杂度压缩")
	prev := 0
	cur := 1
	for i := 2; i <= n; i++ {
		sum := prev + cur
		fmt.Println(sum)
		prev = cur
		cur = sum
	}
}
