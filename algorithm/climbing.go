/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/24
 * Time: 16:05
 * 爬N阶的楼梯，每次可以爬1阶或者2阶，一共有多少种爬法? (递归，动态规划)
 */
package main

import "fmt"

//方法1，递归
func stepOne(num int) int {
	if num <= 2 {
		return num
	}
	return stepOne(num-1) + stepOne(num-2)
}

//方法2，动态规划
func stepTwo(num int) int {
	res := make([]int, num+1)
	res[0] = 1
	res[1] = 1
	for i := 2; i <= num; i++ {
		res[i] = res[i-1] + res[i-2]
	}
	return res[num]
}

//方法3，动态规划，但是要节省内存空间
func stepThree(num int) int {
	res := make([]int, 3)
	res[0] = 1
	res[1] = 1
	for i := 2; i <= num; i++ {
		res[i%3] = res[(i-1)%3] + res[(i-2)%3]
	}
	return res[num%3]
}

func main() {
	fmt.Println(stepOne(10))
	fmt.Println(stepTwo(10))
	fmt.Println(stepThree(10))
}
