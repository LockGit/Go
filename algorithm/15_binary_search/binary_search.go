/**
 * Created by lock
 * 二分查找，递归+非递归
 * 有序无重复的数组
 */
package main

import "fmt"

func recursionBinarySearch(arr []int, left, right, value int) (val, index int) {
	if left > right {
		return val, -1
	}
	//mid := (left + right) / 2
	mid := left + ((right - left) >> 1) //防止整形溢出
	if value == arr[mid] {
		return value, mid
	}
	if arr[mid] < value {
		left = mid + 1
	}
	if arr[mid] > value {
		right = mid - 1
	}
	return recursionBinarySearch(arr, left, right, value)
}

func whileBinarySearch(arr []int, num int) (val, index int) {
	left := 0
	right := len(arr) - 1
	for left <= right {
		//mid := (left + right) / 2
		mid := left + ((right - left) >> 1) //防止整形溢出
		if arr[mid] == num {
			return num, mid
		}
		if num < arr[mid] {
			right = mid - 1
		}
		if num > arr[mid] {
			left = mid + 1
		}
	}
	return num, -1
}

func main() {
	//有序无重复的数组
	arr := []int{0, 1, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	val, index := recursionBinarySearch(arr, 0, len(arr)-1, 9)
	fmt.Println("recursion find res:", val, ",index:", index)

	val, index = whileBinarySearch(arr, 9)
	fmt.Println("for loop find res:", val, ",index:", index)
}
