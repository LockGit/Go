/**
 * Created by lock
回溯实现 + 备忘录 （记录重复的递归项）
假设我们有几种不同币值的硬币 v1，v2，……，vn（单位是元）。如果我们要支付 w 元，求最少需要多少个硬币。
比如，我们有 3 种不同的硬币，1 元、3 元、5 元，我们要支付 9 元，最少需要 3 个硬币（3 个 3 元的硬币）
*/
package main

import "fmt"

var coinMap = map[int]int{}

func minVal(nums ...int) int {
	m := nums[0]
	for _, v := range nums {
		if v < m {
			m = v
		}
	}
	return m
}

func minCoin(money int) int {
	//备忘录，记录重复的递归项
	if cache, ok := coinMap[money]; ok {
		return cache
	}
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
	coin := 1 + minVal(minCoin(money-1), minCoin(money-3), minCoin(money-5))
	coinMap[money] = coin
	return coin
}

func main() {
	fmt.Println(minCoin(9))
}
