/**
 * Created by lock
动态规划，状态转移方程
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

var status = map[int]int{
	1: 1, //支付1元时，至少需要1枚硬币
	2: 2, //支付2元时，至少需要2枚硬币
	3: 1, //支付3元时，至少需要1枚硬币
	4: 2, //支付4元时，至少需要2枚硬币
	5: 1, //支付5元时，至少需要1枚硬币
}

// 递推公式 1 + min(calcMinCoinNum(money-1), calcMinCoinNum(money-3), calcMinCoinNum(money-5))
// 状态转移方程
func dpMinCoin(money int) int {
	for k := 6; k <= money; k++ {
		status[k] = 1 + min(status[k-1], status[k-3], status[k-5])
	}
	return status[money]
}

func main() {
	fmt.Println(dpMinCoin(9))
	fmt.Println(status)
}
