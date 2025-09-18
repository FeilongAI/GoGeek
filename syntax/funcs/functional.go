package main

func functional6() {
	fn := func() string {
		return "Hello World!"
	}
	fn()
}

// 返回一个string的无参方法
func functional7() func() string {
	name := "Hello"
	return func() string {
		return name + " World!"
	}
}

func functional8() {
	//匿名方法立即发起调用
	fn := func() string {
		return "Hello World!"
	}()
	println(fn)
}

//func main() {
//	//functional8()
//	//Defer()
//	//Defer2()
//	DeferV3()
//}
