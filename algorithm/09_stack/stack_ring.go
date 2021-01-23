/**
 * Created by lock
单调栈。Next Greater Number，处理循环数组。
比如输入一个数组 [2,1,2,4,3]，你返回数组 [4,2,4,-1,4]。找下一个元素比当前元素大的元素。
拥有了环形属性，最后一个元素 3 绕了一圈后找到了比自己大的元素 4。
*/
package main

import "fmt"

//将数组翻倍,处理循环问题
func stackRing(arr []int) (ret []int) {
	var stack []int
	n := len(arr)
	result := make([]int, n)
	for i := 2*n - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] <= arr[i%n] {
			//出栈
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			result[i%n] = -1
		} else {
			//栈尾
			result[i%n] = stack[len(stack)-1]
		}
		stack = append(stack, arr[i%n])
	}
	return result
}

func main() {
	arr := []int{2, 1, 2, 4, 3}
	fmt.Println(arr, ",stack ring result is:", stackRing(arr))
}
