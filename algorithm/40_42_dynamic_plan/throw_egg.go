/**
 * Created by lock
高楼扔鸡蛋
你面前有一栋从 1 到 N 共 N 层的楼，然后给你 K 个鸡蛋（K 至少为 1）。
现在确定这栋楼存在楼层 0 <= F <= N，在这层楼将鸡蛋扔下去，鸡蛋恰好没摔碎（高于 F 的楼层都会碎，低于 F 的楼层都不会碎）。
现在问你，最坏情况下，你至少要扔几次鸡蛋，才能确定这个楼层 F 呢？

最坏情况至少要扔几次鸡蛋！！！
*/
package main

import (
	"fmt"
	"math"
)

func minEggVal(a, b int) (min int) {
	if a > b {
		return b
	}
	return a
}

func maxEggVal(a, b int) (min int) {
	if a > b {
		return a
	}
	return b
}

var memoEgg = map[string]int{}

//设鸡蛋的个数为k,楼的层数为n，最坏情况至少要扔几次鸡蛋！！！
func memoThrowEgg(k, n int) (times int) {
	if k == 1 {
		//一个鸡蛋，最坏情况线性探测
		return n
	}
	if n == 0 {
		return 0
	}
	key := fmt.Sprintf("%d_%d", k, n)
	if v, ok := memoEgg[key]; ok {
		return v
	}
	minRes := math.MaxInt64
	//穷举所有可能
	for i := 1; i < n+1; i++ {
		minRes = minEggVal(minRes,
			maxEggVal(
				memoThrowEgg(k, n-i),   //鸡蛋没碎,上面楼层测试
				memoThrowEgg(k-1, n-1), //鸡蛋碎了,下面楼层测试
			)+1, //试一次后需要+1
		)
	}
	memoEgg[key] = minRes
	return minRes
}

//设鸡蛋的个数为k,楼的层数为n，最坏情况至少要扔几次鸡蛋！！！，二分查找优化。
func memoThrowEggBinary(k, n int) (times int) {
	if k == 1 {
		//一个鸡蛋，最坏情况线性探测
		return n
	}
	if n == 0 {
		return 0
	}
	key := fmt.Sprintf("%d_%d", k, n)
	if v, ok := memoEgg[key]; ok {
		return v
	}
	minRes := math.MaxInt64
	low, height := 1, n
	for low <= height {
		mid := low + ((height - low) >> 1)
		broken := memoThrowEggBinary(k-1, n-1)      //碎
		notBroken := memoThrowEggBinary(k-1, n-mid) //没碎
		// minRes = min(max(碎，没碎) + 1)
		if broken > notBroken {
			height = mid - 1
			minRes = minEggVal(minRes, broken+1)
		} else {
			low = mid + 1
			minRes = minEggVal(minRes, notBroken+1)
		}
	}
	memoEgg[key] = minRes
	return minRes
}

//进阶，重新定义状态转移，优化
//设鸡蛋的个数为k,楼的层数为n，最坏情况至少要扔几次鸡蛋！！！
// 给你 k 个鸡蛋，测试 m 次，最坏情况下最多能测试 n 层楼。
//1、无论你在哪层楼扔鸡蛋，鸡蛋只可能摔碎或者没摔碎，碎了的话就测楼下，没碎的话就测楼上。
//2、无论你上楼还是下楼，总的楼层数 = 楼上的楼层数 + 楼下的楼层数 + 1（当前这层楼）。
//dp[k][m] = dp[k][m - 1] + dp[k - 1][m - 1] + 1
//dp[k][m - 1] 就是楼上的楼层数，因为鸡蛋个数 k 不变，也就是鸡蛋没碎，扔鸡蛋次数 m 减一；
//dp[k - 1][m - 1] 就是楼下的楼层数，因为鸡蛋个数 k 减一，也就是鸡蛋碎了，同时扔鸡蛋次数 m 减一。
//此处 m 是一个允许的次数上界，而不是扔了几次。最终要求的扔鸡蛋次数 m，m在状态之中而不是dp数组的结果！
func superThrowEgg(k, n int) (times int) {
	dp := make([][]int, k+1)
	for k, _ := range dp {
		item := make([]int, n+1)
		dp[k] = item
	}
	m := 0
	for dp[k][n] < n {
		m++
		for i := 1; i <= k; i++ {
			dp[i][m] = dp[i][m-1] + dp[i-1][m-1] + 1
		}
	}
	return m
}

func main() {

	fmt.Println(memoThrowEgg(1, 5))
	fmt.Println(memoThrowEgg(3, 9))
	fmt.Println(memoThrowEggBinary(1, 5))
	fmt.Println(memoThrowEggBinary(3, 9))
	fmt.Println(superThrowEgg(1, 5))
	fmt.Println(superThrowEgg(3, 9))
}
