/**
 * Created by lock
改进实现计算器：
比如输入如下字符串，不考虑小数问题：
3 + 9/3 ,calc result is: 6
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

var numListMap = map[string]int{
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

func isDigit(numStr string) (num int, ok bool) {
	if val, ok := numListMap[numStr]; ok {
		return val, true
	}
	return 0, false
}

type strListQueue struct {
	data []string
}

func (s *strListQueue) pop() string {
	if len(s.data) > 0 {
		val := s.data[0]
		s.data = s.data[1:]
		return val
	}
	return ""
}

//在go里，没有stack container,此处必须是引用类型，否则递归时作为值传递，此处弄个struct模拟stack
func helper(strListObj *strListQueue) (sum int) {
	//strList := strListObj.data
	sign := "+"
	var stack []int
	num := 0
	for len(strListObj.data) > 0 {
		c := strListObj.pop()
		if c == " " || c == "" {
			continue
		}
		//多个连续数组，需要组合到一起
		val, isNumber := isDigit(c)
		if isNumber == true {
			numStr := fmt.Sprintf("%d%d", num, val)
			num, _ = strconv.Atoi(numStr)
		}
		if c == "(" {
			//如果遇到(，开始递归计算后面的部分，直到遇到)结束，括号内当做一个整体num计算,返回结果，此处传引用
			num = helper(strListObj)
		}
		//len(strList)==0最后一位是数字的时候要加入stack
		if (isNumber == false && c != " ") || len(strListObj.data) == 0 {
			switch sign {
			case "+":
				stack = append(stack, num)
				break
			case "-":
				stack = append(stack, -num)
				break
			case "*":
				if len(stack) >= 1 {
					top := stack[len(stack)-1]
					stack[len(stack)-1] = top * num
				}
				break
			case "/":
				if len(stack) >= 1 {
					top := stack[len(stack)-1]
					stack[len(stack)-1] = top / num
				}
				break
			}
			num = 0
			sign = c
			// 遇到右括号返回递归结果
			if c == ")" {
				break
			}
		}
	}
	return sumStack(stack)
}

func stackCalcImprove(calcStr string) (sum int) {
	strList := strings.Split(calcStr, "")
	obj := &strListQueue{
		data: strList,
	}
	return helper(obj)
}

func sumStack(stackRes []int) (sum int) {
	for _, item := range stackRes {
		sum += item
	}
	return sum
}

func main() {
	calcStr1 := "3 * (2-6 /(3 -7))"
	////calcStr1 := "3+(9/3)"
	fmt.Println(calcStr1, ",calc result is:", stackCalcImprove(calcStr1))
	calcStr2 := "3+9/3"
	fmt.Println(calcStr2, ",calc result is:", stackCalcImprove(calcStr2))
	calcStr3 := "3*2+3*3+22-2"
	fmt.Println(calcStr3, ",calc result is:", stackCalcImprove(calcStr3))

	//calcStr4 := "0.8/3 + 1.1/3 + 7.1/3"
	//fmt.Println(calcStr4, ",calc result is:", stackCalcImprove(calcStr4))
}
