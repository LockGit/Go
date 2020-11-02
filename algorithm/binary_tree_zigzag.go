/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/24
 * Time: 14:42
 * 二叉树锯齿形层次遍历
Given binary tree {3,9,20,#,#,15,7},
    3
   / \
  9  20
    /  \
   15   7
return its zigzag level order traversal as:
[
  [3],
  [20,9],
  [15,7]
]
*/
package main

import "fmt"

type tree struct {
	data  int
	left  *tree
	right *tree
}

func getZigTreeData(root *tree) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	var queue []*tree
	rev := false
	var level []int
	queue = append(queue, root)
	queue = append(queue, nil)
	for len(queue) > 0 {
		root = queue[0]
		queue = queue[1:]
		if root == nil {
			if len(queue) > 0 {
				queue = append(queue, nil)
			}
			res = append(res, level)
			level = nil
			rev = true
		} else {
			if rev {
				level = append([]int{root.data}, level...)
			} else {
				level = append(level, root.data)
			}
			if root.left != nil {
				queue = append(queue, root.left)
			}
			if root.right != nil {
				queue = append(queue, root.right)
			}
		}
	}
	return res
}

func main() {
	tree1 := new(tree)
	tree2 := new(tree)
	tree3 := new(tree)
	tree4 := new(tree)
	tree5 := new(tree)
	tree1.left = tree2
	tree1.right = tree3
	tree3.left = tree4
	tree3.right = tree5
	tree1.data = 3
	tree2.data = 9
	tree3.data = 20
	tree4.data = 7
	tree5.data = 15
	fmt.Println(getZigTreeData(tree1))
}
