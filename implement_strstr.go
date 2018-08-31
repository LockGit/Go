/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/28
 * Time: 14:54
 * 实现字符串位置查找
 */
package main

import "fmt"

func strstr(haystack string, needle string) int {
	n := len(haystack)
	m := len(needle)
	for i := 0; i < n+1-m; i++ {
		matched := true
		for j := 0; j < m; j++ {
			if haystack[i+j] != needle[j] {
				matched = false
				break
			}
		}
		if matched {
			return i
		}
	}
	return -1
}

func strstr2(haystack string, needle string) int {
	n := len(haystack)
	m := len(needle)
	for i := 0; i < n+1-m; i++ {
		x := 1
		for j := 0; j < m; j++ {
			if haystack[i+j] != needle[j] {
				break
			}
			if x == m {
				return i
			}
			x++
		}
	}
	return -1
}

func main() {
	haystack := "abcdeflockxyz"
	needle := "lock"
	fmt.Println(strstr(haystack, needle))
	fmt.Println(strstr2(haystack, needle))
}
