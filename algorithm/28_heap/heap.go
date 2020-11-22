/**
 * Created by lock
 * 堆
用数组存储堆：
如果数组第1个元素不使用，那么：
数组中下标为i的节点的左子节点就是下标为i∗2 的节点，右子节点就是下标为i∗2+1的节点，父节点就是下标为 i/2​ 的节点。

如果从数组第0个元素开始存储，那么：
如果节点的下标是i，那左子节点的下标就是i*2+1，右子节点的下标就是i*2+2，父节点的下标就是(i−1)/2​。

从0开始，计算左节点多个加法操作
*/
package main

import "fmt"

type heap struct {
	capacity int   //堆的大小
	arr      []int //用数组存储堆的数据,从下标1开始存储数据
	count    int   //堆中已经存储的数据个数
}

func newHeap(capacity int) *heap {
	return &heap{
		capacity: capacity,
		arr:      make([]int, capacity+1),
		count:    0,
	}
}

//交换下标为i和i/2的两个元素
func (h *heap) swap(x, y int) {
	h.arr[x], h.arr[y] = h.arr[y], h.arr[x]
}

func (h *heap) insert(num int) {
	if h.count >= h.capacity {
		return
	}
	h.count++
	h.arr[h.count] = num
	i := h.count
	for i/2 > 0 && h.arr[i] > h.arr[i/2] { // 自下往上堆化，最大值堆
		h.swap(i, i/2)
		i = i / 2
	}
}

//移除堆顶元素
func (h *heap) removeTop() {
	if h.count == 0 {
		return
	}
	h.arr[1] = h.arr[h.count]
	h.count--
	h.heapify(h.capacity, 1)
}

//堆化
func (h *heap) heapify(n, pos int) {
	// 自上往下堆化
	for {
		maxPos := pos
		if pos*2 <= n && h.arr[pos] < h.arr[pos*2] { //左子节点
			maxPos = pos * 2
		}
		if pos*2+1 <= n && h.arr[maxPos] < h.arr[pos*2+1] { //右子节点
			maxPos = pos*2 + 1
		}
		if maxPos == pos {
			break
		}
		h.swap(pos, maxPos)
		pos = maxPos
	}
}

//从下往上堆化
func (h *heap) buildHeap() {
	for i := h.capacity / 2; i >= 1; i-- {
		h.heapify(h.capacity, i)
	}
}

//堆排序
func (h *heap) sort() {
	k := h.capacity
	for k > 1 {
		h.swap(1, k) //放到末尾
		k--
		h.heapify(k, 1) //去除了堆头的剩下元素继续堆化
	}
}

func main() {
	h := newHeap(20)
	for i := 0; i < 30; i++ {
		h.insert(i)
	}
	fmt.Println(h.arr, len(h.arr))
	h.removeTop()
	h.removeTop()
	h.removeTop()
	fmt.Println("remove after:", h.arr)
	h.buildHeap()
	fmt.Println("continue build heap", h.arr)
	h.sort()
	fmt.Println("after sort:", h.arr)
}
