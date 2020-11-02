/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/27
 * Time: 16:17
 * 给定字符串T和S，求S的子串中有多少与T相等。
S = “rabbbit”, T = “rabbit”

Return 3.
*/
package main

import (
	"fmt"
)

func numberDistinct(s string, t string) int {
	if len(s) < len(t) {
		return 0
	}
	n := len(s)
	m := len(t)
	subArr := make([]int, m+1)
	subArr[0] = 1
	for i := 1; i < n+1; i++ {
		for k := 0; k < m; k++ {
			j := m - k
			if s[i-1] == t[j-1] {
				subArr[j] = subArr[j] + subArr[j-1]
			}
		}
	}
	return subArr[m]
}

func main() {
	s := "abcdefabcxyz"
	t := "abc"
	fmt.Println(numberDistinct(s, t))
}
