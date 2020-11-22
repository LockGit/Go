/**
 * Created by lock
 */
package main

import "fmt"

// sort.go
func bubbleSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		swapFlag := false
		for j := i; j < length; j++ {
			if arr[i] > arr[j] {
				tmp := arr[i]
				arr[i] = arr[j]
				arr[j] = tmp
				swapFlag = true
			}
		}
		if !swapFlag {
			break
		}
	}
	return arr
}

func insertSort(arr []int) []int {
	length := len(arr)
	for i := 1; i < length; i++ {
		v := arr[i]
		j := i - 1
		//4 8 1 2 7
		for ; j >= 0; j-- {
			if v < arr[j] {
				arr[j+1] = arr[j] //后移 a[1]=a[0]
			} else {
				break
			}
		}
		arr[j+1] = v //插入
	}
	return arr
}

func selectSort(arr []int) []int {
	length := len(arr)
	//这个用length 和 length-1 的区别是最后一次，不需要多一步swap
	for i := 0; i < length-1; i++ {
		selectMinIndex := i
		for j := i + 1; j < length; j++ { //访问未排序的元素
			if arr[j] < arr[selectMinIndex] {
				selectMinIndex = j
			}
		}
		//swap
		if selectMinIndex != i {
			arr[i], arr[selectMinIndex] = arr[selectMinIndex], arr[i]
		}
	}
	return arr
}

func main() {
	arr := []int{2, 8, 7, 1, 5, 6}
	bubbleSort(arr)
	fmt.Println("bubbleSort res:", arr)

	arr2 := []int{7, 18, 2, 1, 8, 5}
	insertSort(arr2)
	fmt.Println("insertSort res:", arr2)

	arr3 := []int{9, 1, 2, 4, 8, 5, 3, 3, 3, 1}
	selectSort(arr3)
	fmt.Println("selectSort res:", arr3)
}
