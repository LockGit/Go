/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/29
 * Time: 11:47
 * 跳跃游戏
For example:
A = [2,3,1,1,4], return true.
A = [3,2,1,0,4], return false.
数组中的每个数字代表可以跳的步数，看看能不能跳到最后。
*/
package main

import (
	"fmt"
	"math"
)

func canJump(list []int) bool {
	n := len(list)
	if n == 1 {
		return true
	}
	t := 0
	for i := 1; i < n; i++ {
		t = int(math.Max(float64(t), float64(list[i-1]))) - 1
		if t < 0 {
			return false
		}
	}
	return true
}

func main() {
	arr1 := []int{2, 3, 1, 1, 4}
	arr2 := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(arr1))
	fmt.Println(canJump(arr2))
}
