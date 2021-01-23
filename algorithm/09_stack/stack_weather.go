/**
 * Created by lock
单调栈。Next Greater Number，你一个数组 T，这个数组存放的是近几天的天气气温，你返回一个等长的数组。
计算：对于每一天，你还要至少等多少天才能等到一个更暖和的气温；如果等不到那一天，填 0。
比如说给你输入 T = [73,74,75,71,69,76]，你返回 [1,1,3,2,1,0]。
*/
package main

import "fmt"

func stackWeather(arr []int) (ret []int) {
	var stack []int //栈中存的是索引
	result := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		for len(stack) > 0 && arr[i] >= arr[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			result[i] = 0
		} else {
			result[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	return result
}

func main() {
	arr := []int{73, 74, 75, 71, 69, 76}
	fmt.Println(arr, ",result is:", stackWeather(arr))
}
