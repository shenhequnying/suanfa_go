package main

import "fmt"

//链表的操作方式，就是修改next，以及遍历

type Node struct {
	Data int
	Next *Node
}

func (node *Node) iterate() {
	for node != nil {
		fmt.Println(node.Data)
		node = node.Next
	}
}

func (node *Node) delete(node_delete Node) {
	//把删除值的下一个值，赋给删除值
	//再跳过下一个值，把值弄到下下个值
	// node_delete.Data = node_delete.Next.Data
	node_delete.Next = node_delete.Next.Next
}

func (node *Node) revertlist() *Node {
	//双指针遍历
	//什么是双指针遍历
	//通过双指针遍历的过程中，链表中第二个值的data，就是前一个指针最后一次遍历得到的值，next的值，就是之前遍历历史中的值
	cur := node
	var pre *Node = nil
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

func sum(node1, node2 *Node) *Node {
	var node1_tmp *Node = nil
	var node2_tmp *Node = nil
	var final_node *Node = nil
	addition := 0
	for node1.Next != nil || node2.Next != nil {
		//正常实现，防御特殊情况
		node1_tmp = node1
		node2_tmp = node2
		final_node.Data = node1_tmp.Data + node2_tmp.Data + addition
		addition = final_node.Data / 10 % 10
		final_node.Next = new(Node)
		final_node.Next.Data = final_node.Data % 10
		final_node = final_node.Next
	}
	return final_node
}

//这里的repeat 是经过排序的，相邻的值才会出现相等的值
func delete_repeat(node *Node) *Node {
	p := node
	for p != nil {
		if p.Data == p.Next.Data {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return p
}

//链表circle问题
//使用快慢针
func hasCycle(node *Node) bool {
	p1 := node
	p2 := node
	for p1 != nil && p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
		if p1 == p2 {
			return true
		}
	}
	return false
}

func main() {
	//反转链表
	head := new(Node)
	head.Data = 1
	ln2 := new(Node)
	ln2.Data = 2
	ln3 := new(Node)
	ln3.Data = 3
	ln4 := new(Node)
	ln4.Data = 4
	ln5 := new(Node)
	ln5.Data = 5
	head.Next = ln2
	ln2.Next = ln3
	ln3.Next = ln4
	ln4.Next = ln5
	reverse := head.revertlist()
	// // fmt.Println(reverse)
	reverse.iterate()
	fmt.Println(head.Data)
	// head.iterate()

	//两数之和
	// 数1 {1,2,5}(521) 数2 {4,8,9} (984)
	//两链表相加，逢进位加1

}
