/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/26
 * Time: 17:30
 * 给一个数组，其中数组在下标i处的值为A[i]，坐标(i,A[i])和坐标(i,0)构成一条垂直于坐标轴x的直线。现任取两条垂线和x轴组成四边形容器。问其中盛水量最大为多少？
	题目意思就是给一个数组，如[3,4,5,6,8,14,9]，
	其中的3代表从[1,0]到[1,3]的一条线段，
	4代表从[2,0]到[2,4]的线段，
	以此类推，要求求出这些线段中哪两条和x轴一起可以围成一个最大的蓄水池。
*/
package main

import (
	"fmt"
	"math"
)

func maxArea(list []int) int {
	n := len(list)
	i, j := 0, n-1
	maxMulArea := 0
	for i < j {
		maxMulArea = int(math.Max(float64(maxMulArea), float64(j-i)*math.Min(float64(list[i]), float64(list[j]))))
		if list[i] <= list[j] {
			i = i + 1
		} else {
			j = j - 1
		}
	}
	return maxMulArea
}

func main() {
	arr := []int{3, 4, 5, 6, 8, 14, 9}
	fmt.Println(maxArea(arr))
}
