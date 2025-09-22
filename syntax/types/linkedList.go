package main

type LinkedList struct {
	head   *node
	tail   *node
	length int //包外访问
}
type node struct {
}

func (l *LinkedList) Add(id int, value any) error {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList) Append(value any) error {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList) Remove(id int) error {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}
