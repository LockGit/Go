/**
 * Created by lock
 * 希尔排序，不稳定排序
 */
package main

import "fmt"

func shellSort(nums []int) []int {
	//外层步长控制
	for step := len(nums) / 2; step > 0; step /= 2 {
		//开始插入排序
		for i := step; i < len(nums); i++ {
			//满足条件则插入
			for j := i - step; j >= 0 && nums[j+step] < nums[j]; j -= step {
				nums[j], nums[j+step] = nums[j+step], nums[j]
			}
		}
	}
	return nums
}

func main() {
	nums := []int{9, 8, 2, 1, 4, 5, 8, 3}
	shellSort(nums)
	fmt.Println("shell sort:", nums)
}
