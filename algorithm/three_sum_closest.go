/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/18
 * Time: 17:38
 * 给一个整数数组，找到三个数的和与给定target的值距离最短的那个和
 * For example, given array S = {-1 2 1 -4}, and target = 1.
 */
package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	list := []int{-1, 2, 1, -4}
	target := 10
	sort.Ints(list)
	var data []int
	res := list[0] + list[1] + list[2]
	data = []int{list[0], list[1], list[2]}
	length := len(list)
	for i := 0; i < length-2; i++ {
		left := i + 1
		right := length - 1
		for left < right {
			s := list[i] + list[left] + list[right]
			if math.Abs(float64(s-target)) < math.Abs(float64(res-target)) {
				res = s
				data = []int{list[i], list[left], list[right]}
			} else if s == target {
				res = s
				break
			} else if s < target {
				left = left + 1
			} else {
				right = right - 1
			}
		}
	}

	fmt.Println(res)
	fmt.Println(data)
}
