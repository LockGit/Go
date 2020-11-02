/**
 * Created by lock
 * Date: 2020/10/28
 * 给四个数字(0~9)，四个数字可以任意排列，排列的结果需要是一个合法的24小时制的时间，比如23:59，输出合法时间的数量。
 */
package main

import (
	"fmt"
)

/*


ii = [2, 3, 5, 7, 9, 1]
target = 13

def solution(score=0):
    for i in ii:
        if score + i == target:
            yield [i]
        elif score + i < target:
            for rs in solution(score + i):
                yield [i] + rs


for s in solution():
    print(s)

（Pass）
（Pass）
*/

func getNumTimeCombine(firstNum int, aTarget []int) (combineData []int) {
	if len(aTarget) == 4 {
		return aTarget
	}
	aTarget = append(aTarget, firstNum)
	return getNumTimeCombine(firstNum, aTarget)
}

func getNumTimeData(aNum []int) (res [][]int) {
	for _, v := range aNum {
		var aTarget []int
		data := getNumTimeCombine(v, aTarget)
		res = append(res, data)
	}
	return res
}

var dstData []int

func combineNum(index int, arr []int, res *[][]int) {
	if index == len(arr) {
		cpy := make([]int, len(arr))
		copy(cpy, arr)
		*res = append(*res, cpy)
		return
	}
	for i := index; i < len(arr); i++ {
		//交互位置
		if i != index {
			tmp := arr[index]
			arr[index] = arr[i]
			arr[i] = tmp
		}
		combineNum(index+1, arr, res)
		//回溯
		if i != index {
			tmp := arr[index]
			arr[index] = arr[i]
			arr[i] = tmp
		}
	}
	return
}

func main() {
	// 给四个数字(0~9)，四个数字可以任意排列，排列的结果需要是一个合法的24小时制的时间，比如23:59，输出合法时间的数量。
	// 先排列组合，在计算是否合法，也可以在排列组合的时候判断是否合法
	aNum := []int{1, 2, 3, 9}
	//hour 00-23
	//ms 00-59
	res := make([][]int, 0)
	combineNum(0, aNum, &res)
	fmt.Println("res", res)
	fmt.Println("----")
	for _, item := range res {
		a := item[0]
		b := item[1]
		c := item[2]
		d := item[3]
		if (a*10+b) <= 23 && (c*10+d) <= 59 {
			fmt.Println(fmt.Sprintf("%d%d:%d%d", a, b, c, d))
		}
	}
}
