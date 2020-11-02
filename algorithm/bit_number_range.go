/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/24
 * Time: 15:32
 * 等价于 求 m 与 n 二进制编码中 同为1的前缀. [5,7] , 5 & 6 & 7 = 4
Given a range [m, n] where 0 <= m <= n <= 2147483647, return the bitwise AND
of all numbers in this range, inclusive.

For example, given the range [5, 7], you should return 4.
*/
package main

import "fmt"

func rangeBitWise(m int, n int) int {
	var shift uint
	shift = 0
	for n > m {
		shift += 1
		m = m >> 1
		n = n >> 1
	}
	return m << shift
}

func main() {
	fmt.Println(rangeBitWise(5, 6))
	fmt.Println(rangeBitWise(5, 7))
	fmt.Println(rangeBitWise(0, 2147483647))
}
