/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/31
 * Time: 15:32
 * 一个list中只有一个数出现了1次，其他数2次，找出这个出现一次的数字
 */
package main

import "fmt"

func main() {
	arr := []int{1, 1, 2, 2, 5, 3, 3}
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		res = res ^ arr[i]
	}
	fmt.Println(res)
}
