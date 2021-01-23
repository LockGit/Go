/**
 * Created by lock
最小编辑距离
给定两个字符串s1和s2,计算出将s1转换成s2所使用的最少操作数
你可以对一个字符串进行如下操作：
1，插入一个字符
2，删除一个字符
3，替换一个字符
示例：s1="horse",s2="ros"
输出：3
解释：
h替换成r--->rorse
删除r--->rorse--->rose
删除e--->rose--->ros
*/
package main

import "fmt"

func minEditDistanceVal(item ...int) (minVal int) {
	minVal = item[0]
	for i := 1; i < len(item); i++ {
		if item[i] < minVal {
			minVal = item[i]
		}
	}
	return minVal
}

//双指针i,j指向s1,s2末尾，慢慢向前移动，知道i,j等于-1结束
func forceMinEditDistance(s1, s2 string, i, j int) (minTimes int) {
	if i == -1 {
		return j + 1
	}
	if j == -1 {
		return i + 1
	}
	if s1[i] == s2[j] {
		return forceMinEditDistance(s1, s2, i-1, j-1) //什么都不做
	} else {
		//插入，删除，替换 3选1
		return minEditDistanceVal(
			forceMinEditDistance(s1, s2, i, j-1)+1,   //插入
			forceMinEditDistance(s1, s2, i-1, j)+1,   //删除
			forceMinEditDistance(s1, s2, i-1, j-1)+1, //替换
		)
	}
}

var memoMinEdit = map[string]int{}

//增加备忘录优化
func memoMinEditDistance(s1, s2 string, i, j int) (minTimes int) {
	if i == -1 {
		return j + 1
	}
	if j == -1 {
		return i + 1
	}
	key := fmt.Sprintf("%d_%d", i, j)
	if val, ok := memoMinEdit[key]; ok {
		return val
	}
	if s1[i] == s2[j] {
		return forceMinEditDistance(s1, s2, i-1, j-1) //什么都不做
	} else {
		//插入，删除，替换 3选1
		minVal := minEditDistanceVal(
			forceMinEditDistance(s1, s2, i, j-1)+1,   //插入
			forceMinEditDistance(s1, s2, i-1, j)+1,   //删除
			forceMinEditDistance(s1, s2, i-1, j-1)+1, //替换
		)
		memoMinEdit[key] = minVal
		return memoMinEdit[key]
	}
}

//使用动态规划,自底向上
/*
		s1
s2 "" h o r s e
"" 0  1 2 3 4 5
r  1  1 3 2 3 4
o  2  3 1 3 3 4
s  3  4 2 2 2 3
*/
func minEditDistance(s1, s2 string) (minTimes int) {
	dp := make([][]int, len(s1)+1)
	for k, _ := range dp {
		item := make([]int, len(s2)+1)
		dp[k] = item
	}
	for k := 1; k <= len(s1); k++ {
		dp[k][0] = k
	}
	for k := 1; k <= len(s2); k++ {
		dp[0][k] = k
	}
	fmt.Println("init dp table:")
	for _, v := range dp {
		fmt.Println(v)
	}
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = minEditDistanceVal(
					dp[i][j-1]+1,   //插入
					dp[i-1][j]+1,   //删除
					dp[i-1][j-1]+1, //替换
				)
			}
		}
	}
	return dp[len(s1)][len(s2)]
}

func main() {
	s1 := "horse"
	s2 := "ros"
	fmt.Println(s1, "--->", s2, ", force calc min edit distance is:", forceMinEditDistance(s1, s2, len(s1)-1, len(s2)-1))
	fmt.Println(s1, "--->", s2, ", add memo calc min edit distance is:", memoMinEditDistance(s1, s2, len(s1)-1, len(s2)-1))
	fmt.Println(s1, "--->", s2, ", dp calc min edit distance is:", minEditDistance(s1, s2))
}
