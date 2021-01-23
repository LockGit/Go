/**
 * Created by lock
完全背包~~~零钱兑换问题
给定不同面额的硬币和一个总金额，写出函数来计算可以凑成总金额的硬币组合数。假设每种硬币有无限个。
示例：
输入：amount=5,coins=[1,2,5]
输出：4
解释有4种可以凑成总金额等于5的方法：
5 = 5
5 = 2+2+1
5 = 2+1+1+1
5 = 1+1+1+1+1
*/
package main

import "fmt"

func fullPkg(amount int, coins []int) (count int) {
	//dp[i][j]=x , 若只使用 coins 中的前 i 个硬币的面值，若想凑出金额 j，有 dp[i][j] 种凑法
	//baseCase dp[0][..]=0 即不用任何面值的硬币，就无法凑出目标金额， dp[..][0],如果凑出的目标金额为0,dp[0][..]=1
	//状态
	dp := make([][]int, len(coins)+1)
	for k, _ := range dp {
		item := make([]int, amount+1)
		item[0] = 1
		dp[k] = item
		fmt.Println(item)
	}
	//递推
	for i := 1; i <= len(coins); i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i-1] >= 0 {
				//放进与不放进背包的选择之和，共有多少种凑法,两种选择的之和
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			} else {
				//不放进背包，继承之前的结果
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(coins)][amount]
}

func main() {
	amount := 5
	coins := []int{1, 2, 5}
	fmt.Println("amount is:", amount, ",coins is:", coins, ",have count is:", fullPkg(amount, coins))
}
