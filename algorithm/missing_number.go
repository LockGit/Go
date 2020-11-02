/**
* Created by GoLand.
* User: lock
* Date: 2018/8/30
* Time: 15:26
* 寻找连续数组中缺失的数
For example,
Given nums = [0, 1, 3] return 2.
*/
package main

import "fmt"

func main() {
	arr := []int{0, 1, 3}
	n := len(arr)
	s := n * (n + 1) / 2 //这是一个公式
	res := 0
	for _, v := range arr {
		res = res + v
	}
	fmt.Println(s - res)
}
