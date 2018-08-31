/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/30
 * Time: 17:20
 * 一个有序数组，返回去除重复元素后数组的长度。不允许申请新的空间。
 */
package main

import "fmt"

func getDupLength(arr []int) int {
	n := len(arr)
	if n == 1 {
		return 1
	}
	j := 0
	for i := 1; i < n; i++ {
		if arr[i] != arr[j] {
			arr[j+1] = arr[i]
			j = j + 1
		}
	}
	return j
}

func main() {
	arr := []int{1, 2, 2, 3, 4, 5, 5}
	fmt.Println(getDupLength(arr))
}
