/**
 * Created by lock
0,1背包
给你一个可装载重量为 W 的背包和 N 个物品，每个物品有重量和价值两个属性。其中第 i 个物品的重量为 wt[i]，价值为 val[i]，
现在让你用这个背包装物品，最多能装的价值是多少？
示例：
N = 3, W = 4
wt = [2, 1, 3]
val = [4, 2, 3]
算法返回 6，选择前两件物品装进背包，总重量 3 小于 W，可以获得最大价值 6。
*/
package main

import "fmt"

func maxVal(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func pkg01MaxVal(w, n int, wt []int, val []int) (v int) {
	// base case 已初始化
	dp := make([][]int, n+1)
	for k, _ := range dp {
		item := make([]int, w+1)
		dp[k] = item
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= w; j++ {
			if j-wt[i-1] < 0 {
				// 这种情况下只能选择不装入背包
				dp[i][j] = dp[i-1][j]
			} else {
				// 装入或者不装入背包，择优
				dp[i][j] = maxVal(
					dp[i-1][j-wt[i-1]]+val[i-1],
					dp[i-1][j],
				)
			}
		}
	}
	return dp[n][w]
}
func main() {
	n := 3
	w := 4
	wt := []int{2, 1, 3}
	val := []int{4, 2, 3}
	fmt.Println(pkg01MaxVal(w, n, wt, val))
}
