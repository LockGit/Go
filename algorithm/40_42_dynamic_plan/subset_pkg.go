/**
 * Created by lock
子集背包~~~两个子集的元素和相等
给定一个只包含正整数的非空数组。是否可以将这个数组分隔成2个子集，使得两个子集的元素和相等。

示例：
输入：[1,5,11,5]
输出：true
解释：数组可以分隔成 [1,5,5] 和 [11]

示例：
输入：[1,2,3,5]
输出：false
*/
package main

import "fmt"

//两个和相等的数组划分---转换成背包问题
func subsetPkg(arr []int) (b bool) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	if sum%2 == 1 {
		return false
	}
	//我们想求的最终答案就是 dp[n][sum/2]
	halfSum := sum / 2
	//第i个物品，背包容量为w, dp[i][w]=true/false, 是否能装满
	// baseCase dp[i][0]=true
	dp := make([][]bool, len(arr)+1)
	for k, _ := range dp {
		item := make([]bool, halfSum+1)
		dp[k] = item
		dp[k][0] = true
	}
	for i := 1; i <= len(arr); i++ {
		for w := 1; w <= halfSum; w++ {
			if w-arr[i-1] < 0 {
				//背包容量不足，不放入背包
				dp[i][w] = dp[i-1][w]
			} else {
				//不装入或者装入背包
				//如果 w - nums[i-1] 的重量可以被恰好装满，那么只要把第 i 个物品装进去，也可恰好装满 w 的重量; 否则的话，重量 w 肯定是装不满
				dp[i][w] = dp[i-1][w] || dp[i-1][w-arr[i-1]]
			}
		}
	}
	return dp[len(arr)][halfSum]
}

func main() {
	arr := []int{1, 5, 11, 5}
	fmt.Println(arr, ", subset can split is:", subsetPkg(arr))
	arr = []int{1, 2, 3, 5}
	fmt.Println(arr, ", subset can split is:", subsetPkg(arr))

}
