/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/17
 * Time: 17:42
 * 在一个数组中，找到3个元数之和等于0的元素列表组合
 */
package main

import (
	"fmt"
	"sort"
)

func main() {
	list := []int{1, -1, 1, -2, 2, 0}
	fmt.Println("原始数据列表:", list)
	sort.Ints(list)
	var res [][]int
	for i := 0; i < len(list)-2; i++ {
		if i == 0 || (i > 0 && list[i-1] != list[i]) {
			left := i + 1
			right := len(list) - 1
			for left < right {
				sum := list[i] + list[left] + list[right]
				if sum == 0 {
					data := []int{list[i], list[left], list[right]}
					res = append(res, data)
					left = left + 1
					right = right - 1
					for left < right && list[left] == list[left-1] {
						left = left + 1
					}
					for right > left && list[right] == list[right+1] {
						right = right - 1
					}
				} else if sum < 0 {
					left = left + 1
				} else {
					right = right - 1
				}
			}
		}
	}

	for _, v := range res {
		fmt.Println(v, "元素之和为0")
	}

}
