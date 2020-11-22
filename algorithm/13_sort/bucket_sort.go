/**
 * Created by lock
 * 桶排序（是一种稳定的排序，时间复杂度O(n)）
 * 将要排序的数据分到几个有序的桶里，每个桶里的数据再单独进行快速排序。
 * 桶内排完序之后，再把每个桶里的数据按照顺序依次取出
 * 待排序的元素在某一个范围且数据不大
 */
package main

import (
	"fmt"
	"sort"
)

func getMax(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

func bucketSort(arr []int) []int {
	max := getMax(arr)
	bucket := make([][]int, max)
	length := len(arr)
	for i := 0; i < length; i++ {
		bucketIndex := arr[i] * (length - 1) / max
		bucket[bucketIndex] = append(bucket[bucketIndex], arr[i])
	}
	pos := 0
	for i := 0; i < length; i++ {
		bucketLength := len(bucket[i])
		if bucketLength > 0 {
			sort.Ints(bucket[i]) //针对bucket里元素快速排序
			copy(arr[pos:], bucket[i])
			pos = pos + bucketLength
		}
	}
	return arr
}
func main() {
	//分数0~100
	arr := []int{72, 11, 82, 32, 44, 13, 17, 95, 54, 28, 79, 56}
	newArr := bucketSort(arr)
	fmt.Println("bucket sort is:", newArr)
}
