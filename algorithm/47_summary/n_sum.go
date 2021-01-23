/**
 * Created by lock
给定一个target,一个[]arr,计算n个数之和=target
例如：
输入一个数组 nums 和一个目标和 target，请你返回 nums 中能够凑出 target 的两个元素的值，
比如输入 nums = [1,3,5,6], target = 9，那么算法返回两个元素 [3,6]。
可以假设只有且仅有一对儿元素可以凑出 target。
*/
package main

import (
	"fmt"
	"sort"
)

func nSum(arr []int, start int, n int, target int) (res [][]int) {
	if n < 2 || len(arr) < 2 {
		return res
	}
	if n == 2 {
		//base case
		low := start
		height := len(arr) - 1
		for low < height {
			sum := arr[low] + arr[height]
			left := arr[low]
			right := arr[height]
			if sum < target {
				for low < height && left == arr[low] {
					low++
				}
			} else if sum > target {
				for low < height && right == arr[height] {
					height--
				}
			} else if target == sum {
				res = append(res, []int{arr[low], arr[height]})
				for low < height && left == arr[low] {
					low++
				}
				for low < height && right == arr[height] {
					height--
				}
			}
		}
		return
	} else {
		for i := start; i <= len(arr)-1; i++ {
			subSum := nSum(arr, i+1, n-1, target-arr[i])
			for _, item := range subSum {
				item = append(item, arr[i])
				res = append(res, item)
			}
		}
	}
	return
}

func main() {
	arr := []int{1, 21, 3, 4, 5, 6, 7, 7, 8, 8, 9, 15, 18}
	sort.Ints(arr)
	target := 16
	fmt.Println(arr, ",target is:", target, ",res is:", nSum(arr, 0, 4, target))
}
