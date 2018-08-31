/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/31
 * Time: 14:44
 * 反转数组
	For example, with n = 7 and k = 3, the array [1,2,3,4,5,6,7] is rotated to
	[5,6,7,1,2,3,4].
*/

package main

import "fmt"

func rotate(nums []int, k int) []int {
	n := len(nums)
	k = k % n
	nums = reverseArr(nums, 0, n-1)
	nums = reverseArr(nums, 0, k-1)
	nums = reverseArr(nums, k, n-1)
	return nums
}

func reverseArr(nums []int, i int, j int) []int {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i = i + 1
		j = j - 1
	}
	return nums
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	k := 4
	fmt.Println(rotate(arr, k))
}
