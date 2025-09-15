package main

import "unicode/utf8"

func String() {
	println("He said \"Hello World!\"")
	println(`
1
2
3
4
4
`)
	println("Hello" + "World!")
	//统计字节数
	println(len("ab"))
	//字符个数
	println(utf8.RuneCountInString("你好abc"))
}
