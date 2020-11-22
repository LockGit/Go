/**
 * Created by lock
 * 基于权重的轮询
每个节点，用它们的当前值加上它们自己的权重。
选择当前值最大的节点为选中节点，并把它的当前值减去所有节点的权重总和。
*/
package round_node

type node struct {
	host    string
	weight  int
	current int
}

type roundRobin struct {
	nodeList []*node
}

func (r *roundRobin) selectNext() (selectNode *node) {
	if len(r.nodeList) == 0 {
		return
	}
	sum := 0
	for _, node := range r.nodeList {
		sum = sum + node.weight
		node.current += node.weight
		if selectNode == nil || node.current > selectNode.current {
			selectNode = node
		}
	}
	if selectNode == nil {
		return
	}
	selectNode.current -= sum
	return
}
