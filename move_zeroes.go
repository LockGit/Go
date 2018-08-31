/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/30
 * Time: 15:33
 * 将数组中的0移动数组末尾
For example, given nums = [0, 1, 0, 3, 12], after calling your function, nums
should be [1, 3, 12, 0, 0].
*/
package main

import "fmt"

func main() {
	arr := []int{0, 1, 0, 3, 12}
	j := 0
	for i, v := range arr {
		if v != 0 {
			arr[i], arr[j] = arr[j], arr[i]
			j = j + 1
		}
	}
	fmt.Println(arr)
}
