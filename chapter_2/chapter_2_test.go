package chapter2

import (
	"fmt"
	"testing"
)

func Test_InsertionSort(t *testing.T) {
	args := []int{5, 2, 4, 6, 1, 3}
	result := InsertionSortAsc(args)
	fmt.Println("insertion-sort-asc result:", result)

	result = InsertionSortDesc(args)
	fmt.Println("insertion-sort-desc result:", result)
}

func TestMergeSort(t *testing.T) {
	a := []int{2, 4, 5, 7, 1, 2, 3, 6}
	result := MergeSort(a, 0, 7)
	fmt.Println("merge-sort result:", result)
}
