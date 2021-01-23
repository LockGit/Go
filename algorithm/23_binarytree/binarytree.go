/**
 * Created by lock
 * 前、中、后，层序遍历二叉树
 */
package main

import "fmt"

type binaryTree struct {
	data  interface{}
	left  *binaryTree
	right *binaryTree
}

func preOrder(bt *binaryTree) {
	if bt == nil {
		return
	}
	fmt.Println(bt.data)
	if bt.left != nil {
		preOrder(bt.left)
	}
	if bt.right != nil {
		preOrder(bt.right)
	}
}

func inOrder(bt *binaryTree) {
	if bt == nil {
		return
	}
	inOrder(bt.left)
	fmt.Println(bt.data)
	inOrder(bt.right)
}

func levelOrder(bt *binaryTree) {
	queue := []*binaryTree{bt}
	for len(queue) > 0 {
		pop := queue[0]
		queue = queue[1:]
		fmt.Println(pop.data)
		if pop.left != nil {
			queue = append(queue, pop.left)
		}
		if pop.right != nil {
			queue = append(queue, pop.right)
		}
	}
}

func lrdOrder(bt *binaryTree) {
	if bt == nil {
		return
	}
	lrdOrder(bt.left)
	lrdOrder(bt.right)
	fmt.Println(bt.data)
}

func printNodeVal(node *binaryTree, i int) {
	if node == nil {
		return
	}
	if i%2 == 0 {
		fmt.Println(node.data)
	}
	printNodeVal(node.left, i+1)
	printNodeVal(node.right, i+1)
}

func main() {
	bt := &binaryTree{
		data: 0,
		left: &binaryTree{
			data: 1,
			left: &binaryTree{
				data: 3,
			},
			right: &binaryTree{
				data: 4,
			},
		},
		right: &binaryTree{
			data: 2,
			left: &binaryTree{
				data: 5,
			},
			right: &binaryTree{
				data: 6,
				left: &binaryTree{
					data: 7,
				},
			},
		},
	}
	fmt.Println("前序---")
	preOrder(bt)
	fmt.Println("中序---")
	inOrder(bt)
	fmt.Println("后续---")
	lrdOrder(bt)
	fmt.Println("层序---")
	levelOrder(bt)
	fmt.Println("打印偶数层数的节点：")
	printNodeVal(bt, 1)
}
