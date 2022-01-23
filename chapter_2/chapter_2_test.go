package chapter2

import (
	"fmt"
	"testing"
)

func Test_InsertionSort(t *testing.T) {
	origin := []int{5, 2, 4, 6, 1, 3}
	var print []int
	print = append(print, origin...)
	result := InsertionSortAsc(origin)
	fmt.Println("insertion-sort-asc origin:", print)
	fmt.Println("insertion-sort-asc result:", result)

	result = InsertionSortDesc(origin)
	fmt.Println("insertion-sort-desc origin:", print)
	fmt.Println("insertion-sort-desc result:", result)
}

func TestMergeSort(t *testing.T) {
	origin := []int{2, 4, 5, 7, 1, 2, 3, 6}
	var print []int
	print = append(print, origin...)
	result := MergeSort(&origin, 0, 7)
	fmt.Println("merge-sort origin:", print)
	fmt.Println("merge-sort result:", result)
}
