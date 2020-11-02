/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/23
 * Time: 10:20
 * 动态规划
 * 用一个数组表示股票每天的价格，数组的第i个数表示股票在第i天的价格。 如果只允许进行一次交易，也就是说只允许买一支股票并卖掉，求最大的收益。
 * 动态规划法。从前向后遍历数组，记录当前出现过的最低价格，作为买入价格，并计算以当天价格出售的收益，作为可能的最大收益
 * 整个遍历过程中，出现过的最大收益就是所求
 */
package main

import (
	"fmt"
	"math"
)

func planOne(price []int) int {
	maxProfit := 0
	minPrice := price[0]
	for _, v := range price {
		maxProfit = int(math.Max(float64(maxProfit), float64(v-minPrice)))
		minPrice = int(math.Min(float64(minPrice), float64(v)))
	}
	return maxProfit
}

func planTwo(price []int) int {
	maxProfit := 0
	for i := 1; i < len(price); i++ {
		if price[i] > price[i-1] {
			maxProfit = maxProfit + (price[i] - price[i-1])
		}
	}
	return maxProfit
}

func planThree(price []int) int {
	length := len(price)
	if length <= 0 {
		return 0
	}
	a1, a2 := make([]int, length), make([]int, length)
	maxProfit1, maxProfit2 := 0, 0
	minPrice1, maxPrice2 := price[0], price[length-1]
	for i := 0; i < length; i++ {
		maxProfit1 = int(math.Max(float64(maxProfit1), float64(price[i]-minPrice1)))
		a1[i] = maxProfit1
		minPrice1 = int(math.Min(float64(minPrice1), float64(price[i])))
	}
	for i := 0; i < length; i++ {
		maxProfit2 = int(math.Max(float64(maxProfit2), float64(maxPrice2-price[length-1-i])))
		a2[length-1-i] = maxProfit2
		maxPrice2 = int(math.Max(float64(maxPrice2), float64(price[length-1-i])))
	}
	maxProfitRes := 0
	for i := 0; i < length; i++ {
		maxProfitRes = int(math.Max(float64(a1[i]+a2[i]), float64(maxProfitRes)))
	}
	return maxProfitRes
}

func main() {
	price := []int{6, 12, 2, 1, 8, 4, 9, 1, 6, 10}
	fmt.Println("单次最大收益", planOne(price))
	fmt.Println("既然能买卖任意次，那最大收益的方法就是尽可能多的低入高抛。只要明天比今天价格高，就应该今天买入明天再卖出,多次买入卖出后收益", planTwo(price))
	fmt.Println("最多交易两次,手上最多只能持有一支股票，最大收益", planThree(price))
}
