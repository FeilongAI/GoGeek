package main

import "fmt"

func loopExample() {

	for i := 0; i < 10; i++ {
		println(i)
	}
	i := 0
	for ; i < 10; i++ {
		println(i)
	}
	i = 0
	for i < 10 {
		println(i)
		i++
	}
	for {
		println(i)
	}
	for true {
		println(true)
	}

}
func forArray() {

	arr := [3]string{"1", "2", "3"}
	for i := 0; i < len(arr); i++ {
		println(arr[i])
	}
	for index, value := range arr {
		println(index, value)
	}

}
func ForMap() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for key, value := range m {
		println(key, value)
	}
}
func LoopBug() {
	users := []User{
		{Name: "tom"},
		{Name: "jerry"},
	}
	m := make(map[string]*User, 2)
	for _, u := range users {
		fmt.Printf("Address of u: %p\n", &u) // 直接打印地址
		m[u.Name] = &u
	}

	// 查看 map 中存储的地址是否相同
	for k, v := range m {
		fmt.Printf("%s: %p, value: %+v\n", k, v, *v)
	}
}

// 测试闭包捕获循环变量的经典问题
func ClosureBug() {
	var funcs []func()
	for i := 0; i < 3; i++ {
		funcs = append(funcs, func() {
			fmt.Printf("i = %d, address: %p\n", i, &i)
		})
	}

	fmt.Println("Executing closures:")
	for _, f := range funcs {
		f()
	}
}

type User struct {
	Name string
}
