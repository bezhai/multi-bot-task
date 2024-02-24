package ptr

// Ptr returns a pointer to the given value.
func Ptr[T any](v T) *T {
	return &v
}

// Value returns the value of the pointer.
func Value[T any](v *T) T {
	if v == nil {
		var zero T
		return zero
	}
	return *v
}

// ValueDefault takes a pointer to a value of type T and a default value of type T.
// It returns the dereferenced value of the pointer if it is not nil.
// If the pointer is nil, it returns the provided default value.
// This function is useful for safely accessing values that may be nil and providing
// a fallback value in such cases.
func ValueDefault[T any](v *T, defaultValue T) T {
	if v == nil {
		return defaultValue
	}
	return *v
}
