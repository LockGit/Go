/**
 * Created by lock
信封问题，二维递增子序列
给定了一些标记了宽高的信封，宽度和高度（w,h)形式出现，当另一个信封的宽度和高度都比这个信封大的时候，
这个信封就可以放到另一个信封之中，求最多嵌套个数
说明：不允许旋转信封
示例：envelope=[[5,4],[6,4],[6,7],[2,3]]
输出：3
解释：最多信封的个数为3，组合为：[2,3]=>[5,4]=>[6,7]

//先对宽度w进行升序排序，如果遇到w相同的情况，则按照高度h降序排序。
//之后把所有的h作为一个数组，在这个数组上计算 LIS 的长度就是答案。
/*
1,8
2,3
5,4
5,2
6,7
6,4
然后在h上寻找最长递增子序列得到：3，4，7
*/
package main

import (
	"fmt"
	"sort"
)

func sortSeq(arr [][]int) (sortArr [][]int) {
	sort.Slice(arr, func(i, j int) bool {
		if arr[i][0] > arr[j][0] {
			return false
		}
		if arr[i][0] == arr[j][0] {
			if arr[i][1] > arr[j][1] {
				return true
			}
			return false
		}
		return true
	})
	return arr
}

func maxChild(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMaxIncSeqLength(arr [][]int) (maxLength int) {
	if len(arr) == 0 {
		return 0
	}
	//在二维数组h上寻找最长递增子序列即可
	var height []int
	for i := 0; i <= len(arr)-1; i++ {
		height = append(height, arr[i][1])
	}
	if len(height) == 0 {
		return 0
	}
	dp := make([]int, len(height))
	dp[0] = 1
	maxSeqLength := dp[0]
	for i := 1; i < len(dp); i++ {
		maxLen := dp[i-1]
		for k := 0; k <= i; k++ {
			if height[i] > height[k] {
				maxLen = dp[k] + 1
			}
		}
		dp[i] = maxLen
		if dp[i] > maxSeqLength {
			maxSeqLength = dp[i]
		}
	}
	return maxSeqLength
}

func main() {
	arr := [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}
	arr = sortSeq(arr)
	fmt.Println(arr)
	fmt.Println("get max inc seq length:", getMaxIncSeqLength(arr))
}
