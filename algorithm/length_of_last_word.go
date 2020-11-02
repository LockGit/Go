/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/29
 * Time: 20:32
 * 求一个字符串最后一个字符的长度
	For example, Given s = "Hello World",
	return 5.
*/
package main

import "fmt"

func main() {
	s := "Hello World"
	i := len(s) - 1
	for i >= 0 && string(s[i]) == " " {
		i = i - 1
	}
	res := 0
	for i >= 0 && string(s[i]) != " " {
		res = res + 1
		i = i - 1
	}
	fmt.Println(res)
}
