package chapter2

/**
 * 伪码:
 *MERGE(A, p, q, r)
 *  n1 = q - p + 1
 *  n2 = r - q
 *  let L[1..n1 + 1] and R[1..n2 + 1] be new arrays
 *  for i = 1 to n1
 *      L[i] = A[p + i - 1]
 *  for j = 1 to n2
 *      R[j] = A[q + j]
 *  L[n1 + 1] = ∞
 *  R[n2 + 1] = ∞
 *  i = 1
 *  j = 1
 *  for k = p to r
 *      if L[i] <= R[j]
 *          A[k] = L[i]
 *          i = i + 1
 *      else A[k] = R[j]
 *          j = j + 1
 *
 *MERGE-SORT(A, p, r)
 *  if p < r
 *      q = (p + r)/2
 *      MERGE-SORT(A, p, q)
 *      MERGE-SORT(A, q + 1, r)
 *      MERGE(A, p, q, r)
 */

// merge 子序列排序
// a 需要排序的数组
// l 左下标, m 中间分割下标, r 右下标
func merge(a []int, l, m, r int) []int {
	ll := m - l + 1
	rl := r - m
	var L = make([]int, ll)
	var R = make([]int, rl)
	for i := 0; i < ll; i++ {
		L[i] = a[l+i-1]
	}
	for j := 0; j < rl; j++ {
		R[j] = a[m+j]
	}
	i, j := 1, 1
	for ind := l; ind < r; ind++ {
		if L[i] <= R[j] {
			a[ind] = L[i]
			i = i + 1
		} else {
			a[ind] = R[j]
			j = j + 1
		}
	}
	return a
}

// MergeSort 归并排序
func MergeSort(a []int, l, r int) []int {
	if l >= r {
		return a
	}
	m := (l + r) / 2
	MergeSort(a, l, m)
	MergeSort(a, m+1, r)
	return merge(a, l, m, r)
}
