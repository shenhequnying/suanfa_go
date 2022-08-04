package main

//排序就是把某个乱序的数组变成升序或降序的数组

//冒泡，比较所有相邻的元素，如果第一个比第二个大，则交换。
//最后几轮下来，保证最后一个数是最大的。
//bubble

type test_array struct {
	obj_arr []int
}

//标准的冒泡
//2次for循环，相邻元素前后比较，交换
//时间复杂度on2
func (ta test_array) BubbleSort() {

	for i := 0; i < len(ta.obj_arr)-1; i++ {
		for j := 0; j < len(ta.obj_arr)-1; j++ {
			if ta.obj_arr[j] > ta.obj_arr[j+1] {
				temp := ta.obj_arr[j+1]
				ta.obj_arr[j+1] = ta.obj_arr[j]
				ta.obj_arr[j] = temp
			}
		}
	}
}

//选择排序
//找到数组中最小值，放入第一位
//第二小的值，放入第二位
//执行n-1轮
//时间复杂度也是on2
func (ta test_array) SelectSort() {
	for i := 0; i < len(ta.obj_arr); i++ {
		indexMin := i
		for j := i; j < len(ta.obj_arr); j++ {
			if ta.obj_arr[j] < ta.obj_arr[indexMin] {
				indexMin = j
			}
		}
		//和第一位交换
		temp := ta.obj_arr[i]
		ta.obj_arr[i] = ta.obj_arr[indexMin]
		ta.obj_arr[indexMin] = temp
	}
}

//插入排序
//从第二个元素开始，根前边的元素进行比较。
//如果比前边元素小，则把前边的位置向后移
//如果碰到比自己还小的，插入数据的当前位置则是新位置（break）
//如果碰不到比自己小的，就一直循环到第一个数

func (ta test_array) InsertSort() {
	for i := 0; i < len(ta.obj_arr); i++ {
		temp := ta.obj_arr[i]
		j := i
		for j > 0 {
			if ta.obj_arr[j-1] > temp {
				//向后移一位
				ta.obj_arr[j] = ta.obj_arr[j-1]
			} else {
				break
			}
			j--
		}
		//插入数据的最终位置确定
		ta.obj_arr[j] = temp
	}
}

//递归:
//基线条件+递归条件，啥是基线条件？基线条件就是到什么情况下，停止递归
//啥是递归条件，就是每次递归的时候要缩减范围

//归并排序
//把数组拆分，拆分成每一个元素为一个数组
//如何分组？使用递归+中间点下标进行拆分
//进行归并计算，将元素之间，两两比较，小的放前边，大的放后边。
//第二次进行归并计算，比较元素，两两之间，数据的第一个元素之间小的那个，放在新归并数组的头部，再比较第二个元素和刚才的第一个元素（冒泡？）。
//归并时的比较方式很像冒泡，但是比较完的数据进入了新数组，因此整体的遍历中长度比冒泡短的多
//第三次以此类推，最终生成为一个新数组 排序后的。
//不手撕了，以后有机会再聊吧。
func mergeSort(s []int) []int {
	len := len(s)
	if len == 1 {
		return s
	}
	m := len / 2
	leftS := mergeSort(s[:m])
	rightS := mergeSort(s[m:])
	return merge(leftS, rightS)
}

//两个切片合并
func merge(l []int, r []int) []int {
	lLen := len(l)
	rLen := len(r)
	res := make([]int, 0)
	lIndex, rIndex := 0, 0 //两个切片的下标，插入一个数据，下标+1
	for lIndex < lLen && rIndex < rLen {
		if l[lIndex] > r[rIndex] {
			res = append(res, l[lIndex])
			rIndex++
		} else {
			res = append(res, r[rIndex])
			lIndex++
		}
	}
	if lIndex < lLen {
		res = append(res, l[lIndex:]...)
	}
	if rIndex < rLen {
		res = append(res, r[rIndex:]...)
	}
	return res
}

//快速排序
//分治法
//选定pivot中心轴
//将大于中心轴的数字放在左边
//小于中心轴的数字放在右边
//分别对左右子序列重复前三部操作
func partition(list []int, low, high int) int {
	pivot := list[low]
	for low < high {
		for low < high && pivot <= list[high] {
			//指针左移
			high--
		}
		list[low] = list[high]
		for low < high && pivot >= list[low] {
			low++
		}
		list[high] = list[low]
	}
	list[low] = pivot
	return low
}

func QuickSort(list []int, low, high int) {
	if high > low {
		pivot := partition(list, low, high)
		QuickSort(list, low, pivot-1)
		QuickSort(list, pivot+1, high)
	}
}

//顺序搜索
//遍历数组
//找到根目标值相等的元素，返回下标
//没有的就返回-1
func (ta test_array) sequenceSearch(num int) int {
	for i := 0; i < len(ta.obj_arr); i++ {
		if ta.obj_arr[i] == num {
			return i
		}
	}
	return -1
}

//二分搜索
//有序数组
//从中间，向两端搜索
//如果数组本身是乱序，先对数组进行排序
//ologn
func (ta test_array) binarySearch(item int) int {
	low := 0
	high := len(ta.obj_arr) - 1
	for low <= high {
		mid := (low + high) / 2
		element := ta.obj_arr[mid]
		if element > item {
			low = mid + 1
		} else if element < item {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

//递归完成sum
func Sum(array []int) int {
	if len(array) == 0 {
		return 0
	} else {
		//这个为啥这么写?
		return array[0] + Sum(array[1:])
	}
}

func Total(array []int) int {
	if len(array) == 0 {
		return 0
	} else {
		return 1 + Total(array[1:])
	}
}

func Max(array []int) int {
	max := 0
	if len(array) == 0 {
		return max
	} else {
		if max < array[0] {
			max = array[0]
		}
		return Total(array[1:])
	}
}

// func main() {

// }
