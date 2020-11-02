/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/20
 * Time: 23:22
 * 将大于10的数的各个位上的数字相加，若结果还大于0的话，则继续相加，直到数字小于10为止
 */
package main

import "fmt"

func main() {
	num := 38
	for (num / 10) > 0 {
		d := 0
		for num > 0 {
			d += num % 10
			num /= 10
			//fmt.Println(num, d)
		}
		num = d
	}
	fmt.Println(num)

	//方法2
	var num2 = 38
	res := (num2-1)%9 + 1
	fmt.Println(res)
}
