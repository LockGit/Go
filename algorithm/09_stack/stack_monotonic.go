/**
 * Created by lock
 * 单调栈。Next Greater Number，比如说，输入一个数组 nums = [2,1,2,4,3]，你返回数组 [4,2,4,-1,-1]。
解释：
第一个 2 后面比 2 大的数是 4;
1 后面比 1 大的数是 2；
第二个 2 后面比 2 大的数是 4;
4 后面没有比 4 大的数，填 -1；
3 后面没有比 3 大的数，填 -1。
*/
package main

import "fmt"

//nums := []int{2, 1, 2, 4, 3}
func stackMt(nums []int) (ret []int) {
	var stack []int
	result := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] <= nums[i] {
			//出栈
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			result[i] = -1
		} else {
			result[i] = stack[len(stack)-1]
		}
		//放入栈的最底部
		stack = append(stack, nums[i])
		fmt.Println("stack:", stack, nums[i])
	}
	return result
}

func main() {
	nums := []int{2, 1, 2, 4, 3}
	fmt.Println("使用单调栈计算：", stackMt(nums))
}
