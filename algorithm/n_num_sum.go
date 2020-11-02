/**
 * Created by lock
 * Date: 2020/10/30
 * 给一个数字n，找到n个数字，其和为0，要求数字不能重复。
 */
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getNumGroup(n, sum int) (res []int) {
	for k := 0; k < n-1; k++ {
		//要求数字不能重复的逻辑我就不写了，简单用随机因子代替下
		rand.Seed(time.Now().UnixNano() + int64(k))
		num := rand.Intn(1<<16) + 1
		res = append(res, num)
	}
	s := 0
	for _, v := range res {
		s = s + v
	}
	res = append(res, sum-s)
	return
}

func main() {
	//5 5个数字，要求数组不能重复
	n := 5
	targetSum := 0
	var result [][]int
	for x := 0; x < n; x++ {
		ret := getNumGroup(n, targetSum)
		result = append(result, ret)
		//要n组 就把break注释掉
		break
	}
	for _, item := range result {
		s := 0
		for _, child := range item {
			s = s + child
		}
		fmt.Println("方法2,target sum is:", targetSum, ",arr is:", item)
	}

	fmt.Println("-----方法2-----")
	rand.Seed(time.Now().UnixNano())
	baseNum := rand.Intn(1 << 8)
	constNum := 2
	var res2 []int
	for k := 0; k < n-1; k++ {
		res2 = append(res2, baseNum+k+constNum)
	}
	ss := 0
	for _, v := range res2 {
		ss = ss + v
	}
	res2 = append(res2, targetSum-ss)
	fmt.Println("方法2,target sum is,", targetSum, ",arr is:", res2)
}
