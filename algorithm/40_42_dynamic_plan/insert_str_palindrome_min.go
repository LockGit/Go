/**
 * Created by lock
最小插入次数构造回文串

案例：
输入 s = "abcea"，算法返回 2，因为可以给 s 插入 2 个字符变成回文串 "abeceba" 或者 "aebcbea"。
如果输入 s = "aba",则算法返回 0，因为 s 已经是回文串，不用插入任何字符。
*/
package main

import "fmt"

func minInsertTimes(a, b int) (minTimes int) {
	if a < b {
		return b
	}
	return a
}

func insertStrPalindromeMin(s string) (times int) {
	dp := make([][]int, len(s))
	for k, _ := range dp {
		item := make([]int, len(s))
		dp[k] = item
	}
	fmt.Println("init dp table is:", dp)
	//从下向上遍历
	for i := len(s) - 2; i >= 0; i-- {
		//从左向右递推
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp[i][j] = minInsertTimes(dp[i+1][j], dp[i][j-1]) + 1
			}
		}
	}
	fmt.Println("after dp table is:", dp)
	return dp[0][len(s)-1]
}

func main() {
	s1 := "abc"
	fmt.Println("insert str palindrome min times is:", insertStrPalindromeMin(s1))
	s2 := "abcdef"
	fmt.Println("insert str palindrome min times is:", insertStrPalindromeMin(s2))
}
