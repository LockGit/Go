/**
 * Created by lock
动态规划，最长递增子序列，仅返回最长序列长度
例如：
输入：[10,9,2,5,3,7,101,18]
输出：4
解释：最长的上升序列是[2,3,7,101],它的长度是4
*/
package main

import "fmt"

func getMaxLongChildSeq(seq []int) int {
	if len(seq) == 1 {
		return 1
	}
	dp := make([]int, len(seq))
	dp[0] = 1
	maxK := dp[0]
	for i := 1; i < len(seq); i++ {
		maxLen := dp[i-1]
		for k := 0; k <= i; k++ {
			if seq[i] > seq[k] {
				maxLen = dp[k] + 1
			}
		}
		dp[i] = maxLen
		if dp[i] > maxK {
			maxK = dp[i]
		}
	}
	fmt.Println("dp table is:", dp)
	return maxK
}

func main() {
	seq := []int{2, 9, 3, 6, 5, 1, 7}
	n := getMaxLongChildSeq(seq)
	fmt.Println("seq:", seq, ",max child seq is:", n)
}
