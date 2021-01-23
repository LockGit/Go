/**
 * Created by lock
改进实现计算器：
比如输入如下字符串，考虑小数问题：
3 + 9/3 ,calc result is: 6
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

type strListQueue2 struct {
	data []string
}

func (s *strListQueue2) pop() string {
	if len(s.data) > 0 {
		val := s.data[0]
		s.data = s.data[1:]
		return val
	}
	return ""
}

func checkIsSign(c string) (b bool) {
	arr := []string{"+", "-", "*", "/"}
	for _, item := range arr {
		if c == item {
			return true
		}
	}
	return false
}

//在go里，没有stack container,此处必须是引用类型，否则递归时作为值传递，此处弄个struct模拟stack
func helper2(strListObj *strListQueue2) (sum float64) {
	sign := "+"
	var stack []float64
	num := 0.0
	for len(strListObj.data) > 0 {
		c := strListObj.pop()
		if c == " " || c == "" {
			continue
		}
		if !checkIsSign(c) && c != "(" && c != ")" {
			num, _ = strconv.ParseFloat(c, 64)
		}
		if c == "(" {
			//如果遇到(，开始递归计算后面的部分，直到遇到)结束，括号内当做一个整体num计算,返回结果，此处传引用
			num = helper2(strListObj)
		}
		//len(strList)==0最后一位是数字的时候要加入stack
		if checkIsSign(c) || len(strListObj.data) == 0 {
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
	return sumStack2(stack)

}

func preDeal(strList []string) (okStrList []string) {
	var dealStrList []string
	for _, v := range strList {
		if v != " " {
			dealStrList = append(dealStrList, v)
		}
	}
	if len(dealStrList) <= 0 {
		return
	}
	addItem := ""
	for i := 0; i < len(dealStrList)-1; i++ {
		item := dealStrList[i]
		nextItem := dealStrList[i+1]
		if item == " " {
			continue
		}
		if item == "(" || item == ")" || checkIsSign(item) {
			okStrList = append(okStrList, item)
			continue
		}
		addItem += item
		if nextItem == "(" || nextItem == ")" || checkIsSign(nextItem) {
			okStrList = append(okStrList, addItem)
			addItem = ""
		}
	}
	okStrList = append(okStrList, dealStrList[len(dealStrList)-1])
	return
}

func stackCalcImprove2(calcStr string) (sum float64) {
	strList := strings.Split(calcStr, "")
	//pre deal
	strList = preDeal(strList)
	obj := &strListQueue2{
		data: strList,
	}
	return helper2(obj)
}

func sumStack2(stackRes []float64) (sum float64) {
	for _, item := range stackRes {
		sum += item
	}
	return sum
}

func main() {
	calcStr1 := "3 * (2-6 /(3 -7))"
	////calcStr1 := "3+(9/3)"
	fmt.Println(calcStr1, ",calc result is:", stackCalcImprove2(calcStr1))
	calcStr2 := "3+9/3"
	fmt.Println(calcStr2, ",calc result is:", stackCalcImprove2(calcStr2))
	calcStr3 := "3*2+3*3+22-2"
	fmt.Println(calcStr3, ",calc result is:", stackCalcImprove2(calcStr3))

	calcStr4 := "0.8/3 + 1.1/3 + 7.1/3"
	fmt.Println("direct calc:", 0.8/3+1.1/3+7.1/3)
	fmt.Println(calcStr4, ",calc result is:", stackCalcImprove2(calcStr4))
}
