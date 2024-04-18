package langx

import (
	"github.com/bezhai/multi-bot-task/biz/utils/langx/constraints"
)

func IfElse[T any](b bool, t, f T) T {
	if b {
		return t
	}
	return f
}

func MinInteger[T constraints.Integer](a, b T) T {
	if a <= b {
		return a
	}
	return b
}
