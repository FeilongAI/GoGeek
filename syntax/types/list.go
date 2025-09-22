package main

type List interface {
	Add(id int, value any) error
	Append(value any) error
	Remove(id int) error
	Delete(id int) error
}
