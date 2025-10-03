package slice

import "errors"

func deleteAt[T any](slice []T, index int) ([]T, error) {
	if index < 0 || index >= len(slice) {
		return slice, errors.New("index out of range")
	}

	copy(slice[index:], slice[index+1:])
	slice = slice[:len(slice)-1]
	return slice, nil
}
