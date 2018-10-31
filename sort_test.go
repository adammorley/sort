package sort

import "testing"

func TestMerge(t *testing.T) {
	var a1 []int = []int{3}
	var b1 []int = []int{4}
	var c1 []int = []int{3, 4}
	compare(t, merge(a1, b1), c1)
	var a0 []int = []int{3, 4, 12, 44, 53}
	var b0 []int = []int{2, 4, 33, 38, 55}
	var c0 []int = []int{2, 3, 4, 4, 12, 33, 38, 44, 53, 55}
	compare(t, merge(a0, b0), c0)
}
func TestMergeSort(t *testing.T) {
	var a []int = []int{33, 1283, 2, 387, 190, 37, 44, 230}
	var b []int = []int{33, 1283, 2, 387, 190, 37, 44, 230, 55}
	compare(t, []int{2, 33, 37, 44, 190, 230, 387, 1283}, MergeSort(a))
    compare(t, MergeSort(b), []int{2,33,37,44,55,190,230,387,1283})
    compare(t, MergeSort([]int{33, 190, 55}), []int{33, 55, 190})
}
func compare(t *testing.T, a, b []int) {
	if len(a) != len(b) {
        t.Error("wrong length")
	}
	for i := range b {
		if b[i] != a[i] {
            t.Error("wrong value")
		}
	}
}


