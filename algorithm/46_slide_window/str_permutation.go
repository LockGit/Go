/**
 * Created by lock
给定两个字符串s1,s2,写一个函数判断s2是否包含s1的排列。
换句话说，第一个字符串的排列之一是第二个字符串的子串。
示例：
输入：s1="ab",s2="eidbaooo"
输出：true
解释：s2包含s1的排列之一（"ba")

示例：
输入：s1="ab",s2="eidboaoo"
输出:false

输入的 s1 是可以包含重复字符的,相当给你一个 S 和一个 T，请问你 S 中是否存在一个子串，包含 T 中所有字符且不包含其他字符
*/
package main

import "fmt"

//滑动窗口双指针
func checkStrPermutation(t, s string) (b bool) {
	target := make(map[string]int)
	window := make(map[string]int)
	for _, v := range t {
		target[string(v)]++
	}
	left, right, validate := 0, 0, 0
	for right < len(s)-1 {
		char := string(s[right])
		right++
		if v, ok := target[char]; ok && (v >= 1) {
			window[char]++
			if window[char] == target[char] {
				validate++
			}
		}
		//计算什么时候滑动窗口，这里的for 可以改成 if
		for right-left >= len(t) {
			if validate == len(t) {
				return true
			}
			removeChar := string(s[left])
			left++
			if v, ok := target[removeChar]; ok && v >= 1 {
				if window[removeChar] == target[removeChar] {
					validate--
				}
				window[removeChar]--
			}
		}
	}
	return false
}

func main() {
	t := "ab"
	s := "eidbaooo"
	fmt.Println("check :", checkStrPermutation(t, s))
	t = "ab"
	s = "eidboaoo"
	fmt.Println("check :", checkStrPermutation(t, s))
}
