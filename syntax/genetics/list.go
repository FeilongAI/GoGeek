package genetics

// T 类型参数，名字叫做 T，约束是 any，等于没有约束
type List[T any] interface {
	Add(idx int, t T)
	Append(t T)
}

func UseList() {
	var l List[int]
	//l.Append("sdad")
	l.Append(12)
}
