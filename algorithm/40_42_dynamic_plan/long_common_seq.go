/**
 * Created by lock
最长公共子序列
输入: str1 = "abcde", str2 = "ace"
输出: 3
解释: 最长公共子序列是 "ace"，它的长度是 3
*/
package main

import "fmt"

func maxLongCommonSql(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func forceFindFunc(s1, s2 string, i, j int) (length int) {
	if i == -1 || j == -1 {
		//base case
		return 0
	}
	if s1[i] == s2[j] {
		//找到一个 lcs 的元素，继续往前找
		return forceFindFunc(s1, s2, i-1, j-1) + 1
	} else {
		//谁能让最长公共子序列更长
		return maxLongCommonSql(forceFindFunc(s1, s2, i-1, j), forceFindFunc(s1, s2, i, j-1))
	}
}

//暴力搜索
func forceFindLongCommonSeq(s1, s2 string) (maxLen int) {
	return forceFindFunc(s1, s2, len(s1)-1, len(s2)-1)
}

//动态规划
func findLongCommonSeq(s1, s2 string) (maxLen int) {
	//dp table中第一个字符，用的空串，len(s)+1
	dp := make([][]int, len(s1)+1)
	//初始化为0
	for i := 0; i < len(dp); i++ {
		item := make([]int, len(s2)+1)
		dp[i] = item
	}
	fmt.Println("init dp table:", dp)
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				//找到一个 lcs 的元素，自底向上
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				//谁能让最长公共子序列更长
				dp[i][j] = maxLongCommonSql(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	fmt.Println("after dp is:", dp)
	return dp[len(s1)][len(s2)]
}

func main() {
	str1 := "abcde"
	str2 := "ace"
	fmt.Println("force find long common seq is", forceFindLongCommonSeq(str1, str2))
	fmt.Println("find long common seq is", findLongCommonSeq(str1, str2))
}
