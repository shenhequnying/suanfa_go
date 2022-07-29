package main

import (
	"fmt"
	"math"
)

type Node struct {
	Value    string
	Children []*Node
}

type ErNode struct {
	Value               string
	LeftNode, RightNode *ErNode
}

func dfs(node *Node) {
	//直接递归就行了
	//在遍历过程中，第一级执行时，先执行第一个元素的值以及其对应的子相关节点，执行到没有child之后
	//再回来执行第一级的第二个节点及其子元素
	fmt.Println(node.Value)
	for _, node_child := range node.Children {
		dfs(node_child)
	}
}

//先序遍历（优先遍历左侧节点）
//访问跟节点
//对跟节点的左子树先进行先序遍历
//对跟节点的右子树进行先序遍历
//根左右
func xianxudfs(ernode *ErNode) {
	fmt.Println(ernode.Value)
	xianxudfs(ernode.LeftNode)
	xianxudfs(ernode.RightNode)
}

//中序遍历
//先对跟节点左子树进行中序遍历
//访问跟节点
//对跟节点右子树进行中序遍历
//左根右
func zhongxudfs(ernode *ErNode) {
	zhongxudfs(ernode.LeftNode)
	fmt.Println(ernode.Value)
	zhongxudfs(ernode.RightNode)
}

//后序遍历
//先对跟节点左子树进行后续遍历
//对跟节点右子树进行后续遍历
//访问跟节点
//左右根
func houxudfs(ernode *ErNode) {
	houxudfs(ernode.LeftNode)
	houxudfs(ernode.RightNode)
	fmt.Println(ernode.Value)
}

//口诀，新建一个队列，把跟节点入队
//把队头出队并访问
//把队头的children挨个入队
//重复第二第三步，直到队列为空

func bfs(node *Node) {
	final_q := []*Node{node}
	for {
		//正好写反了之前
		//这里应该是保证输出的队列里的长度要大于0
		if len(final_q) == 0 {
			break
		}
		fmt.Println(node.Value)
		final_q = final_q[0:]
		for _, child_node := range node.Children {
			final_q = append(final_q, child_node)
		}
	}
}

//先序遍历就是二叉树的深度优先
//进行深度优先遍历
//遍历过程中记录深度，最后将深度返回
func xianxudfs_deep(ernode *ErNode, count, res int) int {
	// fmt.Println(ernode.Value)
	if (ernode.LeftNode == nil) && (ernode.RightNode == nil) {
		res = int(math.Max(float64(res), float64(count)))
	}
	xianxudfs_deep(ernode.LeftNode, count+1, res)
	xianxudfs_deep(ernode.RightNode, count+1, res)
	return res
}

//二叉树最大深度
func deepbfs(node *ErNode) {
	res := 0
	count := 0
	res = xianxudfs_deep(node, count, res)
	fmt.Println(res)
}

//二叉树最小深度
//使用广度优先队列的模式进行计算
//当广度优先到节点是叶子结点时，返回深度值
//最小的深度值作为最终返回
func mindfs(ernode *ErNode) int {
	res := 0
	deep := 0
	final_q := []*ErNode{ernode}
	for {
		if len(final_q) == 0 {
			break
		}
		deep += 1
		fmt.Println(ernode.Value)
		final_q = final_q[0:]
		if ernode.LeftNode != nil {
			final_q = append(final_q, ernode.LeftNode)
		}
		if ernode.RightNode != nil {
			final_q = append(final_q, ernode.RightNode)
		}
		if (ernode.LeftNode == nil) && (ernode.RightNode == nil) {
			res = int(math.Max(float64(res), float64(deep)))
		}
	}
	return res
}

//二叉树层序遍历
//二叉树每一层的值，作为一个数组返回，整体放入一个大数组中。
//广度优先，每一次遍历，放入一个数组，然后返回？
//这里其实应该使用map的概念，每一级推送的时候，附带层级标签
func cengxutree(ernode *ErNode) [][]string {
	m := [][]string
	final_q := []*ErNode{ernode}
	deep := 0
	for {
		if len(final_q) == 0 {
			break
		}
		// m[deep] = []string{ernode.Value}
		final_q = final_q[0:]
		if (m[deep]==nil){
			m[deep] = []string{ernode.Value}
		}
		if ernode.LeftNode != nil {
			deep += 1
			final_q = append(final_q, ernode.LeftNode)
			m[deep+1] = append(m[deep+1], ernode.LeftNode.Value)
		}
		if ernode.RightNode != nil {
			final_q = append(final_q, ernode.RightNode)
			m[deep+1] = append(m[deep+1], ernode.LeftNode.Value)
		}

	}
	return m
}

//路径总和
//使用深度优先来做

//深度优先
func erdfs(ernode *ErNode, num int){
	if (ernode.left !=nil){
		erdfs(ernode.left, num+ernode.LeftNode.Value)
	}
	if (ernode.right !=nil){
		erdfs(ernode.right, num+ernode.RightNode)
	}
}

//深度优先求路径总和
func hasPathSum(ernode *ErNode, targetsum int) bool{
	if root == nil {
		return false
	}
	if (root.LeftNode == nil) && (root.RightNode == nil){
		return targetsum - ernode.Value
	}
	return hasPathSum(root.LeftNode, targetsum-ernode.Value) || hasPathSum(root.RightNode, targetsum-ernode.Value)
}

//非递归版本的先中后序遍历

func main() {
	//创建树
	var node11 = Node{Value: "c", Children: nil}
	var node12 = Node{Value: "d", Children: nil}
	var node1 = Node{Value: "b", Children: []*Node{&node11, &node12}}
	var node21 = Node{Value: "f", Children: nil}
	var node22 = Node{Value: "g", Children: nil}
	var node2 = Node{Value: "e", Children: []*Node{&node21, &node22}}
	var tree = Node{Value: "a", Children: []*Node{&node1, &node2}}
	fmt.Println(tree)
}
