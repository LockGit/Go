/**
 * Created by lock
给定一个字符串，找出其中最长没有重复子串的长度。
给定字符串"abcabcbb"，最长不重复子串为"abc"，长度为3；
给定字符串"bbbbb"，最长不重复子串为"b"，长度为1；
给定字符串"pwwkew"，最长不重复子串为"wke"或"kew"，长度为3，注意子串必须要连续，"pwke"不符合要求。
*/
package main

import "fmt"

func isUniq(str string, i, j int) (b bool) {
	hashCheck := make(map[string]struct{})
	for index := i; index < j; index++ {
		if _, ok := hashCheck[string(str[index])]; ok {
			return false
		}
		hashCheck[string(str[index])] = struct{}{}
	}
	return true
}

func maxUniqLen(a ...int) (max int) {
	max = a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	return
}

//暴力求解
func getMaxUniqLenStr(str string) (maxLen int) {
	for i := 0; i < len(str); i++ {
		for j := i + 1; j <= len(str); j++ {
			if isUniq(str, i, j) {
				maxLen = maxUniqLen(j-i, maxLen)
			} else {
				break
			}
		}
	}
	return
}

//优化版本
func getMaxUniqLenStrImprove(str string) (maxLen int) {
	hashCount := make(map[string]int)
	start := 0
	for i := 0; i < len(str); i++ {
		if _, ok := hashCount[string(str[i])]; !ok {
			hashCount[string(str[i])] = i
		} else {
			maxLen = maxUniqLen(maxLen, i-hashCount[string(str[i])])
			hashCount[string(str[i])] = i
			start = i
		}
	}
	//最后部分check
	maxLen = maxUniqLen(maxLen, len(str)-start)
	return
}

func main() {
	s := "abcabcbb"
	fmt.Println("raw str is:", s, ",max len is:", getMaxUniqLenStr(s))
	fmt.Println("raw str is:", s, ",max len is:", getMaxUniqLenStrImprove(s))
}
