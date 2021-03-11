/**
 * Created by lock
 * Date: 2021/3/11
给定一个数组x,一个target,找到所有元素之和等于target的值,元素可以重复
	arr := []int{2, 3, 5, 6, 8}
	target := 8
return: [[2 2 2 2] [3 3 2] [5 3] [6 2] [8]]
*/
package main

import "fmt"

var resSumArr [][]int

func calcArrSum(arr []int) (sum int) {
	for _, v := range arr {
		sum += v
	}
	return
}

func calcSumCombine(arr []int, index int, nowArr []int, target int) {
	calcSum := calcArrSum(nowArr)
	if calcSum == target {
		okItem := make([]int, len(nowArr))
		copy(okItem, nowArr)
		resSumArr = append(resSumArr, okItem)
		return
	}
	if calcSum > target {
		return
	}
	for i := 0; i <= index-1; i++ {
		nowArr = append(nowArr, arr[i])
		calcSumCombine(arr, i+1, nowArr, target)
		nowArr = nowArr[0 : len(nowArr)-1]
	}
}

func main() {
	arr := []int{2, 3, 5, 6, 8}
	target := 8
	calcSumCombine(arr, len(arr), []int{}, target)
	fmt.Println(resSumArr)
}
