/**
 * Created by lock
回溯算法 01背包
*/
package main

import "fmt"

var maxW = 0
var weight = []int{2, 2, 4, 6, 3} //物品重量
var n = 5                         //物品个数
var w = 9                         //背包承受的最大重量
var logWeight = [5][10]int{}      //备忘录

func pkgTraceBack(i int, cw int) {
	if cw == w || i == n {
		if cw > maxW {
			maxW = cw
			return
		}
	}
	if logWeight[i][cw] == 1 {
		return
	}
	logWeight[i][cw] = 1
	pkgTraceBack(i+1, cw)
	if cw+weight[i] <= w {
		pkgTraceBack(i+1, cw+weight[i])
	}
}

func main() {
	pkgTraceBack(5, 10)
	fmt.Println(maxW)
}
