/**
 * Created by lock
 * Date: 2021/3/12
字符串解密：
a2bc3d1 ===> aabcbcbcd
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getRepeat(n int, str string) string {
	return strings.Repeat(str, n)
}
func strDecrypt(raw string) (decryptStr string) {
	s := ""
	for _, v := range raw {
		if v >= 48 && v <= 57 {
			n, _ := strconv.Atoi(string(v))
			decryptStr += getRepeat(n, s)
			s = ""
		} else {
			s += string(v)
		}
	}
	return
}

func main() {
	fmt.Println(strDecrypt("a2bc3d1"))
}
