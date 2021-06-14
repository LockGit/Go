/**
 * Created by lock
平衡二叉树check
1，可以是空树。
2，假如不是空树，任何一个结点的左子树与右子树都是平衡二叉树，并且高度之差的绝对值不超过 1。
*/
package main

import "fmt"

type treeAvlNode struct {
	left  *treeAvlNode
	right *treeAvlNode
	value int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
func getHeight(root *treeAvlNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(getHeight(root.left), getHeight(root.right))
}
func isBalanced(root *treeAvlNode) bool {
	if root == nil {
		return true
	}
	if abs(getHeight(root.left)-getHeight(root.right)) > 1 {
		return false
	}
	return isBalanced(root.left) && isBalanced(root.right)
}

func main() {
	tree := &treeAvlNode{
		left:  &treeAvlNode{},
		right: nil,
	}
	fmt.Println("isBalanced tree:", isBalanced(tree))
}
