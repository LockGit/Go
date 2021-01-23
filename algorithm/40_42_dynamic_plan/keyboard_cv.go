/**
 * Created by lock
4键键盘粘贴复制问题,如何在 N 次敲击按钮后得到最多的 'a'？
A,在屏幕打印一个字符'a'
CA,选中整个屏幕
CC,复制选中的内容
CV,粘贴选中的内容到屏幕

示例：
输入：N=3
输出：3
按下3次'a'

输入：N=7
输出：9
解释：
我们最多可以在屏幕上显示9个'a',按键顺序：
A，A，A，CA，CC，CV，CV
*/
package main

import "fmt"

func maxCalc(item ...int) int {
	m := item[0]
	for _, v := range item {
		if v > m {
			m = v
		}
	}
	return m
}

//暴力求解
//可以增加备忘录
func forceCopy(n, count, copy int) (maxV int) {
	if n <= 0 {
		return count
	}
	return maxCalc(
		forceCopy(n-1, count+1, copy),    //A
		forceCopy(n-1, count+copy, copy), //CV 粘贴
		forceCopy(n-2, count, count),     //CA,CC //全选+复制 = 2次操作
	)
}

//动态规划
//最后一次按键要么是 A 要么是 C-V
//最优的操作序列一定是 C-A C-C 接着若干 C-V，用一个变量 j 作为若干 C-V 的起点。那么 j 之前的 2 个操作就应该是 C-A C-C 了：
func dpCopy(N int) (maxA int) {
	dp := make([]int, N+1)
	dp[0] = 0
	for i := 1; i <= N; i++ {
		// 按 A 键
		dp[i] = dp[i-1] + 1
		for j := 2; j < i; j++ {
			// 全选 & 复制 dp[j-2]，连续粘贴 i - j 次
			// 屏幕上共 dp[j - 2] * (i - j + 1) 个 A
			dp[i] = maxCalc(dp[i], dp[j-2]*(i-j+1))
		}
	}
	// N 次按键之后最多有几个 A？
	return dp[N]
}
func main() {
	N := 7
	fmt.Println("N=", N, ",can have max val is:", dpCopy(N))
	fmt.Println("N=", N, ",forceCopy can have max val is:", forceCopy(N, 0, 0))
}
