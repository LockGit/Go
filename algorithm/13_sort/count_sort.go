/**
 * Created by lock
 * 计数排序 （空间换时间，时间复杂度为O(n)）
 * 待排序的元素在某一个范围[MIN, MAX]之间。
 */
package main

import "fmt"

func countSort(arr []int) []int {
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	countArr := make([]int, max)
	for _, v := range arr {
		countArr[v-1]++
	}
	var newArr []int
	for i := 0; i < len(countArr); i++ {
		countIndex := countArr[i]
		for countIndex > 0 {
			newArr = append(newArr, i+1)
			countIndex--
		}
	}
	return newArr
}

func main() {
	arr := []int{8, 1, 2, 9, 4, 1, 2, 12, 5, 12}
	newArr := countSort(arr)
	fmt.Println("sort arr after is:", newArr)
}
