package chapter2

import (
	"fmt"
	"testing"
)

func Test_InsertionSort(t *testing.T) {
	args := []int{5, 2, 4, 6, 1, 3}
	result := insertion_sort(args)
	fmt.Println("result:", result)
}
