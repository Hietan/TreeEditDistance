package model

type Label[T any] struct {
	Value T
}

func NewLabel[T any](value T) *Label[T] {
	return &Label[T]{Value: value}
}
