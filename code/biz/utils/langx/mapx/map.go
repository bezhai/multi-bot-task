package mapx

func ConcatReplace[K comparable, V any](dst, src map[K]V) map[K]V {
	for k, v := range src {
		dst[k] = v
	}
	return dst
}
