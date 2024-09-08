package utils

func Unshift[T any](slice []T, elems ...T) (r []T) {
	slice = append(elems, slice...)

	return slice
}