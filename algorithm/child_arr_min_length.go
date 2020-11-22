/**
 * Created by lock
 * Date: 2020/11/4
 * 给定一个由 n 个正整数组成的数组和一个正整数 s，请找出该数组中满足其和 ≥ s 的最小长度子数组。如果无解，则返回 -1
example:
输入: [2,3,1,2,4,3], s = 7
输出: 2
解释: 子数组 [4,3] 是该条件下的最小长度子数组。

example:
输入: [1, 2, 3, 4, 5], s = 100
输出: -1

解题思路:
如果采用暴力解法，用变量 i 从左到右遍历数组，变量 j 从 i 到数组尾部遍历，将 i 到 j 内的元素遍历求和，找到和大于 s 的最短数组。时间复杂度为 O(n^3)。
对暴力法进行优化，使用累加器来进行求和，那么求和这步只需要 O(1)。总的时间复杂度为 O(n^2)。
使用二分法来继续优化，对于左端点 i，我们用二分法来寻找 j 。首先建立前缀和数组 sum，对于每个 i，在 i 到尾部这段区间上二分查找，找到满足 sum[j] - sum[i] > S 的最小的 j 。总的时间复杂度为 O(nlog(n))。
最优解法是采用滑窗。我们用 2 个指针，一个指向子数组开始的位置，一个指向数子组最后的位置，并维护区间内的和 curr_sum 大于等于 s 同时数组长度最小，实时更新最短的数组长度 res 。时间复杂度为 O(n)。

算法流程:
初始化左指针 left 和右指针 right 指向 0，子数组和 curr_sum 为 0 。
右指针 right 遍历 nums 数组，即不断移动滑窗右端
更新子数组的和，curr_sum += nums[right]
当子数组和满足条件，即 curr_cum >= s 时
更新 res = min(res, right - left + 1)，其中 right - left + 1 是当前子数组的长度
curr_sum -= nums[left]，然后左指针右移，继续判断当前数组和是否满足条件
返回 res

复杂度分析:
时间复杂度：O(n) 。每个指针移动都需要 O(n)的时间。每个元素至多被访问两次，一次被右端点访问，一次被左端点访问。
空间复杂度：O(1)。变量只需要常数空间。
*/
package main

import (
	"fmt"
)

func getMin(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}

func childArrMinLength(arr []int, s int) int64 {
	currentSum := 0
	left := 0
	max := int64(^uint(0) >> 1)
	res := max
	for right := 0; right < len(arr); right++ {
		currentSum = currentSum + arr[right]
		for currentSum >= s {
			res = getMin(res, int64(right-left+1))
			currentSum = currentSum - arr[left]
			left++
		}
	}
	if res == max {
		return -1
	}
	return res
}

func main() {
	arr := []int{2, 3, 1, 2, 4, 3}
	s := 7
	res := childArrMinLength(arr, s)
	fmt.Println(res)
}
