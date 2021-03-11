/**
* Created by lock
* 二叉搜索树想查找，删除，增加节点
  二叉搜索树的特点：在树中的任意一个节点，其左子树中的每个节点的值，都要小于这个节点的值，而右子树节点的值都大于这个节点的值。
*/
package main

import "fmt"

type bsTree struct {
	tree *tree
}

type tree struct {
	data  int
	left  *tree
	right *tree
}

func newBsTree() *bsTree {
	bst := &bsTree{}
	bst.tree = &tree{
		data: 8,
		left: &tree{
			data: 5,
			left: &tree{
				data:  3,
				left:  nil,
				right: nil,
			},
			right: nil,
		},
		right: &tree{
			data:  9,
			left:  nil,
			right: nil,
		},
	}
	return bst
}

//二叉树搜索
func (bst *bsTree) search(val int) *tree {
	root := bst.tree
	for root != nil {
		if val > root.data {
			root = root.right
		} else if val < root.data {
			root = root.left
		} else {
			return root
		}
	}
	return root
}

//二叉树新增节点
func (bst *bsTree) addBstNode(val int) *bsTree {
	if bst.tree == nil {
		bst.tree = &tree{
			data: val,
		}
	}
	newNode := &tree{
		data:  val,
		left:  nil,
		right: nil,
	}
	p := bst.tree
	for p != nil {
		if val < p.data {
			if p.left == nil {
				p.left = newNode
				return bst
			}
			p = p.left
		} else {
			if p.right == nil {
				p.right = newNode
				return bst
			}
			p = p.right
		}
	}
	return bst
}

//删除节点
func (bst *bsTree) deleteNode(val int) {
	//查找要删除的节点
	findNode := bst.tree
	var findNodeParent *tree //要删除节点的父节点
	for findNode != nil && findNode.data != val {
		findNodeParent = findNode
		if val > findNode.data {
			findNode = findNode.right
		} else {
			findNode = findNode.left
		}
	}
	if findNode == nil {
		//未找到
		return
	}

	//待删除节点有2个字节点，从待删除节点右子树中 找到 最小节点
	for findNode.left != nil && findNode.right != nil {
		minNode := findNode.right
		minNodeParent := findNode
		//如果有 left 节点 (可能有右子节点，肯定没有左子节点）
		for minNode.left != nil {
			minNodeParent = minNode
			minNode = minNode.left
		}
		//否则
		findNode.data = minNode.data //min node 数据值替换
		//--数据值替换后，修改要删除的对象，变成删除minNode就可以了
		findNode = minNode
		findNodeParent = minNodeParent
	}

	// 待删除节点是叶子节点或者仅有一个子节点
	var child *tree //待删除节点的子节点
	if findNode.left != nil {
		child = findNode.left
	} else if findNode.right != nil {
		child = findNode.right
	} else {
		child = nil
	}
	if findNodeParent == nil {
		bst.tree = child //删除的是根节点
	} else if findNodeParent.left == findNode {
		findNodeParent.left = child
	} else {
		//findNodeParent.right == findNode
		findNodeParent.right = child
	}
}

//二叉搜索树的高度
func (bst *bsTree) maxDepth() (depth int) {
	if bst.tree == nil {
		return 0
	}
	//进行层序
	queue := []*tree{bst.tree}
	qSize := len(queue)
	depth = 0
	front := 0
	for len(queue) > 0 {
		popItem := queue[0]
		queue = queue[1:]
		front++
		if popItem.left != nil {
			queue = append(queue, popItem.left)
		}
		if popItem.right != nil {
			queue = append(queue, popItem.right)
		}
		if front == qSize {
			front = 0
			qSize = len(queue)
			depth++
		}
	}
	return depth
}

func maxDepthBfs(root *tree) int {
	var depth int
	if root == nil {
		return depth
	}

	stack := make([]*tree, 0)
	stack = append(stack, root)

	for len(stack) > 0 {
		depth++
		tmpStack := make([]*tree, 0)
		for len(stack) > 0 {
			top := stack[len(stack)-1]      //末尾移除了一个
			stack = stack[0 : len(stack)-1] //剩下的
			if top.left != nil {
				tmpStack = append(tmpStack, top.left)
			}
			if top.right != nil {
				tmpStack = append(tmpStack, top.right)
			}
		}
		stack = append(stack, tmpStack...)
	}

	return depth
}

func depthRecursion(tree *tree) int {
	if tree == nil {
		return 0
	}
	return max(depthRecursion(tree.left), depthRecursion(tree.right)) + 1
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func (bst *bsTree) recursionMaxDepth() (depth int) {
	if bst.tree == nil {
		return 0
	}
	depth = depthRecursion(bst.tree)
	return depth
}

//中序遍历，顺序输出
func inOrder(bst *tree) (aData []int) {
	if bst == nil {
		return
	}
	aData = append(aData, inOrder(bst.left)...)
	aData = append(aData, bst.data)
	aData = append(aData, inOrder(bst.right)...)
	return
}

func main() {
	bst := newBsTree()
	target := bst.search(3)
	fmt.Println("search target:", target)
	bst.addBstNode(18)
	bst.addBstNode(18)
	bst.addBstNode(28)
	aData := inOrder(bst.tree)
	fmt.Println("inOrder data:", aData)
	fmt.Println("del node")
	bst.deleteNode(3)
	//bst.deleteNode(5)
	bst.deleteNode(8)
	bst.deleteNode(9)
	bst.deleteNode(18)
	bst.deleteNode(18)
	bst.deleteNode(28)
	aData2 := inOrder(bst.tree)
	fmt.Println("after del", aData2)
	fmt.Println("add two node after,max depth is:", bst.maxDepth())
	aData3 := inOrder(bst.tree)
	fmt.Println("aData3 is:", aData3)
	fmt.Println("add two node after,recursion max depth is:", bst.recursionMaxDepth())
	fmt.Println("add two node after,maxDepthBfs max depth is:", maxDepthBfs(bst.tree))
}
