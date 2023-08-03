package slice

import "ch1/internal/slice"

// Delete 删除 index 处的元素
func Delete[Src any](src []Src, index int) ([]Src, error) {
	res, _, err := slice.Delete[Src](src, index)
	return res, err
}

func Shrink[Src any](src []Src) ([]Src, error) {
	res := slice.Shrink[Src](src)
	return res, nil
}
