/**
 * Created by lock
给你一个字符串S，一个字符串T，请在字符串S里面找出：包含T所有字母的最小子串。
示例：
输入：S="ADOBECODEBANC",T="ABC"
输出："BANC"
*/
package main

import (
	"fmt"
	"math"
)

func checkInStr(needStr, targetStr string) (b bool) {
	trueTimes := 0
	for _, v := range targetStr {
		for _, vv := range needStr {
			if string(v) == string(vv) {
				trueTimes++
				break
			}
		}

	}
	if trueTimes == len(targetStr) {
		return true
	}
	return false
}

//暴力求解
func forceMinSubStr(s, t string) (minStr string) {
	minLen := len(s)
	//此处用: i<=len(s),golang的数组下标
	for i := 0; i <= len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if checkInStr(s[i:j], t) {
				if len(s[i:j]) < minLen {
					minLen = len(s[i:j])
					minStr = s[i:j]
				}
			}
		}
	}
	return
}

//滑动窗口获取最小子串
func getMinSubStrBySlideWindow(s, t string) (minSubStr string) {
	window := make(map[string]int)
	target := make(map[string]int)
	for _, v := range t {
		target[string(v)]++
	}
	left := 0
	right := 0
	validate := 0
	startIndex, minLength := 0, math.MaxInt64
	for right < len(s) {
		//右移窗口
		char := string(s[right])
		right++
		if v, ok := target[char]; ok && v >= 1 {
			window[char]++
			if window[char] == target[char] {
				validate++
			}
		}
		//收缩左窗口
		for validate == len(t) {
			if right-left < minLength {
				startIndex = left
				minLength = right - left
			}
			removeChar := string(s[left])
			left++
			if v, ok := target[removeChar]; ok && v >= 1 {
				if target[removeChar] == window[removeChar] {
					validate--
				}
				window[removeChar]--
			}
		}
	}
	if minLength == math.MaxInt64 {
		return ""
	}
	return s[startIndex : startIndex+minLength]
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(t[0:1])
	fmt.Println("force get min sub str:", forceMinSubStr(s, t))
	fmt.Println("slide window get min sub str:", getMinSubStrBySlideWindow(s, t))
}
