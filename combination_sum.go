/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/25
 * Time: 22:13
 * 回溯算法|递归实现
example:
calcSum([5, 6, 1, 6, 4, 3], 4)

return:
[[1, 1, 1, 1], [1, 3], [4]]
 */
package main

import (
	"sort"
	"fmt"
)

func calcSum(arr []int, target int) [][]int {
	sort.Ints(arr)
	var res [][]int
	var cand [] int
	combinationSum(arr, cand, target, res)
	return res
}

func combinationSum(arr []int, cand []int, target int, res [][]int) {
	if target < 0 {
		return
	} else if target == 0 {
		res = append(res, cand)
		fmt.Println(res)
	} else {
		for k, v := range arr {
			cand = append(cand, v)
			combinationSum(arr[k:], cand, target-v, res)
			cand = cand[:len(cand)-1]
		}
	}
}

func main() {
	arr := [] int{5, 6, 1, 6, 4, 3}
	target := 4
	calcSum(arr, target)
}
