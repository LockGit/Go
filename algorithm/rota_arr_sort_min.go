/**
 * Created by lock
 * Date: 2020/9/7
 * 查找旋转有序数组最小值
 */
package main

import "fmt"

func findRotaSortArrMinVal(arr []int) (min int) {
	right := len(arr) / 2
	left := right - 1
	for {
		if left < 0 {
			min = arr[right]
			break
		}
		if arr[left] > arr[right] {
			min = arr[right]
			break
		} else {
			right--
			left = right - 1
		}
	}
	return min
}

//O(log(n))
func findRotaSortArrMinVal2(arr []int) (min int) {
	left := 0
	right := len(arr) - 1
	for left < right {
		if arr[left] < arr[right] {
			return arr[left]
		}
		mid := left + (right-left)/2
		fmt.Println("left is:", arr[left], ",right is:", arr[right], ",mid is:", arr[mid])
		if arr[left] > arr[mid] {
			// 最小值在[left, mid]
			right = mid
		} else {
			// 最小值在(mid, right]
			left = mid + 1
		}
	}
	return arr[left]
}

func main() {
	fmt.Println("find rota sort arr min value")
	//0 1,2,3,4,5
	var arr = []int{6, 7, 0, 1, 2, 3, 4, 5}
	//var arr = []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("raw arr is:", arr, "arr length is:", len(arr))
	min := findRotaSortArrMinVal(arr)
	fmt.Println("min value is:", min)
	min2 := findRotaSortArrMinVal2(arr)
	fmt.Println("second func min value is:", min2)
}
