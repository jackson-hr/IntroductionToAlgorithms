package chapter2

/**
 * 伪码:
 * INSERTION-SORT(A)
 * for j = 2 to A.length
 * 		key = A[j]
 * 		// insert A[j] into the sorted sequence A[1...j - 1]
 *		i = j - 1
 *		while i > 0 and A[i] > key
 *				A[i + 1] = A[i]
 *				i = i - 1
 *		A[i + 1] = key
 */

// InsertionSortAsc 插入排序(正序)
// 1、从第二个元素(下标为1)开始和前面数据进行比较
// 2、如果比当前值大则右移前面的数值
// 3、直到找到当前数据的对应位置则结束查找插入数据
func InsertionSortAsc(args []int) []int {
	for j := 1; j < len(args); j++ {
		key := args[j]
		i := j - 1
		for ; i >= 0 && args[i] > key; i-- {
			args[i+1] = args[i]
		}
		args[i+1] = key
	}
	return args
}

// InsertionSortDesc 插入排序(倒序)
func InsertionSortDesc(args []int) []int {
	for j := 1; j < len(args); j++ {
		key := args[j]
		i := j - 1
		for ; i >= 0 && args[i] < key; i-- {
			args[i+1] = args[i]
		}
		args[i+1] = key
	}
	return args
}
