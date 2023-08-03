package slice

func calCapacity(c, l int) (int, bool) {
	if c <= 64 {
		return c, false
	}
	if c > 2048 && (c/2 >= l) {
		factor := 0.6
		return int(float32(c) * float32(factor)), true
	}
	if c <= 2048 && (c/4 >= l) {
		return c / 2, true
	}
	return c, false
}

func Shrink[T any](src []T) []T {
	c, l := cap(src), len(src)
	n, changed := calCapacity(c, l)
	if !changed {
		return src
	}
	s := make([]T, 0, n)
	s = append(s, src...)
	return s
}
