/**
 * Created by lock
 * Date: 2020/3/16
 * 全排列 递归+回溯
 */
package main

import (
	"fmt"
	"strings"
)

func calcCombination(arr []string, start int, choice *[]string) {
	if start == len(arr) {
		*choice = append(*choice, strings.Join(arr, ""))
		return
	}
	for i := start; i < len(arr); i++ {
		//交换位置
		if i != start {
			tmp := arr[start]
			arr[start] = arr[i]
			arr[i] = tmp
		}
		calcCombination(arr, start+1, choice)
		//回溯
		if i != start {
			tmp := arr[start]
			arr[start] = arr[i]
			arr[i] = tmp
		}
	}
}

func main() {
	arr := []string{
		"a", "b", "c",
	}
	res := make([]string, 0)
	calcCombination(arr, 0, &res)
	fmt.Println(res)
}
