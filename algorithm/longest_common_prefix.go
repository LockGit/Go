/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/30
 * Time: 10:47
 * 字符串最长公共前缀
 */
package main

import "fmt"

func main() {
	s := []string{"abc", "abcd", "abe"}
	res := s[0]
	for _, item := range s[1:] {
		n := len(item)
		for k := range res {
			if k >= n || res[k] != item[k] {
				res = res[:k]
			}
		}
	}
	fmt.Println(res)
}
