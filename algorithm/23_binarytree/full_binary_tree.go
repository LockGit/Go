/**
 * Created by lock
 * 判断是否是满二叉树
 * 深度为k且有 2^k －1 个结点的二叉树称为满二叉树
 */
package main

import (
	"fmt"
	"math"
)

type treeFull struct {
	left  *treeFull
	right *treeFull
	val   interface{}
}

func getMaxDepth(tree *treeFull) (deep int) {
	if tree == nil {
		return 0
	}
	leftDeep := getMaxDepth(tree.left)
	rightDeep := getMaxDepth(tree.right)
	if leftDeep > rightDeep {
		return leftDeep + 1
	}
	return rightDeep + 1
}

//获取一个树的所有节点个数
func getNodeCnt(tree *treeFull) (cnt int) {
	if tree == nil {
		return 0
	}
	left := getNodeCnt(tree.left)
	right := getNodeCnt(tree.right)
	cnt = left + right + 1
	return cnt
}

func main() {
	/*
			    1
			10    11
		 2    3  4   5
	*/
	tree := &treeFull{
		left: &treeFull{
			left: &treeFull{
				val: 2,
			},
			right: &treeFull{
				val: 3,
			},
			val: 10,
		},
		right: &treeFull{
			left: &treeFull{
				val: 4,
			},
			right: &treeFull{
				val: 5,
			},
			val: 11,
		},
		val: 1,
	}
	//fmt.Println("get max tree deep:", getMaxDepth(tree, 0))
	fmt.Println("tree depth is:", getMaxDepth(tree))
	fmt.Println("tree node cnt is:", getNodeCnt(tree))
	fmt.Println("is full tree:", math.Pow(float64(2), float64(getMaxDepth(tree)))-1 == float64(getNodeCnt(tree)))
}
