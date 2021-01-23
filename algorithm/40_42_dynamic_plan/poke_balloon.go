/**
 * Created by lock
戳气球问题
给你一组气球[0,n-1]，每个气球的分值各不相同，戳破一个气球的得分为：score[now-1]*score[now]score[now+1]

假设这组气球的最左边与最右边分值均为1，nums[-1]==nums[n]=1
示例：
输入: [3 1 5 8]
输出: 3*1*5 + 3*5*8 + 1*3*8 + 1*8*1 = 167
*/
package main

import (
	"fmt"
	"math"
)

var maxScore = math.MinInt64

func removeElement(arr []int, i int) (removeArr []int) {
	for k := 0; k < len(arr); k++ {
		if i != k {
			removeArr = append(removeArr, arr[k])
		}
	}
	return
}

//使用回溯法暴力求解计算
func forcePokeBalloon(arr []int, score int) (mScore int) {
	if len(arr) == 0 {
		if score > maxScore {
			maxScore = score
		}
		return maxScore
	}
	for i := 0; i < len(arr); i++ {
		leftScore := 1
		if i-1 >= 0 {
			leftScore = arr[i-1]
		}
		rightScore := 1
		if i+1 <= len(arr)-1 {
			rightScore = arr[i+1]
		}
		nowScore := leftScore * arr[i] * rightScore
		//戳破其中一个气球
		newArr := removeElement(arr, i)
		//回溯
		forcePokeBalloon(newArr, score+nowScore)
		//撤销操作，这里是slice类型,不用做相关操作了
	}
	return maxScore
}

func maxPokeScore(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//使用动态规划
//戳破第k个气球，使得dp[i][j]值最大
//最值在 dp[0][n+1]
func dpPokeBalloon(arr []int) (maxScore int) {
	n := len(arr)
	//添加虚拟气球，left,right=1的分配一个空间，初始化为1
	points := make([]int, n+2)
	points[0] = 1
	points[n+1] = 1
	for x := 1; x <= n; x++ {
		points[x] = arr[x-1]
	}
	dp := make([][]int, n+2)
	for y, _ := range dp {
		item := make([]int, n+2)
		dp[y] = item
	}
	//从下向上，从左到右遍历
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n+2; j++ {
			for k := i + 1; k < j; k++ {
				//对于一组给定的 i 和 j，穷举 i < k < j 的所有气球 k，选择得分最高的作为 dp[i][j] 的值即可，这也就是状态转移方程：
				dp[i][j] = maxPokeScore(
					dp[i][j],
					dp[i][k]+points[i]*points[j]*points[k]+dp[k][j],
				)
			}
		}
	}
	fmt.Println("---- dp table is:")
	for _, v := range dp {
		fmt.Println(v)
	}
	fmt.Println("---- ")
	return dp[0][n+1]
}

func main() {
	arr := []int{3, 1, 5, 8}
	fmt.Println(arr, ",use force backtrace calc val is:", forcePokeBalloon(arr, 0))
	fmt.Println(arr, ",use dp calc val is:", dpPokeBalloon(arr))
}
