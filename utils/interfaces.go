package utils

// Get returns the value of an interface or the default value.
func Get[T any](Interface any) (_ T) {
	if Interface == nil {
		return
	}
	return Interface.(T)
}

// GetOr returns the value of an interface or the argument.
func GetOr[T any](Interface any, val T) (_ T) {
	if Interface == nil {
		return val
	}
	return Interface.(T)
}
