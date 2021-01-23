/**
 * Created by lock
找所有字母异位词
给定一个字符串s和一个非空字符串p,找到s中所有满足p的字母异位词的子串，返回这些子串的起始索引
说明：
字母异位词指字母相同，但排列不同的字符串。
不考虑答案输出的顺序。
示例：
输入：s="cbaedabacd",p="abc"
输出：[0,6]

解释：
字符串索引从0开始的子串，cba,是abc的字母异位词
字符串索引从6开始的子串，bac,是abc的字母异位词

相当于，输入一个串 S，一个串 T，找到 S 中所有 T 的排列，返回它们的起始索引。
*/
package main

import "fmt"

func findStrAnagrams(s, p string) (res []int) {
	window := make(map[string]int)
	target := make(map[string]int)
	for _, v := range p {
		target[string(v)]++
	}
	left, right := 0, 0
	validate := 0
	for right < len(s) {
		char := string(s[right])
		right++
		if v, ok := target[char]; ok && v >= 1 {
			window[char]++
			if window[char] == target[char] {
				validate++
			}
		}
		//收缩左窗口,此处 for right-left 可以用 if right-left == len(p)
		for right-left >= len(p) {
			if right-left == validate {
				res = append(res, left)
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
	return
}

func main() {
	s := "cbaedabacd"
	p := "abc"
	fmt.Println("find res:", findStrAnagrams(s, p))
}
