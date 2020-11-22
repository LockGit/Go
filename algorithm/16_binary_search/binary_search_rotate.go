/**
 * Created by lock
 * 4,5,6,7,0,1,2
 * 有序数组是一个循环有序数组,求“值等于给定值”的二分查找
 */
package main

import "fmt"

func searchRotate(arr []int, target int) (val, index int) {
	left := 0
	right := len(arr) - 1
	//如果中间的数小于最右边的数，则右半段是有序的
	//若中间数大于最右边数，则左半段是有序的
	//只要在有序的半段里用首尾两个数组来判断目标值是否在这一区域内
	for left <= right {
		mid := left + ((right - left) >> 1)
		if arr[mid] == target {
			return arr[mid], mid
		} else if arr[mid] < arr[right] {
			//右边半段有序
			if arr[mid] < target && arr[right] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			//左边半段有序
			if arr[left] <= target && arr[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return target, -1
}

func main() {
	arr := []int{4, 5, 6, 7, 0, 1, 2}
	val, index := searchRotate(arr, 2)
	fmt.Println(val, index)
}
