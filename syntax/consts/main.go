package consts

const External = "包外"
const external = "包内"
const (
	a = 123
)
const (
	statusA float64 = iota
	statusB
	statusC
	statusD = iota + 12
	statusF = iota + 12
)

const (
	DayA = iota*12 + 13
	DayB
	DayC
)
