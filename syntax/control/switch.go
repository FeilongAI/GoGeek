package main

import "fmt"

func Switch(status int) {
	switch status {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("default")
	}

}
func SwitchBool(age int) {
	switch {
	//case true

	case age >= 18:
		fmt.Println("1")
	case age > 12:
		fmt.Println("2")
	default:
		fmt.Println("default")
	}

}
