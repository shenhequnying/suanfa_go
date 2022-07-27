package main

import "strings"

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
	m := make(map[string]string)
	m["["] = "]"
	m["{"] = "}"
	m["("] = ")"
	kuohao_slice := strings.Split(kuohao, " ")
	for _, label := range kuohao_slice {
		if label == "" {
			continue
		}
		if label == "(" {
			kuohao_stack = append(kuohao_stack, label)
		}
	}
}
