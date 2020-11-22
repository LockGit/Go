/**
 * Created by lock
 * 顺序栈，用数组实现的栈，演算一个四则运算过程 3+2-1*2+6/2
 */
package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type stackArr struct {
	Items []interface{} //数组
	Size  int           //栈的大小
	Count int           //栈中元素个数
}

//压栈
func (s *stackArr) push(item interface{}) bool {
	if s.Size == s.Count {
		return false
	}
	s.Items = append(s.Items, item)
	s.Count++
	return true
}

//出栈
func (s *stackArr) pop() (item interface{}, err error) {
	if s.Count == 0 {
		return nil, errors.New("no item")
	}
	item = s.Items[s.Count-1]
	s.Count = s.Count - 1
	return
}

var opLevel = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
}

func compareOp(opTop string, nowOp string) bool {
	if opLevel[nowOp] > opLevel[opTop] {
		//直接入栈场景
		return true
	} else if opLevel[nowOp] <= opLevel[opTop] {
		//操作
		return false
	}
	return false
}

func calcStack(numStack []int, opStack []string) (newNumStack []int, newOpStack []string) {
	for len(opStack) > 0 && len(numStack) >= 2 {
		opTop := opStack[len(opStack)-1]
		opNum1 := numStack[len(numStack)-2]
		opNum2 := numStack[len(numStack)-1]
		fmt.Println(opNum1, opTop, opNum2)
		var res int
		switch opTop {
		case "+":
			res = opNum1 + opNum2
		case "-":
			res = opNum1 - opNum2
		case "*":
			res = opNum1 * opNum2
		case "/":
			res = opNum1 / opNum2
		}
		numStack = numStack[0 : len(numStack)-2]
		numStack = append(numStack, res)
		opStack = opStack[0 : len(opStack)-1]
	}
	newNumStack = numStack
	newOpStack = opStack
	return
}

//演算一个四则运算过程
func stackCalcNumStr(str string) (res int) {
	//str := "3+2-1*2+6/2"
	var numStack []int
	var opStack []string
	//当遇到运算符，就与运算符栈的栈顶元素进行比较。
	for _, v := range str {
		if strings.Contains("+-*/", string(v)) {
			if len(opStack) > 0 {
				top := opStack[len(opStack)-1]
				nowOp := string(v)
				if compareOp(top, nowOp) {
					//如果比运算符栈顶元素的优先级高，就将当前运算符压入栈,
					opStack = append(opStack, string(v))
				} else {
					//如果比运算符栈顶元素的优先级低或者相同，从运算符栈中取栈顶运算符，从操作数栈的栈顶取 2 个操作数，然后进行计算
					numStack, opStack = calcStack(numStack, opStack)
					opStack = append(opStack, string(v))
				}
			} else {
				opStack = append(opStack, string(v))
				fmt.Println("opStack item", opStack)
			}
		} else {
			num, _ := strconv.Atoi(string(v))
			numStack = append(numStack, num)
			fmt.Println("numStack item", numStack)
		}
	}
	numStack, opStack = calcStack(numStack, opStack)
	fmt.Println("------- calc result: -------")
	fmt.Println("numStack", numStack)
	fmt.Println("opStack", opStack)
	res = numStack[0]
	return
}

//tips:demo,非线程安全
func main() {
	stackSize := 100
	s := &stackArr{
		Items: make([]interface{}, stackSize),
		Size:  stackSize,
		Count: 0,
	}
	for i := 0; i < 8; i++ {
		ret := s.push(i)
		fmt.Println("push:", i, ",result:", ret)
	}
	for i := 0; i < 8; i++ {
		if ret, err := s.pop(); err == nil {
			fmt.Println("pop:", i, ",result:", ret)
		}
	}
	//-------
	fmt.Println("演算一个四则运算过程 3+2-1*2+6/2")
	str := "3+2-1*2+6/2"
	res := stackCalcNumStr(str)
	fmt.Println("calc res:", res)
}
