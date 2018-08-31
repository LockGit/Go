/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/30
 * Time: 17:42
 * 给定一个数组和一个值，在这个数组中移除指定值和返回移除后新的数组长度。不要为其他数组分配额外空间,O(1)复杂度
 */
package main

import "fmt"

func removeArrEle(arr []int, val int) int {
	n := len(arr)
	j := 0
	for i := 0; i < n; i++ {
		if arr[i] != val {
			arr[j] = arr[i]
			j = j + 1
		}
	}
	return j
}

func main() {
	arr := []int{1, 2, 2, 3, 4, 5, 5}
	fmt.Println(removeArrEle(arr, 2))
}
