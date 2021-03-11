/**
 * Created by lock
找出aab类型的字符串
*/
package main

import "fmt"

func findAabStr(s string) (res [][]string, newStr string) {
	endIndex := 0
	c := 0
	for i := 0; i <= len(s)-3; i++ {
		now := string(s[i])
		next := string(s[i+1])
		end := string(s[i+2])
		if now == next && end != now {
			if c == 0 {
				newStr = newStr + s[0:i]
			}
			if c > 0 {
				newStr = newStr + s[endIndex:i]
			}
			endIndex = i + 3
			c++
			res = append(res, []string{now, next, end})
		}
	}
	newStr = newStr + s[endIndex:]
	return res, newStr
}

func main() {
	s := "xaabxyzmmnxxzfkkm"
	res, newStr := findAabStr(s)
	fmt.Println(s, ",find aab str:", res, newStr)

	s2 := "xxabxz"
	res2, newStr2 := findAabStr(s2)
	fmt.Println(s, ",find aab str:", res2, newStr2)
}
