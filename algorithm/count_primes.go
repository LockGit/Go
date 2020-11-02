/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/27
 * Time: 00:02
 * 求小于给定非负数n的质数个数
 */
package main

import "fmt"

func countPrimes(n int) int {
	res := 0
	for i := 2; i < n; i++ {
		if isPrime(i) {
			res = res + 1
		}
	}
	return res
}

func isPrime(k int) bool {
	i := 2
	for i*i <= k {
		if k%i == 0 {
			return false
		}
		i = i + 1
	}
	return true
}

func main() {
	fmt.Println(countPrimes(500))
}
