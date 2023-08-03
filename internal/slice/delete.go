package slice

import "ch1/internal/errs"

func Delete[T any](src []T, index int) ([]T, T, error) {
	length := len(src)
	if index < 0 || index >= length {
		var zero T
		return nil, zero, errs.NewErrIndexOutOfRange(length, index)
	}
	j := 0
	res := src[index]
	for i, v := range src {
		if i != index {
			src[j] = v
			j++
		}
	}
	src = src[:j]
	return src, res, nil
}
