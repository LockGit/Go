/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/22
 * Time: 18:52
 * 高度平衡二叉树判断
 * 如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。
 */
package main

import (
	"fmt"
	"math"
)

type Element interface{}

type Root struct {
	Data  Element
	left  *Root
	right *Root
}

func isBalanced(root *Root) bool {
	if root == nil {
		return true
	} else {
		if isBalanced(root.left) && isBalanced(root.right) {
			return math.Abs(float64(depth(root.left)-depth(root.right))) <= 1
		} else {
			return false
		}
	}
	return false
}

func depth(root *Root) int {
	if root == nil {
		return -1
	} else {
		return int(math.Max(float64(depth(root.left)), float64(depth(root.right)))) + 1
	}
}

func main() {
	node1 := new(Root)
	node2 := new(Root)
	node3 := new(Root)
	node4 := new(Root)
	node5 := new(Root)
	node1.left = node2
	node1.right = node3
	node2.left = node4
	node4.left = node5
	check := isBalanced(node1)
	fmt.Println(check)
}
