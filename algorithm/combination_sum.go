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
	"fmt"
	"sort"
)

func calcSum(arr []int, target int, aRes *[][]int) {
	sort.Ints(arr)
	var cand []int
	combinationSum(arr, cand, target, aRes)
	return
}

func combinationSum(arr []int, cand []int, target int, aRes *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		*aRes = append(*aRes, cand)
		fmt.Println(cand, aRes)
		return
	}
	for k, v := range arr {
		cand = append(cand, v) // append cand 末尾
		combinationSum(arr[k:], cand, target-v, aRes)
		cand = cand[:len(cand)-1] //移除cand最后一个,回溯状态
	}
}

func main() {
	arr := []int{5, 6, 1, 6, 4, 3}
	target := 4
	aRes := make([][]int, 0)
	calcSum(arr, target, &aRes)
	fmt.Println(aRes)
}
