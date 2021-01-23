/**
 * Created by lock
求排列
给定一个arr=[1,2,3,4]
*/
package main

import "fmt"

var permutationRes [][]int

func containsPerm(arr []int, item int) (b bool) {
	for _, v := range arr {
		if item == v {
			return true
		}
	}
	return false
}

func traceBackPermutation(arr []int, track []int) {
	if len(track) == len(arr) {
		copyTrack := make([]int, len(arr))
		copy(copyTrack, track)
		permutationRes = append(permutationRes, copyTrack)
		return
	}
	for i := 0; i <= len(arr)-1; i++ {
		if containsPerm(track, arr[i]) {
			continue
		}
		track = append(track, arr[i])
		traceBackPermutation(arr, track)
		track = track[:len(track)-1]
	}
}
func main() {
	arr := []int{1, 2, 3, 4}
	traceBackPermutation(arr, []int{})
	fmt.Println(arr, ",permutation is:", permutationRes)
}
