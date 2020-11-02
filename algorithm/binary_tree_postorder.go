/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/24
 * Time: 11:42
 * 二叉树的后序遍历
Given binary tree {1,#,2,3},
   1
    \
     2
    /
   3
return [3,2,1].
*/
package main

import "fmt"

type postOrderTree struct {
	data  int
	left  *postOrderTree
	right *postOrderTree
}

func reverseInts(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}

func getTreeData(root *postOrderTree) []int {
	var res []int
	stack := []*postOrderTree{root}
	for len(stack) > 0 {
		root = stack[0]
		stack = stack[1:]
		if root.left != nil {
			stack = append(stack, root.left)
		}
		if root.right != nil {
			stack = append(stack, root.right)
		}
		res = append(res, root.data)
	}
	return reverseInts(res)
}

func main() {
	tree1 := new(postOrderTree)
	tree2 := new(postOrderTree)
	tree3 := new(postOrderTree)
	tree1.right = tree2
	tree2.left = tree3
	tree1.data = 1
	tree2.data = 2
	tree3.data = 3
	fmt.Println("获取二叉树遍历结果", getTreeData(tree1))
}
