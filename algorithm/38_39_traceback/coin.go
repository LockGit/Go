/**
 * Created by lock
回溯实现
假设我们有几种不同币值的硬币 v1，v2，……，vn（单位是元）。如果我们要支付 w 元，求最少需要多少个硬币。
比如，我们有 3 种不同的硬币，1 元、3 元、5 元，我们要支付 9 元，最少需要 3 个硬币（3 个 3 元的硬币）
*/
package main

import "fmt"

func min(nums ...int) int {
	m := nums[0]
	for _, v := range nums {
		if v < m {
			m = v
		}
	}
	return m
}

func calcMinCoinNum(money int) (minCoinNum int) {
	if money == 1 {
		return 1
	}
	if money == 2 {
		return 2
	}
	if money == 3 {
		return 1
	}
	if money == 4 {
		return 2
	}
	if money == 5 {
		return 1
	}
	return 1 + min(calcMinCoinNum(money-1), calcMinCoinNum(money-3), calcMinCoinNum(money-5))
}

func main() {
	fmt.Println(calcMinCoinNum(9))
}
