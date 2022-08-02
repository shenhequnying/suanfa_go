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

//归并排序
//把数组拆分，拆分成每一个元素为一个数组
//如何分组？使用递归+中间点下标进行拆分
//进行归并计算，将元素之间，两两比较，小的放前边，大的放后边。
//第二次进行归并计算，比较元素，两两之间，数据的第一个元素之间小的那个，放在新归并数组的头部，再比较第二个元素和刚才的第一个元素（冒泡？）。
//归并时的比较方式很像冒泡，但是比较完的数据进入了新数组，因此整体的遍历中长度比冒泡短的多
//第三次以此类推，最终生成为一个新数组 排序后的。
//不手撕了，以后有机会再聊吧。

//快速排序
//分治法
