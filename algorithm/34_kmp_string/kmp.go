/**
 * Created by lock
 * kmp算法实现
 */
package main

import "fmt"

func getNext(pattern string) []int {
	m := len(pattern)
	next := make([]int, m)
	for index := range next {
		next[index] = -1
	}
	for i := 1; i < m-1; i++ {
		j := next[i-1]
		for pattern[j+1] != pattern[i] && j >= 0 {
			j = next[j]
		}
		if pattern[j+1] == pattern[i] {
			j += 1
		}
		next[i] = j
	}

	return next
}

func findByKMP(s string, pattern string) int {
	n := len(s)
	m := len(pattern)
	if n < m {
		return -1
	}
	next := getNext(pattern)
	j := 0
	for i := 0; i < n; i++ {
		for j > 0 && s[i] != pattern[j] {
			j = next[j-1] + 1
		}

		if s[i] == pattern[j] {
			if j == m-1 {
				return i - m + 1
			}
			j += 1
		}
	}
	return -1
}

func main() {
	s := "abc abcdab abcdabcdabde"
	pattern := "bcdabd"
	fmt.Println(findByKMP(s, pattern)) //16

	s = "aabbbbaaabbababbabbbabaaabb"
	pattern = "abab"
	fmt.Println(findByKMP(s, pattern)) //11

	s = "aabbbbaaabbababbabbbabaaabb"
	pattern = "ababacd"
	fmt.Println(findByKMP(s, pattern)) //-1

	s = "hello"
	pattern = "ll"
	fmt.Println(findByKMP(s, pattern)) //2
}
