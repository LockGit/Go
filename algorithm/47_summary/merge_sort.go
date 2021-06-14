/**
 * Created by lock
 * Date: 2021/3/12
[2,4,6],[1,2,5]
两个有序数组合并后： 1,2,2,4,5,6
*/
package main

import "fmt"

func mergeSort(a1, a2 []int) (arr []int) {
	for len(a1) > 0 && len(a2) > 0 {
		x := a1[0]
		y := a2[0]
		if x >= y {
			arr = append(arr, y)
			a2 = a2[1:]
		} else {
			arr = append(arr, x)
			a1 = a1[1:]
		}
	}
	if len(a1) > 0 {
		arr = append(arr, a1...)
	}
	if len(a2) > 0 {
		arr = append(arr, a2...)
	}
	return
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	a1 := []int{2, 4, 6}
	a2 := []int{1, 2, 5}
	fmt.Println(mergeSort(a1, a2))
}
