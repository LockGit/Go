/**
 * Created by lock
 * 归并排序 (这两个是分治思想）
 */
package main

import "fmt"

func merge(arr []int, start, mid, end int) {
	tmp := make([]int, end-start+1)
	i := start
	j := mid + 1
	k := 0
	for ; i <= mid && j <= end; k++ {
		if arr[i] <= arr[j] {
			tmp[k] = arr[i]
			i++
		} else {
			tmp[k] = arr[j]
			j++
		}
	}
	//判断哪个子数组中有剩余的数据,将剩余的数据拷贝到临时数组tmp
	for ; i <= mid; i++ {
		tmp[k] = arr[i]
		k++
	}
	//判断哪个子数组中有剩余的数据,将剩余的数据拷贝到临时数组tmp
	for ; j <= end; j++ {
		tmp[k] = arr[j]
		k++
	}
	//复制到原始数组对应位置
	copy(arr[start:end+1], tmp)
}

func mergeSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	mergeSort(arr, start, mid)
	mergeSort(arr, mid+1, end)
	merge(arr, start, mid, end)
}

func main() {
	arr2 := []int{9, 1, 2, 7, 4, 5, 8}
	mergeSort(arr2, 0, len(arr2)-1)
	fmt.Println("merge sort res:", arr2)
}
