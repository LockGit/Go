/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/28
 * Time: 15:35
 * 反转二叉树
Invert a binary tree.

     4
   /   \
  2     7
 / \   / \
1   3 6   9
to
     4
   /   \
  7     2
 / \   / \
9   6 3   1

*/
package main

import "fmt"

type InvertTree struct {
	data  int
	left  *InvertTree
	right *InvertTree
}

func invertTree(tree *InvertTree) *InvertTree {
	if tree == nil {
		return nil
	}
	left := invertTree(tree.left)
	right := invertTree(tree.right)
	tree.left = right
	tree.right = left
	return tree
}

func main() {
	tree := new(InvertTree)
	tree.data = 1
	//add some tree node
	fmt.Println(invertTree(tree))
}
