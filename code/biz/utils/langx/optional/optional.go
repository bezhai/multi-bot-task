package optional

type Optional[T any] struct {
	value T
	ok    bool
}

func Of[T any](value T) Optional[T] {
	return Optional[T]{
		value: value,
		ok:    true,
	}
}

func Empty[T any]() Optional[T] {
	return Optional[T]{}
}

func (o Optional[T]) IsPresent() bool {
	return o.ok
}

func (o Optional[T]) Get() (t T, ok bool) {
	return o.value, o.ok
}

func (o Optional[T]) OrElse(t T) T {
	if o.ok {
		return o.value
	}
	return t
}

func (o Optional[T]) OrElseLazy(f func() T) T {
	if o.ok {
		return o.value
	}
	return f()
}
