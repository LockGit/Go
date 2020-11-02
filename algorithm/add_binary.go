/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/19
 * Time: 22:58
 * 给定两个二进制字符串，返回它们的和（也是一个二进制字符串）。 例如， a = "11" b = "1" 返回"100"。
 */
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	a := "1"
	b := "11"

	if len(a) < len(b) {
		a, b = b, a
	}
	n, m := len(a), len(b)

	c, r := 0, 0
	var res []string
	for k := 0; k < n; k++ {
		i := n - 1 - k
		if k < m {
			j := m - 1 - k
			x, err := strconv.Atoi(a[i : i+1])
			if err != nil {
				fmt.Println("字符串转整形出错了")
				break
			}
			y, err := strconv.Atoi(b[j : j+1])
			if err != nil {
				fmt.Println("字符串转整形出错了")
				break
			}
			r = (x + y + c) % 2
			c = (x + y + c) / 2
		} else {
			x, err := strconv.Atoi(a[i : i+1])
			if err != nil {
				fmt.Println("字符串转整形出错了")
				break
			}
			r = (x + c) % 2
			c = (x + c) / 2
		}
		//往头部插入一个元素,还要把整形转回字符串
		res = append([]string{strconv.Itoa(r)}, res...)
	}

	if c == 1 {
		res = append([]string{strconv.Itoa(c)}, res...)
	}

	fmt.Println(strings.Join(res, ""))
}
