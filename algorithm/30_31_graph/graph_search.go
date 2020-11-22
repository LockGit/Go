/**
* Created by lock
bfs dfs graph search
图的广度优先，深度优先搜索
*/
package main

import (
	"container/list"
	"fmt"
)

//无向图
type graph struct {
	linkedList []*list.List
	v          int
}

func newGraph(v int) *graph {
	g := &graph{
		linkedList: make([]*list.List, v),
		v:          v,
	}
	for i, _ := range g.linkedList {
		g.linkedList[i] = list.New()
	}
	return g
}

//insert edge，一条边存2次
func (g *graph) addEdge(s int, t int) {
	g.linkedList[s].PushBack(t)
	g.linkedList[t].PushBack(s)
}

//搜索一条从 s 到 t 的路径
func (g *graph) bfs(s int, t int) {
	if s == t {
		return
	}
	//init prev
	prev := make([]int, g.v)
	for index, _ := range prev {
		prev[index] = -1
	}

	//search by queue
	var queue []int
	visited := make([]bool, g.v)
	queue = append(queue, s)
	visited[s] = true
	isFound := false
	for len(queue) > 0 && !isFound {
		top := queue[0]
		queue = queue[1:]
		linkedList := g.linkedList[top]
		for e := linkedList.Front(); e != nil; e = e.Next() {
			k := e.Value.(int)
			if !visited[k] {
				prev[k] = top
				if k == t {
					isFound = true
					break
				}
				queue = append(queue, k)
				visited[k] = true
			}
		}
	}
	if isFound {
		printPrev(prev, s, t)
	} else {
		fmt.Printf("no path found from %d to %d\n", s, t)
	}
}

//递归打印s->t的路径,prev数组中存储的就是搜索的路径。不过，这个路径是反向存储的
func printPrev(prev []int, s int, t int) {
	if t == s || prev[t] == -1 {
		fmt.Printf("%d ", t)
		return
	} else {
		printPrev(prev, s, prev[t])
		fmt.Printf("%d ", t)
	}

}

//搜索一条从 s 到 t 的路径
func (g *graph) dfs(s int, t int) {
	prev := make([]int, g.v)
	for i := range prev {
		prev[i] = -1
	}
	visited := make([]bool, g.v)
	visited[s] = true
	isFound := false
	g.recurse(s, t, prev, visited, isFound)
	printPrev(prev, s, t)
}

func (g *graph) recurse(s int, t int, prev []int, visited []bool, isFound bool) {
	if isFound {
		return
	}
	visited[s] = true
	if s == t {
		isFound = true
		return
	}
	linkedList := g.linkedList[s]
	for e := linkedList.Front(); e != nil; e = e.Next() {
		k := e.Value.(int)
		if !visited[k] {
			prev[k] = s
			g.recurse(k, t, prev, visited, false)
		}
	}

}

func main() {
	graph := newGraph(8)
	graph.addEdge(0, 1)
	graph.addEdge(0, 3)
	graph.addEdge(1, 2)
	graph.addEdge(1, 4)
	graph.addEdge(2, 5)
	graph.addEdge(3, 4)
	graph.addEdge(4, 5)
	graph.addEdge(4, 6)
	graph.addEdge(5, 7)
	graph.addEdge(6, 7)
	//搜索一条从 s 到 t 的路径
	graph.bfs(0, 7)
	fmt.Println()
	graph.bfs(1, 3)
	fmt.Println()
	graph.dfs(0, 7)
	fmt.Println()
	graph.dfs(1, 3)
	fmt.Println()
}
