package main

import (
	"fmt"
	"math"
	"strings"
)

func removeDuplicateElement(to_set []int) []int {
	//通过map的key方式做去重
	result := make([]int, 0, len(to_set))
	temp := map[int]struct{}{}
	for _, item := range to_set {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

//双数组求集合
//go求交集必须用到map
func intersection(array_1, array_2 []int) []int {
	m := make(map[int]int)
	result_list := make([]int, 0)
	for _, ele1 := range array_1 {
		m[ele1] += 1
	}
	for _, ele2 := range array_2 {
		//这里可能需要做异常判断
		count, _ := m[ele2]
		if count > 0 {
			result_list = append(result_list, ele2)
		}
	}
	return result_list
}

func kuohao_new(kuohao string) bool {
	//map不计算在空间复杂度内，是o1的
	kuohao_stack := []string{}
	m := make(map[string]string)
	m["["] = "]"
	m["{"] = "}"
	m["("] = ")"
	kuohao_slice := strings.Split(kuohao, " ")
	for _, label := range kuohao_slice {
		if label == "" {
			continue
		}
		if _, ok := m[label]; ok {
			kuohao_stack = append(kuohao_stack, label)
		}
		//还是需要获取栈顶元素
		//判断栈顶是否和map中的值对应起来
		if m[kuohao_stack[len(kuohao_stack)-1]] == label {
			kuohao_stack = kuohao_stack[0 : len(kuohao_stack)-1]
		}
	}
	if (len(kuohao_stack)) > 0 {
		return false
	} else {
		return true
	}
}

//两数之和
//双指针遍历？
//用map实现更简单...
func liangshu(object_slice []int, target_num int) bool {
	//先创建一个map
	m := make(map[int]int)
	for label, number := range object_slice {
		m[number] = label
	}
	for number := range object_slice {
		if number <= target_num {
			if _, ok := m[target_num-number]; ok {
				fmt.Println("找到2数，分别为: ", number, m[number], target_num-number, m[target_num-number])
				return true
			}
		}
	}
	return false
}

//无重复字符串最长子串
//记录每一个不重复字符串的长度
//如果出现重复字符，则该记录清零重新计算
//将计算结果放入slice中，取最大值
//这种算法可能不准，不如双指针滑动窗口
func norepeatstr(object_slice []string) int {
	m := make(map[string]int)
	j := []int{}
	sum := 0
	max_len := 0
	for _, char := range object_slice {
		sum += 1
		if _, ok := m[char]; ok {
			m[char] += 1
			if m[char] > 1 {
				j = append(j, sum)
				sum = 0
			}
		}
	}
	//这里其实可以使用max函数
	for _, sum_num := range j {
		if sum_num > max_len {
			max_len = sum_num
		}
	}
	return max_len
}

//双指针滑动窗口法
//右指针遍历，左指针在出现重复字段时定位到之前记录的重复字段的位置+1
//节省，在窗口滑动的过程中，有可能不停的变化，这时候需要记录最大值，使用max函数实现
func norepeatstr_new(object_slice []string) int {
	left := 0
	right := 0
	max_len := 0.0
	m := make(map[string]int)
	for right = 0; right < len(object_slice); right++ {
		if _, ok := m[object_slice[right]]; ok {
			//滑动窗口的+1，在这里需要是获取到的重复值的下标+1，总感觉不太对的样子？
			left = m[object_slice[right]] + 1
			max_len = math.Max(max_len, float64(right-left))
		}
		m[object_slice[right]] = right
	}
	return int(max_len)
}

//最小字符串，有点难不看了。
