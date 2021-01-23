/**
* Created by lock
积水问题---用一个数组表示一个条形图，问这个条形图最多能接多少水。
arr=[0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}]
输出：6
*/
package main

import "fmt"

func getLeftRightMax(arr []int, index int) (leftMax, rightMax int) {
	leftMax = arr[0]
	for i := 0; i < index; i++ {
		if arr[i] > leftMax {
			leftMax = arr[i]
		}
	}
	rightMax = arr[len(arr)-1]
	for i := index; i <= len(arr)-1; i++ {
		if arr[i] > rightMax {
			rightMax = arr[i]
		}
	}
	return
}

func getMin(a, b int) (minVal int) {
	if a > b {
		return b
	}
	return a
}

func getMax(a, b int) (maxVal int) {
	if a > b {
		return a
	}
	return b
}

func waterArea(arr []int) (v int) {
	//v[i]=min(max(left),max(right))-height[i]
	for i := 1; i < len(arr)-1; i++ {
		leftMax, rightMax := getLeftRightMax(arr, i)
		vi := getMin(leftMax, rightMax) - arr[i]
		if vi > 0 {
			v += vi
		}
	}
	return
}

//备忘录优化版本,提前计算好左右max
func waterAreaImprove(arr []int) (v int) {
	leftMax := make([]int, len(arr))
	rightMax := make([]int, len(arr))
	leftMax[0] = arr[0]
	rightMax[len(arr)-1] = arr[len(arr)-1]
	for i := 1; i <= len(arr)-1; i++ {
		leftMax[i] = getMax(arr[i], leftMax[i-1])
	}
	for i := len(arr) - 2; i >= 0; i-- {
		rightMax[i] = getMax(arr[i], rightMax[i+1])
	}
	for i := 0; i <= len(arr)-1; i++ {
		v += getMin(leftMax[i], rightMax[i]) - arr[i]
	}
	return
}

//使用双指针，边走边算，节省下空间复杂度
func waterAreaPointer(arr []int) (v int) {
	left := 0
	right := len(arr) - 1
	leftMax := arr[0]
	rightMax := arr[len(arr)-1]
	for left <= right {
		leftMax = getMax(leftMax, arr[left])
		rightMax = getMax(rightMax, arr[right])
		if leftMax < rightMax {
			v += leftMax - arr[left]
			left++
		} else {
			v += rightMax - arr[right]
			right--
		}
	}
	return
}

func main() {
	arr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(arr, ",waterArea v is:", waterArea(arr))
	fmt.Println(arr, ",waterAreaImprove v is:", waterAreaImprove(arr))
	fmt.Println(arr, ",waterAreaPointer v is:", waterAreaPointer(arr))
	//最长回文串
}
