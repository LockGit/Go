/**
 * Created by lock
最长无重复子串,给定一个字符串，找出其中不含重复字符的最长子串长度
示例：
输入："abcabcbb"，返回：3
解释：无重复字符串abc的长度为3

示例：
输入："bbbb",返回：1
解释:无重复字符串为b,长度为1
*/
package main

import "fmt"

func findLongWithoutRepeatSubStr(s string) (length int) {
	window := make(map[string]int)
	left, right := 0, 0
	for right < len(s) {
		char := string(s[right])
		right++
		window[char]++
		for window[char] > 1 {
			removeChar := string(s[left])
			window[removeChar]--
			left++
		}
		if right-left > length {
			length = right - left
		}
	}
	return
}

func main() {
	s1 := "abcabcbb"
	fmt.Println("find:", s1, ",without repeat sub str is:", findLongWithoutRepeatSubStr(s1))
	s2 := "bbbb"
	fmt.Println("find:", s2, ",without repeat sub str is:", findLongWithoutRepeatSubStr(s2))
}
