/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/31
 * Time: 16:40
 * 两数之和
Example:
Given nums = [2, 7, 11, 15], target = 9,
Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/
package main

import "fmt"

func calcTwoNumSum(arr []int, target int) (interface{}, interface{}) {
	d := make(map[int]interface{})
	for k, v := range arr {
		if d[v] != nil {
			return d[v].(int), k
		}
		d[target-v] = k
	}
	return nil, nil
}

func main() {
	arr := []int{2, 3, 8, 7, 15}
	target := 9
	a, b := calcTwoNumSum(arr, target)
	fmt.Println(a, b)
}
