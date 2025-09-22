package main

import "fmt"

func NewUser() {
	//初始化结构体
	u := User{}
	fmt.Printf("%+v", u)
	// up是一个指针
	up := &User{}
	//返回一个指针
	fmt.Printf("%+v", up)
	up2 := new(User)
	fmt.Printf("%+v", up2)
	//声明一个指针再赋值
	var up3 *User
	up3 = new(User)
	fmt.Printf("%+v", up3)
	u4 := User{Name: "Tom", Age: 20}
	fmt.Printf("%+v", u4)
	u5 := User{Name: "Tom", Age: 20}
	fmt.Printf("%+v", u5)
}

type User struct {
	Name      string
	FirstName string
	Age       int
}

func (u User) ChangeName(name string) {
	fmt.Printf("change name 中 u 的地址 %p \n", &u)
	u.Name = name
}

//func ChangeName(u User, name string) {
//
//}

func (u *User) ChangeAge(age int) {
	fmt.Printf("change age 中 u 的地址 %p \n", u)
	u.Age = age
}

//
//func ChangeAge(u *User, age int) {
//
//}

func ChangeUser() {
	u1 := User{Name: "Tom", Age: 18}
	fmt.Printf("u1 的地址 %p \n", &u1)
	//编译器自动 (&u1).ChangeAge(35)
	u1.ChangeAge(35)
	// 这一步执行的时候，其实相当于复制了一个 u1，改的是复制体
	// 所以 u1 原封不动
	u1.ChangeName("Jerry")
	fmt.Printf("%+v\n", u1)

	up1 := &User{}
	fmt.Printf("up1 的地址 %p \n", &up1)
	// (*up1).ChangeName("Jerry")
	up1.ChangeName("Jerry")
	up1.ChangeAge(35)
	fmt.Printf("%+v", up1)
}
