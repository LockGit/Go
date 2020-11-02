/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/24
 * Time: 09:53
 * 返回所有根到叶节点的路径
	  1
	 /   \
	2     3
	 \
	  5
	All root-to-leaf paths are:

	["1->2->5", "1->3"]
*/
package main

import "fmt"

type PathTree struct {
	data  int
	left  *PathTree
	right *PathTree
}

func getTreePath(root *PathTree) [][]int {
	var res [][]int
	var cand []int
	recursiveTreePath(root, cand, res)
	return res
}

func recursiveTreePath(root *PathTree, cand []int, res [][]int) {
	if root == nil {
		return
	} else {
		cand = append(cand, root.data)
		if root.left == nil && root.right == nil {
			res = append(res, cand)
			fmt.Println(res)
		}
		recursiveTreePath(root.left, cand, res)
		recursiveTreePath(root.right, cand, res)
		cand = cand[:len(cand)-1]
	}
}

func main() {
	tree1 := new(PathTree)
	tree2 := new(PathTree)
	tree3 := new(PathTree)
	tree4 := new(PathTree)
	tree1.left = tree2
	tree1.right = tree3
	tree2.right = tree4
	tree1.data = 1
	tree2.data = 2
	tree3.data = 3
	tree4.data = 5
	fmt.Println("树的所有路径")
	getTreePath(tree1)
}
