package chapter2

import (
	"fmt"
	"testing"
)

func Test_InsertionSort(t *testing.T) {
	args := []int{5, 2, 4, 6, 1, 3}
	result := insertion_sort_asc(args)
	fmt.Println("result:", result)

	result = insertion_sort_desc(args)
	fmt.Println("result:", result)
}
