/**
 * Created by lock
最长回文子序列（这题还可以结合马拉车算法）
示例：
输入："bbbb"
输出：4
一个可能最长的回文自序列为"bbbb"
*/
package main

import "fmt"

func maxChildSeq(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findLongPalindromeChildSeq(s string) (maxLen int) {
	dp := make([][]int, len(s))
	for k, _ := range dp {
		item := make([]int, len(s))
		dp[k] = item
	}
	for index := 0; index < len(dp); index++ {
		dp[index][index] = 1
	}
	//倒着遍历
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = maxChildSeq(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][len(s)-1]
}

func main() {
	s := "bbbb"
	fmt.Println("long palindrome child seq length:", findLongPalindromeChildSeq(s))
	s2 := "aba"
	fmt.Println("long palindrome child seq length:", findLongPalindromeChildSeq(s2))
}
