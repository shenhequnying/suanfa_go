package main

import "fmt"

var ping_result = []int{}

//统计请求次数
func ping_return(t int) int {
	count := 0
	ping_result = append(ping_result, t)
	for ping_result[0] < t-3000 {
		//先出
		ping_result = ping_result[0:]
		count += 1
	}
	return count
}
func main() {
	//先进先出
	queue_slice := []int{1, 2, 3, 4}
	//进
	queue_slice = append(queue_slice, 5)
	fmt.Println(queue_slice)
	//出
	queue_slice = queue_slice[0:]
	fmt.Println(queue_slice)
}
