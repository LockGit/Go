/**
 * Created by lock
原地去重
快慢指针技巧,当当前数据上操作，这个操作会破坏原始数组结构
*/
package main

import "fmt"

func originalUnique(arr []int) (res []int) {
	if len(arr) <= 1 {
		return arr
	}
	slow := 0
	fast := 0
	for i := 0; i <= len(arr)-1; i++ {
		if arr[slow] != arr[fast] {
			slow = slow + 1
			arr[slow] = arr[fast]
		}
		fast++
	}
	return arr[0 : slow+1]
}

func uniqueArr(arr []int) (res []int) {
	j := 0
	for i := 1; i < len(arr); i++ {
		if arr[j] == arr[i] {
			continue
		}
		j++
		arr[j] = arr[i]
	}
	return arr[0 : j+1]
}

func main() {
	arr := []int{0, 1, 1, 2, 2, 3, 4, 5, 6, 6, 7, 8}
	fmt.Println("uniq arr res is:", originalUnique(arr))

	arr2 := []int{0, 0, 1, 1, 2, 3, 4, 5, 5}
	fmt.Println("uniq arr res is:", uniqueArr(arr2))
}
