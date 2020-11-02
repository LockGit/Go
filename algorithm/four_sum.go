/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/18
 * Time: 20:38
 * 给定一个数组[a,b,c,d,e],一个target值，找出满足a+b+c+d==target的不重复集合
 */
package main

import (
	"fmt"
	"sort"
)

func main() {
	list := []int{1, 0, -1, 0, -2, 2, 3}
	target := 0
	sort.Ints(list)
	length := len(list)
	var res [][]int
	for i := 0; i < length-3; i++ {
		if i > 0 && list[i-1] == list[i] {
			continue
		}
		for j := i + 1; j < length-2; j++ {
			if j > i+1 && list[j-1] == list[j] {
				continue
			}
			left := j + 1
			right := length - 1
			for left < right {
				sum := list[i] + list[j] + list[left] + list[right]
				if sum == target {
					data := []int{list[i], list[j], list[left], list[right]}
					res = append(res, data)
					left = left + 1
					right = right - 1
					for left < right && list[left] == list[left-1] {
						left = left + 1
					}
					for left < right && list[right] == list[right+1] {
						right = right - 1
					}
				} else if sum < target {
					left = left + 1
				} else {
					right = right - 1
				}
			}
		}
	}

	fmt.Println(res)
}
