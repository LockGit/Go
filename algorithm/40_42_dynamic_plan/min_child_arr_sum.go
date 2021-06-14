/**
 * Created by lock
 * Date: 2021/3/12
动态规划之最小连续子数组之和：
arr = [1,-1,-2,1]
minSum=(-2)+(1)=-3
*/
package main

import (
	"fmt"
	"math"
)

func minAb(a ...int) int {
	m := a[0]
	for _, v := range a {
		if v < m {
			m = v
		}
	}
	return m
}

func minChildArrSum(arr []int) (minSum int) {
	if len(arr) <= 0 {
		return 0
	}
	dp := make([]int, len(arr))
	dp[0] = arr[0]
	for i := 1; i <= len(arr)-1; i++ {
		dp[i] = minAb(arr[i], dp[i-1]+arr[i])
	}
	minSum = math.MaxInt64
	for _, v := range dp {
		if v < minSum {
			minSum = v
		}
	}
	return minSum
}

func main() {
	arr := []int{1, -1, -2, 1, -5}
	fmt.Println("min child arr mul sum:", minChildArrSum(arr))
}
