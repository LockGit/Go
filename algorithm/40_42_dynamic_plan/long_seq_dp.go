/**
 * Created by lock
动态规划，最长递增子序列
有一个数字序列包含 n 个不同的数字，如何求出这个序列中的最长递增子序列长度？比如 2, 9, 3, 6, 5, 1, 7
这样一组数字序列，它的最长递增子序列就是 2, 3, 5, 7，所以最长递增子序列的长度是 4。
*/
package main

import (
	"fmt"
	"math"
)

var maxDist = math.MinInt64
var maxSeq []int

// 回溯探索
func calcLongSeqBackTrace(a []int, i, j, aLength, dist int, seq []int) {
	if i == aLength {
		if dist > maxDist {
			maxDist = dist
			maxSeq = seq
		}
		return
	}

	if j == -1 || a[i] > a[j] {
		nSeq := append(seq, a[i])
		calcLongSeqBackTrace(a, i+1, i, aLength, dist+1, nSeq)
	}
	calcLongSeqBackTrace(a, i+1, j, aLength, dist, seq)
}

// 动态规划解法
func calcLongSeqDp(a []int, n int) int {
	if n < 0 || len(a) <= 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	// 递推数组
	maxLis := make([]int, n)
	// 选择路径记录
	optRecord := make([]int, n)

	maxLis[0] = 1
	optRecord[0] = -1

	for i := 1; i < n; i++ {
		v := 0
		pre := -1
		for k := 0; k < i; k++ {
			if a[i] > a[k] && maxLis[k] > v {
				v = maxLis[k]
				pre = k
			}
		}
		maxLis[i] = v + 1
		optRecord[i] = pre
	}

	var k = 0
	for i := 1; i < n; i++ {
		if maxLis[i] > maxLis[k] {
			k = i
		}
	}
	path := make([]int, 0, n)
	for t := k; t >= 0; {
		path = append(path, a[t])
		t = optRecord[t]
	}
	//for i := len(path) - 1; i >= 0; i-- {
	//	fmt.Print(path[i], " ")
	//}
	return maxLis[k]
}

func main() {
	arr := []int{2, 9, 3, 6, 5, 1, 7}
	var seq []int
	calcLongSeqBackTrace(arr, 0, -1, len(arr), 0, seq)
	fmt.Println("max child seq:", maxSeq)

	fmt.Println("use dp calc max seq length:", calcLongSeqDp(arr, len(arr)))
}
