package utils

import (
	rt "reflect"
	st "strings"
)

// Name returns the name of a type.
func Name(item any) (_ string) {
	return rt.TypeOf(item).Name()
}

// NameLower returns the lower case name of a type.
func NameLower(item any) (_ string) {
	return st.ToLower(Name(item))
}

// Default returns the default value for a type.
func Default[T any]() (_ T) {
	return
}

// IsDefault returns whether the value is the default value.
func IsDefault[T comparable](comp T) (_ bool) {
	return comp == Default[T]()
}

// E returns true if the value is not the default value.
func E[T comparable](val T) (_ bool) {
	return val != Default[T]()
}

// IfDefault returns the value or the argument.
func IfDefault[T comparable](comp T, val T) (_ T) {
	if IsDefault(comp) {
		return val
	}
	return comp
}

// Nil returns nil if the value is the default value.
func Nil[T comparable](comp T) (_ any) {
	if IsDefault(comp) {
		return nil
	}
	return comp
}
