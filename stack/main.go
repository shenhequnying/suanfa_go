package main

import (
	"fmt"
	"strings"
)

func kuohaofunc(kuohao string) bool {
	kuohao_stack := []string{}
	kuohao_slice := strings.Split(kuohao, " ")
	for _, label := range kuohao_slice {
		if label == "" {
			continue
		}
		if label == "(" {
			kuohao_stack = append(kuohao_stack, label)
		}
		if label == "{" {
			kuohao_stack = append(kuohao_stack, label)
		}
		if label == "[" {
			kuohao_stack = append(kuohao_stack, label)
		}
		if label == ")" && kuohao_stack[len(kuohao_stack)-1] == "(" {
			kuohao_stack = kuohao_stack[0 : len(kuohao_stack)-1]
		}
		if label == "]" && kuohao_stack[len(kuohao_stack)-1] == "[" {
			kuohao_stack = kuohao_stack[0 : len(kuohao_stack)-1]
		}
		if label == "}" && kuohao_stack[len(kuohao_stack)-1] == "{" {
			kuohao_stack = kuohao_stack[0 : len(kuohao_stack)-1]
		}
	}
	if len(kuohao_stack) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	stack_slice := []int{1, 2, 3, 4}
	//从末尾追加
	stack_slice = append(stack_slice, 5)
	fmt.Println(stack_slice)
	//从末尾删除
	//如果需要返回多个值，则改变1的值
	stack_slice_pop := stack_slice[0 : len(stack_slice)-1]
	fmt.Println(stack_slice_pop)
}
