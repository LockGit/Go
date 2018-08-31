/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/28
 * Time: 15:50
 *  同构字符串
For example,

Given "egg", "add", return true.

Given "foo", "bar", return false.

Given "paper", "title", return true.
*/
package main

import (
	"fmt"
)

func checkStruct(s string, t string) bool {
	a := make(map[interface{}]interface{})
	for k, v := range s {
		if a[v] == nil {
			a[v] = t[k]
		} else {
			if a[v] != t[k] {
				return false
			}
		}
	}
	b := make(map[interface{}]interface{})
	for k, v := range t {
		if b[v] == nil {
			b[v] = s[k]
		} else {
			if b[v] != s[k] {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(checkStruct("agg", "emm"))
	fmt.Println(checkStruct("agg", "emn"))
	fmt.Println(checkStruct("foo", "bar"))
	fmt.Println(checkStruct("paper", "title"))
}
