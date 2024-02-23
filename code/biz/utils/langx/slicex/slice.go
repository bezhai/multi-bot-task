package slicex

import (
	"github.com/bezhai/multi-bot-task/biz/utils/langx/optional"
)

func Get[T any](a []T, index int) optional.Optional[T] {
	if index < 0 {
		index += len(a)
	}
	if index < 0 || index >= len(a) {
		return optional.Empty[T]()
	}
	return optional.Of(a[index])
}

func Map[From any, To any](src []From, fn func(From) To) []To {
	if len(src) == 0 {
		return nil
	}

	dst := make([]To, 0, len(src))
	for _, srcElem := range src {
		dst = append(dst, fn(srcElem))
	}

	return dst
}
