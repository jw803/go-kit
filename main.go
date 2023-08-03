package main

import "ch1/slice"

func main() {
	nationalities := []string{"China", "Japan", "Korea", "Singapore"}
	newNationalities, _ := slice.Delete(nationalities, 2)
	println("%v", newNationalities)

	numbers := []int{0, 1, 2, 3}
	newNumbers, _ := slice.Delete(numbers, 2)
	println("%v", newNumbers)

	shrinkNumbers := make([]int, 0, 100)
	exampleNumbers := []int{0, 1}
	shrinkNumbers = append(shrinkNumbers, exampleNumbers...)
	newShrinkNumbers, _ := slice.Shrink(shrinkNumbers)
	println("%v", newShrinkNumbers)
}
