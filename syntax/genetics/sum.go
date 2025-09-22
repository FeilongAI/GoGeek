package genetics

func Sum[T Number](vals ...T) T {
	var res T
	for _, val := range vals {
		res = res + val
	}
	return res
}

type Number interface {
	//联合类型约束：| 符号表示"或"，意思是实现这个接口的类型必须是其中之一
	//~int int类型或者int的衍生类型
	~int | int8 | int16 | int32 | int64
}

/*
- T 是类型参数
- Number 是类型约束，限制 T 只能是之前定义的整数类型
*/
func SumV1[T Number](vals ...T) T {
	var t T
	return t
}
