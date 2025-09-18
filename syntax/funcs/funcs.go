package main

import "fmt"

func Fun1() {
	fmt.Println("Fun1")
}
func Fun2(a, b string) {
	fmt.Println(a, b)
}
func Fun3(a float64, b, c string) {
	fmt.Println(a, b, c)
}
func Fun4(a float64, b, c string) (string, string) {
	return b, c
}
func Fun5() (name string, a int) {
	//对应数据类型的默认值
	return
}

// 要么都有名字,要么都没有名字
func Fun6() (name string, c int) {
	//对应数据类型的默认值
	return
}
