/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/23
 * Time: 15:06
 * 层序遍历二叉树
For example:
Given binary tree {3,9,20,#,#,15,7},
    3
   / \
  9  20
    /  \
   15   7
return its level order traversal as:
[
  [3],
  [9,20],
  [15,7]
]
*/
package main

import "fmt"

type BtreeNode struct {
	Data  int
	left  *BtreeNode
	right *BtreeNode
}

func layerOrder(node *BtreeNode) []int {
	if node != nil {
		stack := []*BtreeNode{node}
		var res []int
		for len(stack) > 0 {
			node = stack[0]
			stack = stack[1:]
			if node.left != nil {
				stack = append(stack, node.left)
			}
			if node.right != nil {
				stack = append(stack, node.right)
			}
			res = append(res, node.Data)
		}

		return res
	}
	return nil
}

//有层次的遍历,放到二维数组中,方法1
func layerOrderArr(root *BtreeNode) [][]int {
	if root != nil {
		queue := []*BtreeNode{root}
		queue = append(queue, nil)
		var level []int
		var res [][]int
		for len(queue) > 0 {
			root = queue[0]
			queue = queue[1:]
			if root == nil {
				res = append(res, level[:])
				level = nil
				if len(queue) > 0 {
					queue = append(queue, nil)
				}
			} else {
				level = append(level, root.Data)
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
	return nil
}

//有层次的遍历,放到二维数组中,方法2
func layerOrderArr2(root *BtreeNode) [][]int {
	if root != nil {
		queue := []*BtreeNode{root}
		var level []int
		var res [][]int
		for len(queue) > 0 {
			n := len(queue)
			level = nil
			for i := 0; i < n; i++ {
				root = queue[0]
				queue = queue[1:]
				if root.left != nil {
					queue = append(queue, root.left)
				}
				if root.right != nil {
					queue = append(queue, root.right)
				}
				level = append(level, root.Data)
			}
			res = append(res, level[:])
		}
		return res
	}
	return nil
}

func main() {
	tree1 := new(BtreeNode)
	tree1.Data = 3
	tree2 := new(BtreeNode)
	tree2.Data = 9
	tree3 := new(BtreeNode)
	tree3.Data = 20
	tree4 := new(BtreeNode)
	tree4.Data = 15
	tree5 := new(BtreeNode)
	tree5.Data = 7
	tree1.left = tree2
	tree1.right = tree3
	tree3.left = tree4
	tree3.right = tree5
	stack := layerOrder(tree1)
	fmt.Println(stack)
	res2 := layerOrderArr(tree1)
	fmt.Println(res2)
	res3 := layerOrderArr2(tree1)
	fmt.Println(res3)
}
