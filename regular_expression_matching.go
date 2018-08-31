/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/30
 * Time: 16:40
 * 实现一个正则匹配算法
 */
package main

import "fmt"

func isMatch(s string, p string) bool {
	if p != "" && s == "" {
		return false
	}
	if s == "" {
		return false
	}
	return isMatchAux(s, p, 0, 0)
}

func isMatchAux(s string, p string, si int, pi int) bool {
	if pi == len(p) {
		return si == len(s)
	}
	if pi < len(p)-1 && string(p[pi+1]) != "*" || pi == len(p)-1 {
		isCurMatched := si < len(s) && (p[pi] == s[si] || string(p[pi]) == ".")
		isNextMatched := isMatchAux(s, p, si+1, pi+1)
		return isCurMatched && isNextMatched
	}
	for si < len(s) && pi < len(p) && (p[pi] == s[si] || string(p[pi]) == ".") {
		if isMatchAux(s, p, si, pi+2) {
			return true
		}
		si = si + 1
	}
	return isMatchAux(s, p, si, pi+2)
}

func main() {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "aa"))
	fmt.Println(isMatch("aaa", "aa"))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("aa", ".*"))
	fmt.Println(isMatch("ab", ".*"))
	fmt.Println(isMatch("aab", "c*a*b"))
}
