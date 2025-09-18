package main

func Defer() {
	defer func() {
		println("第一个 defer")
	}()
	defer func() {
		println("第二个 defer")
	}()
}
func Defer2() {
	i := 0
	defer func() {
		//执行的时候值
		println(i)
	}()
	i = 1
}
func DeferV3() {
	i := 0
	defer func(val int) {
		//执行的时候值
		println(i)
		println(val)
	}(i)
	i = 1
}
func DeferReturn() int {
	a := 0
	defer func() {
		a = 1
	}()
	//执行到这里的时候a已经写死了
	return a
}

// 改的是返回值本体
func DeferReturnV2() (a int) {
	a = 0
	defer func() {
		a = 1
	}()
	return a
}
