package env_utils

import (
	"os"
)

func LazyValue(key string) func() string {
	return func() string {
		return os.Getenv(key)
	}
}

func Value(key string) string {
	return os.Getenv(key)
}
