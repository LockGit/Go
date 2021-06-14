/**
 * Created by lock
 * Date: 2021/3/12
n x n 的二维矩阵旋转90° （怎么转随便，一行看成一个整体比较方便旋转）
demo1,例如：
[
	[1,2],
	[3,4]
]
旋转得到：
[
	[3,1],
	[4,2]
]

demo2,例如：
1 2 3
4 5 6
7 8 9
旋转得到：
7 4 1
8 5 2
9 6 3
*/
package main

import "fmt"

func roateArr(rawArr [][]int) (newArr [][]int) {
	roate := make([][]int, len(rawArr))
	length := len(rawArr) - 1
	for i := 0; i <= length; i++ {
		for j := 0; j <= length; j++ {
			roate[j] = append([]int{rawArr[i][j]}, roate[j]...)
		}
	}
	return roate
}

func main() {
	rawArr := [][]int{
		[]int{1, 2},
		[]int{3, 4},
	}
	fmt.Println(roateArr(rawArr))
	fmt.Println("---------------------")
	rawArr2 := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	fmt.Println(roateArr(rawArr2))
}
