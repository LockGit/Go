/**
 * Created by lock
烧饼排序
假设盘子上有 n 块面积大小不一的烧饼，你如何用一把锅铲进行若干次翻转，让这些烧饼的大小有序（小的在上，大的在下）
求具体的反转操作序列
*/
package main

import "fmt"

var cakeRes []int

//数组索引从 0 开始，而我们要返回的结果是从 1 开始算
func cakeSort(arr []int, n int) {
	if n == 1 {
		return
	}
	// 寻找最大饼的索引
	maxCake := 0
	maxCakeIndex := 0
	for i := 0; i < n; i++ {
		if arr[i] > maxCake {
			maxCakeIndex = i
			maxCake = arr[i]
		}
	}
	// 第一次翻转，将最大饼翻到最上面
	reverseCake(arr, 0, maxCakeIndex)
	cakeRes = append(cakeRes, maxCakeIndex+1)
	//第二次翻转，将最大饼翻到最下面
	reverseCake(arr, 0, n-1)
	cakeRes = append(cakeRes, n)

	//递归
	cakeSort(arr, n-1)
}

func reverseCake(arr []int, i, j int) {
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

func main() {
	arr := []int{3, 2, 4, 1}
	cakeSort(arr, len(arr))
	fmt.Println("sort cake nums a per arr:", cakeRes)
	fmt.Println("after sort", arr)
}
