/**
 * Created by lock
正则表达式
给定一个字符串s和一个模式串p,实现支持"."和"*"的正则匹配
"."匹配任意单个字符
"*"匹配0个或者多个前面的字符
示例：
输入：s="aa",p="a*"
返回：true

示例：
输入：s="aab"，p="c*a*b"
输出：true
解释：'c'可以出现0次，'a'可以被重复一次。因此可以匹配 "aab"
*/
package main

import "fmt"

var memoRegex = map[string]bool{}

func regexMatch(s *string, i int, p *string, j int) (b bool) {
	m := len(*s)
	n := len(*p)
	sVal, pVal := *s, *p
	// base case
	if j == n {
		return i == m
	}
	if i == m {
		if (n-j)%2 == 1 {
			return false
		}
		for ; j+1 < n; j += 2 {
			if string(pVal[j+1]) != "*" {
				return false
			}
		}
		return true
	}

	// 记录状态 (i, j)，消除重叠子问题
	key := fmt.Sprintf("%d_%d", i, j)
	if _, ok := memoRegex[key]; ok {
		return memoRegex[key]
	}

	res := false
	if sVal[i] == pVal[j] || string(pVal[j]) == "." {
		if j < n-1 && string(pVal[j+1]) == "*" {
			res = regexMatch(s, i, p, j+2) || regexMatch(s, i+1, p, j)
		} else {
			res = regexMatch(s, i+1, p, j+1)
		}
	} else {
		if j < n-1 && string(pVal[j+1]) == "*" {
			res = regexMatch(s, i, p, j+2)
		} else {
			res = false
		}
	}
	// 将当前结果记入备忘录
	memoRegex[key] = res
	return res
}

func main() {
	s := "aa"
	p := "a*"
	fmt.Println(s, ", pattern is:", p, ",res is:", regexMatch(&s, 0, &p, 0))
	s1 := "aab"
	p1 := "c*a*b"
	fmt.Println(s1, ", pattern is:", p1, ",res is:", regexMatch(&s1, 0, &p1, 0))
}
