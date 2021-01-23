/**
 * Created by lock
数组前缀和
给定一个数组：arr = []int{2,2,3,1,4}
找到target_sum=4的所有(连续子数组)
*/
package main

import "fmt"

//前缀和，只返回符合的个数，不返回具体的元素
func getChildArrCountByPreSum(arr []int, targetSum int) (c int) {
	preSum := make([]int, len(arr)+1)
	preSum[0] = 0
	for i := 0; i <= len(arr)-1; i++ {
		preSum[i+1] = arr[i] + preSum[i]
	}
	for i := 0; i <= len(arr); i++ {
		for j := 0; j < i; j++ {
			if preSum[i]-preSum[j] == targetSum {
				fmt.Println("find:", j, i, arr[j:i])
				c++
			}
		}
	}
	return
}

//优化版本，结合hash表计算符合的个数
func getChildArrCountByPreSumHash(arr []int, targetSum int) (c int) {
	sumMap := make(map[int]int)
	sumMap[0] = 1
	sumI := 0
	for i := 0; i < len(arr); i++ {
		sumI += arr[i]
		// 这是我们想找的前缀和 nums[0..j]
		sumTarget := sumI - targetSum
		// 如果前面有这个前缀和，则直接更新答案
		if count, ok := sumMap[sumTarget]; ok {
			c += count
		}
		// 把前缀和 nums[0..i] 加入并记录出现次数
		if count, ok := sumMap[sumI]; ok {
			sumMap[sumI] = count + 1
		} else {
			sumMap[sumI] = 1
		}
	}
	return
}

func calcSumIj(arr []int, i, j int) (sum int) {
	for index := i; index <= j; index++ {
		sum += arr[index]
	}
	return
}

func forceGetArrPreSum(arr []int, targetSum int) (res [][]int) {
	for i := 0; i <= len(arr)-1; i++ {
		for j := i; j <= len(arr)-1; j++ {
			if calcSumIj(arr, i, j) == targetSum {
				if i == j {
					res = append(res, []int{arr[i]})
				} else {
					res = append(res, []int{arr[i], arr[j]})
				}
			}
		}
	}
	return
}

func main() {
	arr := []int{1, 2, 2, 3, 4, 2, 2, 3, 1, -5, 9}
	targetSum := 4
	fmt.Println(arr, ",targetSum:", targetSum, "forceGetArrPreSum:", forceGetArrPreSum(arr, targetSum))

	fmt.Println(arr, ",target sum:", targetSum, ",getChildArrCountByPreSum count is:", getChildArrCountByPreSum(arr, targetSum))
	fmt.Println(arr, ",target sum:", targetSum, ",getChildArrCountByPreSumHash count is:", getChildArrCountByPreSumHash(arr, targetSum))
}
