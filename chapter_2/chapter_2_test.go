package chapter2

import (
	"fmt"
	"testing"
)

func Test_InsertionSort(t *testing.T) {
	args := []int{5, 2, 4, 6, 1, 3}
	result := insertionSortAsc(args)
	fmt.Println("asc result:", result)

	result = insertionSortDesc(args)
	fmt.Println("desc result:", result)
}
