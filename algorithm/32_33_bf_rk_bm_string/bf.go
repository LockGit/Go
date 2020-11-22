/**
 * Created by lock
 * bf算法实现
 */
package main

import "fmt"

//暴力搜素
func bfStrSearch(s, pattern string) int {
	sLen := len(s)
	patternLen := len(pattern)
	if sLen == 0 || patternLen == 0 || patternLen > sLen {
		return -1
	}
	findRange := sLen - patternLen
	for i := 0; i <= findRange; i++ {
		if (s[i : i+patternLen]) == pattern {
			return i
		}
	}
	return -1
}

func main() {
	s := "aaabcdefaacbcd"
	t := "acb"
	index := bfStrSearch(s, t)
	fmt.Println("bf find str index:", index)
}
