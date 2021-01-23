/**
 * Created by lock
求一段字符串中的最长回文串
求最长回文串，还有个算法叫马拉车算法，可以将时间复杂度降到线性。
这个问题还可以用动态规划方法解决，时间复杂度一样，但是空间复杂度至少要 O(N^2) 来存储 DP table。
这道题是少有的动态规划非最优解法的问题！！！
*/
package main

import "fmt"

func checkIsPalindromeStr(str string, left int, right int) (childRes string) {
	for left >= 0 && right < len(str) {
		if str[left] != str[right] {
			break
		}
		left--
		right++
	}
	//取出回文子串，如果没有回文子串，这是空串返回
	return str[left+1 : right]
}

func getMaxChild(a, b string) (maxStr string) {
	if len(a) > len(b) {
		return a
	}
	return b
}

func longPalindromeChildStr(str string) (res string) {
	for i := 0; i < len(str); i++ {
		// 以 s[i] 为中心的最长回文子串
		childRes1 := checkIsPalindromeStr(str, i, i) //奇数情况
		// 以 s[i] 和 s[i+1] 为中心的最长回文子串
		childRes2 := checkIsPalindromeStr(str, i, i+1) //偶数情况
		resTmp := getMaxChild(childRes1, childRes2)
		if len(resTmp) > len(res) {
			res = resTmp
		}
	}
	return
}

func main() {
	s := "babad"
	fmt.Println(s, ",longPalindromeChildStr is:", longPalindromeChildStr(s))
	s2 := "cbbd"
	fmt.Println(s2, ",longPalindromeChildStr is:", longPalindromeChildStr(s2))
	s3 := "axmnabcdcbaxyz"
	fmt.Println(s3, ",longPalindromeChildStr is:", longPalindromeChildStr(s3))
}
