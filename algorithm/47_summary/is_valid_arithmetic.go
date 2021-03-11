/**
* Created by lock
是否是合法的四则运算,假设不存在负数
*/
package main

import (
	"fmt"
	"strings"
)

var numMap = map[string]bool{
	"1": true,
	"2": true,
	"3": true,
	"4": true,
	"5": true,
	"6": true,
	"7": true,
	"8": true,
	"9": true,
}

var signMap = map[string]bool{
	"+": true,
	"-": true,
	"*": true,
	"/": true,
}

func isNum(s string) (b bool) {
	if _, ok := numMap[s]; ok {
		return true
	}
	return false
}

func isSign(s string) (b bool) {
	if _, ok := signMap[s]; ok {
		return true
	}
	return false
}

func isValidArithmetic(str string) (b bool) {
	lastIndex := len(str) - 1
	for i := 0; i < len(str)-1; i += 2 {
		if i+1 == lastIndex {
			if !isNum(string(str[i+1])) {
				return false
			}
		}
		if isNum(string(str[i])) && isSign(string(str[i+1])) {
			continue
		} else {
			return false
		}
	}
	return true
}

//使用双指针
func isValidArithmeticPointer(str string) (b bool) {
	right := len(str) - 1
	left := 0
	for left < right {
		if left%2 == 0 {
			if isNum(string(str[left])) && isNum(string(str[right])) {
				left++
				right--
			} else {
				return false
			}
		} else {
			if isSign(string(str[left])) && isSign(string(str[right])) {
				left++
				right--
			} else {
				return false
			}
		}
	}
	return true
}

func checkIsOk(q *QueueStr) (b bool) {
	for len(q.Data) > 0 {
		item := q.pop()
		if item == "(" {
			b = checkIsOk(q)
			if b == false {
				return false
			}
		}
		nextItem := q.pop()
		//if isNum(item) && nextItem == "" {
		//	continue
		//}
		if isNum(item) && isSign(nextItem) {
			b = true
		} else {
			b = false
			return
		}
		if item == ")" {
			break
		}
	}
	return true
}

type QueueStr struct {
	Data []string
}

func (q *QueueStr) pop() (v string) {
	if len(q.Data) > 0 {
		data := q.Data[0]
		q.Data = q.Data[1:]
		return data
	}
	return ""
}

func validHelper(str string) (b bool) {
	arr := strings.Split(str, "")
	q := &QueueStr{
		Data: arr,
	}
	return checkIsOk(q)
}

func main() {
	s := "3*((2+2)/2*2-2)+(2*2+3)+1+2"
	//s := "1+2-3"
	fmt.Println(s, ",validHelper:", validHelper(s))
	s = "3+2-2/2+2+2+2+3"
	fmt.Println(s, ",isValidArithmetic result is:", isValidArithmetic(s))
	fmt.Println(s, ",isValidArithmetic result is:", isValidArithmeticPointer(s))
	s2 := "3*((2+2/2*2-2)+(2*2+3)+1+2"
	fmt.Println(s2, ",isValidArithmetic result is:", isValidArithmetic(s2))
}
