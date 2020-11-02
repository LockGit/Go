/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/22
 * Time: 17:52
 * 找出字符串中所有的变位词,重复出现的
 * Input: "abab" output: "ab"
 */
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "ababcdef"
	dict := make(map[string]int)
	for _, v := range s {
		_, exist := dict[string(v)]
		if exist == false {
			dict[string(v)] = 1
		} else {
			dict[string(v)] = dict[string(v)] + 1
		}
	}

	var res []string
	for k, v := range dict {
		if v > 1 {
			res = append(res, k)
		}
	}
	fmt.Println(strings.Join(res, ""))
}
