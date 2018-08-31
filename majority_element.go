/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/30
 * Time: 14:30
 * 从一个数组中找出出现半数以上的元素
 */
package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 1, 2, 2, 2, 3}
	half := len(arr) / 2
	d := make(map[int]interface{})
	for _, v := range arr {
		if d[v] == nil {
			d[v] = 1
		} else {
			d[v] = d[v].(int) + 1
		}
	}
	for k, item := range d {
		if item.(int) >= half {
			fmt.Println(k)
		}
	}
}
