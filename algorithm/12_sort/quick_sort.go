/**
 * Created by lock
 * 快速排序
 */
package main

import "fmt"

func partition(start, end int, arr []int) int {
	//选取最后一位数字作为对比数字
	pivot := arr[end]
	i := start
	//原地排序,空间复杂度得是O(1)
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			//加入已处理区间的尾部
			if i != j {
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
	}
	arr[i], arr[end] = arr[end], arr[i]
	return i
}
func quickSort(left, right int, arr []int) {
	if left >= right {
		return
	}
	pivot := partition(left, right, arr)
	quickSort(left, pivot-1, arr)
	quickSort(pivot+1, right, arr)
}

func main() {
	arr1 := []int{8, 1, 2, 6, 3, 9, 7}
	quickSort(0, len(arr1)-1, arr1)
	fmt.Println("quick sort res:", arr1)
}
