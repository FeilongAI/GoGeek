package main

import "fmt"

func Array() {

	a1 := [3]int{1, 2, 3}
	fmt.Printf("a1:%v,len:%d,cap:%d\n", a1, len(a1), cap(a1))
	a2 := [3]int{1, 2}
	fmt.Printf("a2:%v,len:%d,cap:%d\n", a2, len(a1), cap(a1))
	var a3 = [3]int{}
	fmt.Printf("a3:%v,len:%d,cap:%d\n", a3, len(a3))
}
