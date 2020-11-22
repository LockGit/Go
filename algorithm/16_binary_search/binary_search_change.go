/**
 * Created by lock
 * 二分查找的变体
 * 假设查找的元素中存在重复元素
 * 1，查找第一个值等于给定值的元素
 * 2，查找最后一个值等于给定值的元素
 * 3，查找第一个大于等于给定值的元素
 * 4，查找最后一个小于等于给定值的元素
 */
package main

import "fmt"

//1，查找第一个值等于给定值的元素
func searchFirstEq(arr []int, target int) (val int, index int) {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := left + ((right - left) >> 1) // (left+right)/2 ，比较好的编译器应该自己会优化这个代码吧
		if arr[mid] > target {
			right = mid - 1
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			if (mid == 0) || arr[mid-1] != target {
				return arr[mid], mid
			} else {
				right = mid - 1
			}
		}
	}
	return val, -1
}

//2，查找最后一个值等于给定值的元素
func searchLastEq(arr []int, target int) (val, index int) {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := left + ((right - left) >> 1) //(left+right)/2
		if arr[mid] > target {
			right = mid - 1
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			if (arr[len(arr)-1] == target) || (arr[mid+1] != target) {
				return arr[mid], mid
			} else {
				left = mid + 1
			}
		}
	}
	return val, -1
}

//3，查找第一个大于等于给定值的元素
func searchFirstGt(arr []int, target int) (val, index int) {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := left + ((right - left) >> 1)
		if arr[mid] >= target {
			if mid == 0 || arr[mid-1] < target {
				return arr[mid], mid
			} else {
				right = mid - 1
			}
		} else {
			left = mid + 1
		}
	}
	return val, -1
}

//4，查找最后一个小于等于给定值的元素
func searchLastLt(arr []int, target int) (val, index int) {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := left + ((right - left) >> 1)
		if arr[mid] <= target {
			if mid == len(arr)-1 || arr[mid+1] > target {
				return arr[mid], mid
			} else {
				left = mid + 1
			}
		} else {
			right = mid - 1
		}
	}
	return val, -1
}

func main() {
	arr := []int{1, 2, 3, 3, 6, 6, 6, 7, 8, 8, 8, 8, 8, 11, 12, 15}
	val, index := searchFirstEq(arr, 6)
	fmt.Println("查找第一个值等于给定值的元素:", val, ",下标:", index)

	val2, index2 := searchLastEq(arr, 6)
	fmt.Println("查找最后一个值等于给定值的元素:", val2, ",下标:", index2)

	val3, index3 := searchFirstGt(arr, 3)
	fmt.Println("查找第一个大于等于给定值的元素:", val3, ",下标:", index3)

	val4, index4 := searchLastLt(arr, 8)
	fmt.Println("查找最后一个小于等于给定值的元素:", val4, ",下标:", index4)
}
