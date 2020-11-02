/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/27
 * Time: 18:45
 * 在旋转排序数组中搜索(指定的目标)
	假定一个有序数组在一个预先不知道的轴点地方被旋转。
	例如，0 1 2 4 5 6 7可能会变为4 5 6 7 0 1 2。
	给你一个目标值去搜索，如果在数组中找到了则返回它的索引，否则返回-1。
	你可以假定数组中没有重复的元素。
*/
package main

import "fmt"

func findMin(list []int) int {
	n := len(list)
	left := 0
	right := n - 1
	if n == 1 || list[left] < list[right] {
		return list[0]
	}
	for left <= right {
		mid := left + (right-left)/2
		if mid > 0 && list[mid-1] > list[mid] {
			return list[mid]
		} else if list[mid] > list[right] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	list := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(findMin(list))
}
