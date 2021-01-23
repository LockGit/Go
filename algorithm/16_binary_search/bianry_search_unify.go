/**
 * Created by lock
二分查找的unify统一：
Knuth 大佬（发明 KMP 算法的那位）都说二分查找：思路很简单，细节是魔鬼。
很多人喜欢拿整型溢出的 bug 说事儿，但是二分查找真正的坑根本就不是那个细节问题，而是在于到底要给 mid 加一还是减一，
while 里到底用 <= 还是 <。
*/
package main

import "fmt"

//基本二分查找
func binarySearch(arr []int, target int) (index int) {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := left + ((right - left) >> 1)
		if target == arr[mid] {
			return mid
		} else if target > arr[mid] {
			left = mid + 1
		} else if target < arr[mid] {
			right = mid - 1
		}
	}
	return -1
}

//左边界的二分
func binarySearchLeft(arr []int, target int) (index int) {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := left + ((right - left) >> 1)
		if target < arr[mid] {
			right = mid - 1
		} else if target > arr[mid] {
			left = mid + 1
		} else if target == arr[mid] {
			//锁定右边解
			right = mid - 1
		}
	}
	if left >= len(arr) || arr[left] != target {
		return -1
	}
	return left
}

//右边界的二分
func binarySearchRight(arr []int, target int) (index int) {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := left + ((right - left) >> 1)
		if target < arr[mid] {
			right = mid - 1
		} else if target > arr[mid] {
			left = mid + 1
		} else if target == arr[mid] {
			//锁定左边界
			left = mid + 1
		}
	}
	if right < 0 || arr[right] != target {
		return -1
	}
	return right
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 5, 5, 5, 5, 5, 6, 8, 9}
	fmt.Println("index is:", binarySearch(arr, 3))
	fmt.Println("index is:", binarySearch(arr, 5))
	fmt.Println("index is:", binarySearchLeft(arr, 5))
	fmt.Println("index is:", binarySearchRight(arr, 5))
	fmt.Println("index is:", binarySearchLeft(arr, 18))
	fmt.Println("index is:", binarySearchRight(arr, 18))
}
