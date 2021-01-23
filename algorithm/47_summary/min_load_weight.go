/**
* Created by lock
基于二分搜索的问题---最小运载能力
有重量weights [1,2,3,4,5,6,7,8,9,10] , d=5
返回如果需要在在d天内将weights中所有物品，通过船装载传输完成，那么船的最低运载能力为多少？

Input: weights = [1,2,3,4,5,6,7,8,9,10], d = 5
Output: 15
Explanation:
A ship capacity of 15 is the minimum to ship all the packages in 5 days like this:
1st day: 1, 2, 3, 4, 5
2nd day: 6, 7
3rd day: 8
4th day: 9
5th day: 10
*/
package main

import (
	"fmt"
	"math"
)

func getMinSumArr(arr []int) (minCap, sumCap int) {
	minCap = math.MinInt64
	for _, v := range arr {
		sumCap += v
		if v > minCap {
			minCap = v
		}
	}
	return
}

func canLoadFinish(arr []int, cap, d int) (b bool) {
	//因为物品不能分隔，那么不能使用这种
	//for i := 1; i <= d; i++ {
	//	sumWeight = sumWeight - cap
	//	if sumWeight <= 0 {
	//		return true
	//	}
	//}
	i := 0
	for day := 1; day <= d; day++ {
		maxCap := cap
		for {
			maxCap = maxCap - arr[i]
			if maxCap >= 0 {
				i++
				if i == len(arr) {
					return true
				}
			} else {
				break
			}
		}
	}
	return false
}

func getMinLoadCap(arr []int, d int) (minLoad int) {
	//right in [min[arr]~~~sum(arr)]
	left, right := getMinSumArr(arr)
	for left < right {
		midCap := left + ((right - left) >> 1)
		if canLoadFinish(arr, midCap, d) {
			right = midCap
		} else {
			left = midCap + 1
		}
	}
	return left
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	d := 5
	fmt.Println(arr, ",d=", d, ",min load cap is:", getMinLoadCap(arr, d))
}
