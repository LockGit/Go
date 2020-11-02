/**
* Created by GoLand.
* User: lock
* Date: 2018/8/23
* Time: 22:55
* 在二叉树中找一条路径，使得该路径的和最大。该路径可以从二叉树任何结点开始，也可以到任何结点结束。
*
      1
     / \
    2   3
* return 6
*/
package main

import (
	"fmt"
)

type maxTreeSum struct {
	data  int
	left  *maxTreeSum
	right *maxTreeSum
}

//golang 的max函数只支持2个参数你说坑不坑? lock还需要自己实现一个
func max(maxVal int, args ...int) int {
	for _, v := range args {
		if maxVal < v {
			maxVal = v
		}
	}
	return maxVal
}

func calcTreeMaxSum(root *maxTreeSum) int {
	maxRes := []*maxTreeSum{new(maxTreeSum)}
	res := maxSum(root, maxRes)
	return max(res, maxRes[0].data)
}

func maxSum(root *maxTreeSum, maxRes []*maxTreeSum) int {
	if root == nil {
		return 0
	} else {
		leftMax := maxSum(root.left, maxRes)
		rightMax := maxSum(root.right, maxRes)
		rootMax := max(root.data, root.data+leftMax, root.data+rightMax)
		if maxRes[0] != nil {
			maxRes[0].data = max(maxRes[0].data, rootMax, root.data+leftMax+rightMax)
		} else {
			maxRes[0].data = max(rootMax, root.data+leftMax+rightMax)
		}
		return rootMax
	}
}

func main() {
	tree1 := new(maxTreeSum)
	tree2 := new(maxTreeSum)
	tree3 := new(maxTreeSum)
	tree1.data = 1
	tree2.data = 2
	tree3.data = 3
	tree1.left = tree2
	tree1.right = tree3
	fmt.Println("binary tree max path sum is", calcTreeMaxSum(tree1))
}
