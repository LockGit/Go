/**
* Created by GoLand.
* User: lock
* Date: 2018/8/29
* Time: 12:00
* 从一个未经排序的数组中找出第k大的元素。注意是排序之后的第k大，而非第k个不重复的元素。
  堆排序是利用堆这种数据结构而设计的一种排序算法，堆排序是一种选择排序，它的最坏，最好，平均时间复杂度均为O(nlogn)，它也是不稳定排序。
  https://www.cnblogs.com/chengxiao/p/6129630.html
*/
package main

import "fmt"

func find(arr []int, k int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		pivot := partition(arr, left, right)
		if pivot == k-1 {
			return arr[pivot]
		} else if pivot < k-1 {
			left = pivot + 1
		} else {
			right = pivot - 1
		}
	}
	return 0
}

func partition(arr []int, left int, right int) int {
	pivot := right
	j := left
	for i := left; i < right; i++ {
		if arr[i] > arr[pivot] {
			arr[i], arr[j] = arr[j], arr[i]
			j = j + 1
		}
	}
	arr[j], arr[pivot] = arr[pivot], arr[j]
	return j
}

func main() {
	arr := []int{3, 2, 1, 5, 6, 4}
	k := 2
	fmt.Println(find(arr, k))
}
