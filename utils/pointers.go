package utils

// Star returns the value of a pointer or the default value.
func Star[T any](pointer *T) (_ T) {
	if pointer == nil {
		return
	}
	return *pointer
}

// StarOr returns the value of a pointer or the argument.
func StarOr[T any](pointer *T, val T) (_ T) {
	if pointer == nil {
		return val
	}
	return *pointer
}

// IfNil returns the pointer or the argument.
func IfNil[T any](pointer *T, val *T) (_ *T) {
	if pointer == nil {
		return val
	}
	return pointer
}

func New[T any](_ *T) (_ T) {
	return
}

func Init[T any](pointer **T) (t T) {
	*pointer = &t
	return
}
