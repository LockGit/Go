/**
 * Created by lock
 */
package main

import "fmt"

type LinkedRev struct {
	Value interface{}
	Next  *LinkedRev
}

func createLinkedRev(n int) *LinkedRev {
	l := &LinkedRev{
		Value: "linked-header",
		Next:  nil,
	}
	for i := 0; i < 6; i++ {
		nextNode := &LinkedRev{
			Value: i,
			Next:  nil,
		}
		l.Next = nextNode
		l = l.Next
	}
	//此时指针已经指向链表末尾，链表数据是完整的
	return l
}

//用结构体方法的形式添加节点，好处是外面的l指针扔然指向头部
func (link *LinkedRev) addNode(n int) {
	for i := 0; i <= n; i++ {
		node := &LinkedRev{
			Value: i,
			Next:  nil,
		}
		link.Next = node
		link = link.Next
	}
	return
}

func (link *LinkedRev) revLink() *LinkedRev {
	newLink := new(LinkedRev)
	//header->0->1->2->3->4->5->6 (反转前）
	//header<-0<-1<-2<-3<-4<-5<-6 (反转后）
	for link != nil {
		next := link.Next   //下一个节点
		link.Next = newLink //当前节点
		newLink = link      //当前节点在赋值给新节点
		link = next
	}
	return newLink
}

func main() {
	//linked := createLinkedRev(6)
	//for linked.Next != nil {
	//	fmt.Println(linked.Value)
	//	linked = linked.Next
	//}
	l := new(LinkedRev)
	l.Value = "header"
	//增加n个节点
	l.addNode(6)
	//for l != nil {
	//	fmt.Println(l.Value)
	//	l = l.Next
	//}
	fmt.Println("start revers...")
	//start rev , 反转前，不要把指针提前指到链表尾部，所以上面的代码先注释掉
	newLink := l.revLink()
	var copyLink *LinkedRev
	copyLink = newLink
	for copyLink != nil {
		fmt.Println(copyLink.Value)
		copyLink = copyLink.Next
	}
	//继续反转还原
	oldLink := new(LinkedRev)
	for newLink != nil {
		next := newLink.Next
		newLink.Next = oldLink
		oldLink = newLink
		newLink = next
	}
	fmt.Println("------end----")
	for oldLink != nil {
		fmt.Println(oldLink.Value)
		oldLink = oldLink.Next
	}
}
