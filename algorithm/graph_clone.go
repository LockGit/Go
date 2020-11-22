/**
* Created by GoLand.
* User: lock
* Date: 2018/8/24
* Time: 17:37
* 无向图复制 dfs,bfs
   1
  / \
 /   \
0 --- 2
     / \
     \_/
*/
package main

import (
	"fmt"
)

type graph struct {
	label     int
	neighbors []*graph
}

func cloneGraph(node *graph) *graph {
	if node == nil {
		return nil
	}
	var queue []*graph
	var cloneNeighbors []*graph
	var cloneNeighbor *graph
	newNode := new(graph)
	newNode.label = node.label
	visited := make(map[*graph]*graph)
	dict := make(map[*graph]*graph)
	dict[node] = newNode
	queue = append(queue, node)
	for len(queue) > 0 {
		node = queue[0]
		queue = queue[1:]
		visited[node] = node
		cloneNode := dict[node]
		cloneNeighbors = nil
		for _, neighbor := range node.neighbors {
			if visited[neighbor] == nil {
				queue = append(queue, neighbor)
			}
			if dict[neighbor] == nil {
				cloneNeighbor = new(graph)
				cloneNeighbor.label = neighbor.label
				dict[neighbor] = cloneNeighbor
			} else {
				cloneNeighbor = dict[neighbor]
			}
			cloneNeighbors = append(cloneNeighbors, cloneNeighbor)
		}
		cloneNode.neighbors = cloneNeighbors
	}
	return newNode
}

func main() {
	node1 := new(graph)
	node2 := new(graph)
	node3 := new(graph)
	node1.label = 1
	node1.neighbors = []*graph{node2, node3}
	node2.label = 0
	node2.neighbors = []*graph{node1, node3}
	node3.label = 2
	node3.neighbors = []*graph{node1, node2}
	fmt.Println(node1.label, node1.neighbors)
	cloneNode := cloneGraph(node1)
	fmt.Println(cloneNode.label, cloneNode.neighbors)
}
