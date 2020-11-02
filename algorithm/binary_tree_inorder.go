/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/23
 * Time: 13:35
 * 中序遍历二叉树
 */
package main

import "fmt"

type Btree struct {
	Data  interface{}
	left  *Btree
	right *Btree
}

func showBtreeData(btree *Btree) {
	fmt.Println(btree.Data)
}
func inOrder(btree *Btree) {
	if btree != nil {
		if btree.left != nil {
			inOrder(btree.left)
		}
		showBtreeData(btree)
		if btree.right != nil {
			inOrder(btree.right)
		}
	}
}

func main() {
	tree1 := new(Btree)
	tree1.Data = 1
	tree2 := new(Btree)
	tree2.Data = 2
	tree3 := new(Btree)
	tree3.Data = 3
	tree4 := new(Btree)
	tree4.Data = 4
	tree1.left = tree2
	tree1.right = tree3
	tree3.left = tree4
	inOrder(tree1)
}
