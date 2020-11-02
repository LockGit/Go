/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/30
 * Time: 14:56
 * 求连续的数组值，加和最大
For example, given the array [−2,1,−3,4,−1,2,1,−5,4],
the contiguous subarray [4,−1,2,1] has the largest sum = 6.
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	maxSum := arr[0]
	maxCurrnet := maxSum
	for i := 1; i < len(arr); i++ {
		maxCurrnet = int(math.Max(float64(arr[i]), float64(maxCurrnet+arr[i])))
		maxSum = int(math.Max(float64(maxSum), float64(maxCurrnet)))
	}
	fmt.Println(maxSum)
}
