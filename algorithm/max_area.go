/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/29
 * Time: 18:21
 * 单调栈---最大面积
	For example,
	Given heights = [2, 1, 5, 6, 2, 3],
	return 10.
	大致题意：
	给出一组数据，分别代表柱状图的高度，求出柱状图中最大矩形的面积。每个单位长度默认为1。
	https://blog.csdn.net/x_kh_2001/article/details/80465027
*/
package main

import (
	"fmt"
	"math"
)

func getMaxArea(height []int) int {
	if height == nil {
		return 0
	}
	if len(height) == 1 {
		return height[0]
	}
	var stack []int
	maxArea := 0
	n := len(height)
	var w int
	for i := 0; i < n+1; i++ {
		for len(stack) > 0 && (i == n || height[stack[len(stack)-1]] > height[i]) {
			pop := stack[len(stack)-1]
			stack = append([]int{}, stack[0:len(stack)-1]...)
			h := height[pop]
			if len(stack) > 0 {
				w = i - stack[len(stack)-1] - 1
			} else {
				w = i
			}
			maxArea = int(math.Max(float64(maxArea), float64(h*w)))
		}
		stack = append(stack, i)
	}
	return maxArea
}

func main() {
	height := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(getMaxArea(height))
}
