/**
 * Created by lock
单调队列
实现单调队列数据结构：deque，即双端队列。支持对头，对尾插入元素，返回对头，对尾。
输入：nums = [1,3,-1,-3,5,3,6,7],和 k=3
输出：[3,3,5,5,6,7]
窗口大小为3，在nums上滑动过去，每滑动一次，取当前窗口中的最大值
*/
package main

import "fmt"

type windowQueue struct {
	data []int
}

func (w *windowQueue) push(n int) {
	for len(w.data) > 0 && w.data[len(w.data)-1] < n {
		w.data = w.data[0 : len(w.data)-1]
	}
	w.data = append(w.data, n)
}

func (w *windowQueue) pop(n int) {
	if len(w.data) > 0 && w.data[0] == n {
		w.data = w.data[1:]
	}
}

func (w *windowQueue) max() (maxVal int) {
	maxVal = w.data[0]
	return maxVal
}

func windowMtQueue(nums []int, k int) (result []int) {
	window := new(windowQueue)
	for i := 0; i <= len(nums)-1; i++ {
		if i < k-1 {
			//先填满窗口的 (前 k-1)
			window.push(nums[i])
		} else {
			//窗口向前滑动
			window.push(nums[i])
			result = append(result, window.max())
			window.pop(nums[i-k+1])
		}
	}
	return result
}
func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println("nums is:", nums, ",window k=", k, ",result is:", windowMtQueue(nums, k))
}
