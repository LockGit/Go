/**
 * Created by lock
求组合
给定一个arr=[1,2,3,4]，c=2
1,2
1,3
1,4
2,3
2,4
3,4
*/
package main

import "fmt"

var combineRes [][]int

func traceBackCombine(arr []int, start, stop, k int, track []int) {
	if len(track) == k {
		copyTrack := make([]int, len(track))
		copy(copyTrack, track)
		combineRes = append(combineRes, copyTrack)
		return
	}
	for i := start; i <= stop; i++ {
		track = append(track, arr[i])
		traceBackCombine(arr, i+1, stop, k, track)
		track = track[:len(track)-1]
	}
}

func main() {
	arr := []int{1, 2, 3, 4}
	k := 2
	traceBackCombine(arr, 0, len(arr)-1, k, []int{})
	fmt.Println(arr, ",combine is:", combineRes)
}
