/**
 * Created by lock
动态规划：最小下降路径和
输入为一个n * n的二维数组matrix，请你计算从第一行落到最后一行，经过的路径和最小为多少。(假设所有元素都大于0)
每次下降，可以向下、向左下、向右下三个方向移动一格。
也就是说，可以从matrix[i][j]降到matrix[i+1][j]或matrix[i+1][j-1]或matrix[i+1][j+1]三个位置。
比如：
matrix =
[
	[2,1,3],
	[6,5,4],
	[7,8,9],
]
输出：13
解释：有两条路径：1，5，7 和 1，4，8
*/
package main

import (
	"fmt"
	"math"
)

var memo [][]int

func minVal(v ...int) (minVal int) {
	minVal = v[0]
	for _, item := range v {
		if item < minVal {
			minVal = item
		}
	}
	return minVal
}

func dpDownSum(matrix [][]int, n, j int) (v int) {
	if n < 0 || j < 0 || j >= len(matrix[0]) {
		return math.MaxInt64
	}
	if memo[n][j] > 0 {
		return memo[n][j]
	}
	if n == 0 {
		return matrix[n][j]
	}
	v = matrix[n][j] + minVal(
		dpDownSum(matrix, n-1, j),
		dpDownSum(matrix, n-1, j-1),
		dpDownSum(matrix, n-1, j+1),
	)
	memo[n][j] = v
	return
}

func minDownSum(matrix [][]int) (minSum int) {
	n := len(matrix) - 1
	memo = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		item := make([]int, n+1)
		memo[i] = item
	}
	minSum = math.MaxInt64
	for i := 0; i <= n; i++ {
		minSum = minVal(minSum, dpDownSum(matrix, n, i))
	}
	return minSum
}

func main() {
	matrix := [][]int{
		{2, 1, 3},
		{6, 5, 4},
		{7, 8, 9},
	}
	fmt.Println("min down sum is:", minDownSum(matrix))
}
