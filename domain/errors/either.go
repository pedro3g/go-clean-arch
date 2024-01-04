package errors

type Either interface {
	isEither()
}

type Left[T any] struct {
	Value T
	Error error
}

func (l Left[T]) isEither() {}

type Right[T any] struct {
	Value T
}

func (r Right[T]) isEither() {}
