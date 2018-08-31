/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/30
 * Time: 15:53
 * 海明重量---位操作
写下一个能够接受无符号整型的数字作为输入并返回这个数字共有多少位为’１‘的函数（也称作海明重量）。
举例，32位的数字’11‘用二进制表示为00000000000000000000000000001011，因此，编写的函数应该返回3。
*/
package main

import "fmt"

func main() {
	num := 11
	res := 0
	for num > 0 {
		if num&1 == 1 {
			res = res + 1
		}
		num = num >> 1
	}
	fmt.Println(res)
}
