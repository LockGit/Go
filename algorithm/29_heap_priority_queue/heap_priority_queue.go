/**
 * Created by lock
 * 优先级队列
 */
package main

import "fmt"

type node struct {
	value    int
	priority int
}

type priorityQueue struct {
	heap     []node
	capacity int
	used     int
}

func newPriorityQueue(capacity int) *priorityQueue {
	return &priorityQueue{
		heap:     make([]node, capacity+1),
		capacity: capacity,
		used:     0,
	}
}

func (p *priorityQueue) swap(x, y int) {
	p.heap[x], p.heap[y] = p.heap[y], p.heap[x]
}

//堆化，调整堆
func (p *priorityQueue) adjustHeap(start, end int) {
	if start >= end {
		return
	}
	//保证优先级最高的节点在 h.heap[1] 的位置即可
	for i := end / 2; i >= start; i-- {
		high := start
		if start*2 <= end && p.heap[start].priority < p.heap[i*2].priority { //左子节点公式 2*start
			high = start * 2
		}
		if start*2+1 <= end && p.heap[high].priority < p.heap[start*2+1].priority { //右子节点2*start+1
			high = start*2 + 1
		}
		if i == high {
			break
		}
		p.swap(start, high)
		start = high
	}
}

func (p *priorityQueue) Push(node node) {
	if p.used >= p.capacity {
		return
	}
	p.used++
	p.heap[p.used] = node
	//p.adjustHeap(1, p.used)
}

func (p *priorityQueue) Pop() node {
	if p.used == 0 {
		return node{-1, -1}
	}
	//先堆化，在取出堆顶元素
	p.adjustHeap(1, p.used)
	node := p.heap[1]
	p.heap[1] = p.heap[p.used] //末尾元素移到堆顶
	p.used--
	p.adjustHeap(1, p.used) //取出后在堆化一次
	return node
}

//获取队列头部元素
func (p *priorityQueue) GetTop() (topNode node) {
	if p.used == 0 {
		return node{-1, -1}
	}
	p.adjustHeap(1, p.used)
	return p.heap[1]
}

func main() {
	pq := newPriorityQueue(10)
	for i := 1; i <= 6; i++ {
		nodeItem := node{
			value:    i,
			priority: i*8 - 5,
		}
		pq.Push(nodeItem)
	}
	fmt.Println("raw queue:", pq.heap)
	fmt.Println("after pop,queue list:", pq.heap, len(pq.heap))
	fmt.Println("get top:", pq.GetTop())
	pq.Push(node{
		value:    20,
		priority: 40,
	})
	fmt.Println("after push:", pq.heap)
	fmt.Println("pop item:", pq.Pop())
	fmt.Println("after pop item:", pq.heap, pq.used)
	fmt.Println(pq.Pop())
	fmt.Println(pq.Pop())
	fmt.Println(pq.Pop())
	fmt.Println(pq.Pop())
	fmt.Println(pq.Pop())
	fmt.Println(pq.Pop())
	fmt.Println(pq.Pop())
}
