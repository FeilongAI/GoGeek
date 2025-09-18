package main

func getAge() {
	age := 20
	if age >= 20 {
		println(age)
	} else if age == 20 {
		println(age)
	} else if age < 20 {
		println(age)
	}
}

func diff(start, end int) {
	if distance := start - end; distance < 0 {
		println(distance)
	} else if distance == 0 {
		println(distance)
	} else if distance > 0 {
		println(distance)
	}
}
