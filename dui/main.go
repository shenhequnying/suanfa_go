package main

//数组表示堆
//左侧子节点的位置是2*index + 1
//右侧子节点的位置是2*index+2
//父节点是（index-1）/2 商

//功能：快速高效的找出最大值和最小值
//时间复杂度o（1）
//找出第k个最大元素
//构建一个最小堆，并将元素依次插入堆中
//当堆的容量超过k，就删除堆顶
//插入结束后，堆顶就是第k个最大元素
//在插入的过程中，所有剩余的值都是比他小的了，保证这个之后，逻辑成立。

type minHeap struct {
	head []int
}

//交换
func (mh minHeap) Swap(index1, index2 int) {
	temp := mh.head[index1]
	mh.head[index1] = mh.head[index2]
	mh.head[index2] = temp
}

//获取父节点
func (mh minHeap) GetParentIndex(index int) int {
	//根据公式..
	return (index - 1) / 2
}

func (mh minHeap) getLeftIndex(index int) int {
	return index*2 + 1
}

func (mh minHeap) getRightIndex(index int) int {
	return index*2 + 2
}

//节点上移，接收元素下标
func (mh minHeap) ShiftUp(index int) int {
	parent_index := mh.GetParentIndex(index)
	if mh.head[parent_index] > mh.head[index] {
		mh.Swap(mh.head[parent_index], mh.head[index])
		//index变了，修改
		mh.ShiftUp(parent_index)
	}
	return len(mh.head)
}

//将值，插入堆的底部，即数组的尾部。
//然后上移，将这个值和他的父节点进行交换，直到父节点小于等于这个插入的值。
func (mh minHeap) insert(number int) {
	mh.head = append(mh.head, number)
	//自动shift，完成堆的目标
	mh.ShiftUp(len(mh.head) - 1)
}

//下移
//不断和子节点交换，直到大于等于当前节点的值
func (mh minHeap) shiftDown(index int) {
	leftindex := mh.getLeftIndex(index)
	rightindex := mh.getRightIndex(index)
	if mh.head[leftindex] < mh.head[index] {
		mh.Swap(leftindex, index)
		mh.shiftDown(leftindex)
	}
	if mh.head[rightindex] < mh.head[index] {
		mh.Swap(rightindex, index)
		mh.shiftDown(rightindex)
	}
}

//用数组尾部元素替换堆顶，直接删除堆顶会破坏堆结构
//然后下移，将新堆顶和他的子节点进行交换，直到子节点大雨等于这个新堆顶
func (mh minHeap) pop() {
	mh.head[0] = mh.head[len(mh.head)-1]
	mh.head = mh.head[:len(mh.head)-1]
	mh.shiftDown(0)
}

func (mh minHeap) top() int {
	return mh.head[0]
}

func (mh minHeap) size() int {
	return len(mh.head)
}

//topk
func (mh minHeap) topK(topk int, numbers []int) int {
	for _, number := range numbers {
		mh.insert(number)
		if mh.size() > topk {
			mh.pop()
		}
	}
	return mh.top()
}

//太难了不看了，先看简单的
// func (mh minHeap) topKfre(nums []int, k int) {
// 	fremap := make(map[int]int)
// 	for _, num := range nums {
// 		fremap[num] += 1
// 	}

// }
