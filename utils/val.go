package utils

// Val is a type that can have a value or be empty.
type Val[T any] struct {
	Val T
	E   bool
}

// Empty returns an empty Val.
func Empty[T any]() (_ Val[T]) {
	return Val[T]{E: true}
}

// V returns the value of a Val or the default value.
func (v Val[T]) V() (_ T) {
	if v.E {
		return
	}
	return v.Val
}

// ValOr returns the value of a Val or the argument.
func (v Val[T]) ValOr(val T) (_ T) {
	if v.E {
		return val
	}
	return v.Val
}
