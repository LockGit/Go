/**
 * Created by lock
给定一个arr,求子集：
arr = [1,2,3,4]
*/
package main

import "fmt"

var subsetRes [][]int

func traceBackSubset(arr []int, start int, track []int) {
	//slice是引用
	copyTrack := make([]int, len(track))
	copy(copyTrack, track)
	subsetRes = append(subsetRes, copyTrack)
	for i := start; i < len(arr); i++ {
		track = append(track, arr[i])
		traceBackSubset(arr, i+1, track)
		track = track[:len(track)-1]
	}
}

func main() {
	arr := []int{1, 2, 3, 4}
	//fmt.Println(arr[:3])
	var track []int
	traceBackSubset(arr, 0, track)
	fmt.Println(arr, ",subset is:", subsetRes)
}
