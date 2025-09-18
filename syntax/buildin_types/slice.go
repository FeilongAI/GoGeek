package main

import "fmt"

func Slice() {
	a1 := []int{1, 2, 3}
	fmt.Printf("a1:%v,len:%d,cap:%d\n", a1, len(a1), cap(a1))
	//3 长度 4 容量
	s2 := make([]int, 3, 4)
	fmt.Printf("s2:%v,len:%d,cap:%d\n", s2, len(s2), cap(s2))
	//长度和容量都是4,{0,0,0,0}
	s3 := make([]int, 4)
	//{} s3 := make([]int, 0,4)
	fmt.Printf("s2:%v,len:%d,cap:%d\n", s3, len(s3), cap(s3))
	//长度已经放了多少个，容量能放多少个，可以扩容
	s4 := make([]int, 0, 4)
	s4 = append(s4, 0)
	fmt.Printf("s2:%v,len:%d,cap:%d\n", s4, len(s4), cap(s4))
}
func Subslice() {

	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//左闭右开
	//容量从start 开始算  s1已经申请空间的元素
	//s2:[2 3],len:2,cap:8
	s2 := s1[1:3]
	fmt.Printf("s2:%v,len:%d,cap:%d\n", s2, len(s2), cap(s2))
}
func ShareSlice() {
	s1 := []int{1, 2, 3, 4}
	s2 := s1[2:]
	fmt.Printf("s2: %v, len=%d, cap=%d \n", s2, len(s2), cap(s2))

	s2[0] = 99
	fmt.Printf("s2: %v, len=%d, cap=%d \n", s2, len(s2), cap(s2))
	fmt.Printf("s1: %v, len=%d, cap=%d \n", s1, len(s1), cap(s1))

	s2 = append(s2, 199)
	fmt.Printf("s2: %v, len=%d, cap=%d \n", s2, len(s2), cap(s2))
	fmt.Printf("s1: %v, len=%d, cap=%d \n", s1, len(s1), cap(s1))
}
