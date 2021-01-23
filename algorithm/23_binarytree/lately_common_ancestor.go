/**
 * Created by lock
二叉树的最近公共祖先
*/
package main

import "fmt"

type commonTree struct {
	data  interface{}
	left  *commonTree
	right *commonTree
}

func detectCommonAncestor(root *commonTree, p, q *commonTree) (ct *commonTree) {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	left := detectCommonAncestor(root.left, p, q)
	right := detectCommonAncestor(root.right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil && right == nil {
		return nil
	}
	if left != nil {
		return left
	}
	return right
}

func main() {
	p := &commonTree{
		data:  1,
		left:  nil,
		right: nil,
	}
	q := &commonTree{
		data:  2,
		left:  nil,
		right: nil,
	}
	bt := &commonTree{
		data:  "root",
		left:  p,
		right: q,
	}
	node := detectCommonAncestor(bt, p, q)
	fmt.Println("lately common ancestor root node data is:", node.data)
}
