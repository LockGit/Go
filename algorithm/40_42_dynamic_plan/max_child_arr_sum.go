/**
 * Created by lock
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
arr = [1,-2,3,10,-4,7,2,-5]
最大子数组和为:3+10+(-4)+7+2=18
*/
package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMaxChildArrSum(arr []int) (sum int) {
	if len(arr) == 0 {
		return 0
	}
	dp := make([]int, len(arr))
	fmt.Println("init dp table:", dp)
	dp[0] = arr[0]
	for i := 1; i <= len(arr)-1; i++ {
		dp[i] = max(arr[i], arr[i]+dp[i-1])
	}
	maxSum := 0
	for j := 0; j < len(dp); j++ {
		if dp[j] > maxSum {
			maxSum = dp[j]
		}
	}
	return maxSum
}

//dp table 状态压缩
func getMaxChildArrSumOptimize(arr []int) (sum int) {
	if len(arr) == 0 {
		return 0
	}
	maxSum := 0
	preSum := arr[0]
	//最大，当前两个状态有关
	for i := 1; i <= len(arr)-1; i++ {
		dpNowSum := max(arr[i], arr[i]+preSum)
		preSum = dpNowSum
		if dpNowSum > maxSum {
			maxSum = dpNowSum
		}
	}
	return maxSum
}

func main() {
	arr := []int{1, -2, 3, 10, -4, 7, 2, -5}
	fmt.Println("get max child arr sum is:", getMaxChildArrSum(arr))
	fmt.Println("get max child arr optimize sum is:", getMaxChildArrSumOptimize(arr))
}
