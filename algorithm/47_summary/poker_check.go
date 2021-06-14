/**
 * Created by lock
 * Date: 2021/3/12
斗地主中，大小连续的牌可以作为顺子，有时候我们把对子拆掉，结合单牌，可以组合出更多的顺子，可能更容易赢。
给你输入一个升序排列的数组nums（可能包含重复数字），请你判断nums是否能够被分割成若干个长度至少为 3 的子序列，每个子序列都由连续的整数组成。
比如题目举的例子，输入nums = [1,2,3,3,4,4,5,5]，算法返回 true。
因为nums可以被分割成[1,2,3,4,5]和[3,4,5]两个包含连续整数子序列。
但如果输入nums = [1,2,3,4,4,5]，算法返回 false，因为无法分割成两个长度至少为 3 的连续子序列。
*/
package main

import "fmt"

func pokerCheck(arr []int) (b bool) {
	freq := make(map[int]int)
	need := make(map[int]int)
	for _, v := range arr {
		freq[v]++
	}
	for _, v := range arr {
		if freq[v] == 0 {
			// 已经被用到其他子序列中
			continue
		}
		if need[v] > 0 {
			//判断v能否接到其他子序列后面
			freq[v]--
			need[v]--
			need[v+1]++
		} else if freq[v] > 0 && freq[v+1] > 0 && freq[v+2] > 0 {
			// 将v作为开头，新建一个长度为 3 的子序列 [v,v+1,v+2]
			freq[v]--
			freq[v+1]--
			freq[v+2]--
			// 对v+3的需求加一
			need[v+3]++
		} else {
			return false
		}
	}
	return true
}

func main() {
	arr := []int{1, 2, 3, 3, 4, 4, 5, 5}
	fmt.Println("check arr:", arr, ",res is:", pokerCheck(arr))
	arr1 := []int{1, 2, 3, 4, 4, 5}
	fmt.Println("check arr:", arr1, ",res is:", pokerCheck(arr1))
}
