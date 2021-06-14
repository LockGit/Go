/**
 * Created by lock
 * 完全二叉树check
深度为 k 的，有n个结点的二叉树，当且仅当其每个结点都与深度为 k 的满二叉树中编号从 1 至 n 的结点一一对应，称之为完全二叉树。
（除最后一层外，每一层上的节点数均达到最大值；在最后一层上只缺少右边的若干结点）
*/
package main

import "fmt"

var ids map[int]bool

type completeTreeNode struct {
	left  *completeTreeNode
	right *completeTreeNode
}

func isCompleteTree(root *completeTreeNode) bool {
	ids = make(map[int]bool)
	dfs(root, 0)
	for i := 0; i < len(ids); i++ {
		//最后一层尽可能向左靠近 [1,2,3,4,5,null,7]
		if !ids[i] {
			return false
		}
	}
	return true
}

func dfs(node *completeTreeNode, id int) {
	if node == nil {
		return
	}
	ids[id] = true
	id = id << 1 // x2
	dfs(node.left, id+1)
	dfs(node.right, id+2)
}
func main() {
	tree := &completeTreeNode{
		left:  &completeTreeNode{},
		right: nil,
	}
	fmt.Println("isCompleteTree:", isCompleteTree(tree))
}
