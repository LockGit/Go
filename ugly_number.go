/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/31
 * Time: 16:57
 * 是否是丑陋数，所谓丑陋数就是其质数因子只能是2,3,5。
	那么最直接的办法就是不停的除以这些质数，如果剩余的数字是1的话就是丑陋数
*/
package main

import "fmt"

func isUglyNum(num int) bool {
	if num == 0 {
		return false
	}
	item := []int{2, 3, 5}
	for _, v := range item {
		for num%v == 0 {
			num = num / v
		}
	}
	if num == 1 {
		return true
	}
	return false
}

func main() {
	fmt.Println(isUglyNum(10))
	fmt.Println(isUglyNum(14))
	fmt.Println(isUglyNum(6))
}
