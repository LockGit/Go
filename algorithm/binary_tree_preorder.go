/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/24
 * Time: 13:41
 * 二叉树前序遍历
For example:
Given binary tree {1,#,2,3},
   1
    \
     2
    /
   3

return [1,2,3].
*/
package main

import "fmt"

type preOrderTree struct {
	data  int
	left  *preOrderTree
	right *preOrderTree
}

func getPreTreeData(root *preOrderTree) []int {
	var res []int
	stack := []*preOrderTree{root}
	for len(stack) > 0 {
		root = stack[0]
		stack = stack[1:]
		res = append(res, root.data)
		if root.left != nil {
			stack = append(stack, root.left)
		}
		if root.right != nil {
			stack = append(stack, root.right)
		}
	}
	return res
}

func main() {
	tree1 := new(preOrderTree)
	tree2 := new(preOrderTree)
	tree3 := new(preOrderTree)
	tree1.right = tree2
	tree2.left = tree3
	tree1.data = 1
	tree2.data = 2
	tree3.data = 3
	fmt.Println("获取二叉树前序遍历结果", getPreTreeData(tree1))
}
