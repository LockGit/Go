/**
 * Created by lock
 * Date: 2020/10/30
 * 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。说明：解集不能包含重复的子集。
输入: nums = [1,2,3]

输出:
[1]
[1 2]
[1 2 3]
[1 3]
[2]
[2 3]
[3]

*/
package main

import (
	"fmt"
)

func numCombineTrack(aTrack, aChoice []int, startIndex int, res *[][]int) {
	if len(aTrack) > 0 {
		cpy := make([]int, len(aTrack))
		copy(cpy, aTrack)
		*res = append(*res, cpy)
	}
	for i := startIndex; i < len(aChoice); i++ {
		aTrack = append(aTrack, aChoice[i])
		numCombineTrack(aTrack, aChoice, i+1, res)
		aTrack = aTrack[0 : len(aTrack)-1]
	}
}

func main() {
	aNum := []int{1, 2, 3}
	var res [][]int
	var aTrack []int
	numCombineTrack(aTrack, aNum, 0, &res)
	for _, v := range res {
		fmt.Println(v)
	}
}
