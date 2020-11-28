/**
 * Created by lock
	在一个非负整数a中，我们希望从中移除k个数字，让剩下的数字值最小，如何选择移除哪k个数字
	例如：3255155,删除4位,用topK最后结果是321,但实际最小是155，155最小那么需要删除3255
*/
package main

import (
	"fmt"
	"strings"
)

func removeTopKNum(num string, k int) string {
	stack := []byte{}
	for i := range num {
		digit := num[i]
		for k > 0 && len(stack) > 0 && digit < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, digit)
	}
	stack = stack[:len(stack)-k]
	ans := strings.TrimLeft(string(stack), "0")
	if ans == "" {
		ans = "0"
	}
	return ans
}

func main() {
	fmt.Println(removeTopKNum("4556847594546", 8))
	fmt.Println(removeTopKNum("3255155", 4))
}
