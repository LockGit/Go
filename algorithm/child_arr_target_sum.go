/**
 * Created by lock
 * Date: 2020/11/4
 * 给定一个由 n 个正整数组成的数组和一个正整数 s，求所有大于等于targetSum的子数组,子数组内元素可以重复使用
example:
输入: [2, 3, 1], s = 2
输出:
	[2]
	[1 1]
	min child arr length is: 1
解释: 子数组 [2] 是该条件下的最小长度子数组。
递归+回溯
*/
package main

import "fmt"

func getSumTargetArr(arr []int, selectIndex int, targetSum int, aWaitNum []int, res *[][]int) {
	if targetSum < 0 {
		return
	}
	if targetSum == 0 {
		cpy := make([]int, len(aWaitNum))
		copy(cpy, aWaitNum)
		*res = append(*res, cpy)
		return
	}
	for i := selectIndex; i < len(arr); i++ {
		num := arr[i]
		aWaitNum = append(aWaitNum, num)
		getSumTargetArr(arr, i, targetSum-num, aWaitNum, res)
		aWaitNum = aWaitNum[:len(aWaitNum)-1]
	}
}

func checkMinLengthChildArr(arr []int, targetSum int) int {
	var childTargetSumArr [][]int
	getSumTargetArr(arr, 0, targetSum, []int{}, &childTargetSumArr)
	if len(childTargetSumArr) == 0 {
		return -1
	}
	min := int(^uint(0) >> 1) //put max to min
	for _, v := range childTargetSumArr {
		fmt.Println(v)
		itemLen := len(v)
		if itemLen < min {
			min = itemLen
		}
	}
	return min
}

func main() {
	arr := []int{2, 3, 1}
	targetSum := 2
	index := checkMinLengthChildArr(arr, targetSum)
	fmt.Println("min child arr length is:", index)
}
