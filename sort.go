package sort

func MergeSort(a []int) []int {
	if len(a) == 1 {
		return a
	}
	var b []int = a[:len(a)/2]
	a = a[len(a)/2:]
	return mergeSort(a, b)
}
func mergeSort(a, b []int) []int {
	if len(a) == 1 && len(b) == 1 {
		return merge(a, b)
	} else if len(a) == 1 {
		b = mergeSort(b[:len(b)/2], b[len(b)/2:])
		return merge(a, b)
	} else if len(b) == 1 {
		a = mergeSort(a[:len(a)/2], a[len(a)/2:])
		return merge(a, b)
	}
	var c, d []int = a[0 : len(a)/2], a[len(a)/2:]
	var e, f []int = b[0 : len(b)/2], b[len(b)/2:]
	var g []int = mergeSort(c, d)
	var h []int = mergeSort(e, f)
	return merge(g, h)
}

// merge merges two sorted slices, ensuring the result is sorted
// things to improve:
//      remove pass-by-value (eg: use pointers perhaps)
//      re-use result array
func merge(a, b []int) []int {
	var r []int = make([]int, 0)
	var i, j int = 0, 0
	for {
		// if a has a lower or equal value, put it on the result
		if a[i] < b[j] || a[i] == b[j] {
			r = append(r, a[i])
			i++
		} else if b[j] < a[i] { // same but for b
			r = append(r, b[j])
			j++
		}
		if i == len(a) { // a exhausted
			for ; j < len(b); j++ {
				r = append(r, b[j])
			}
		} else if j == len(b) { // b exhausted
			for ; i < len(a); i++ {
				r = append(r, a[i])
			}
		}
		if j == len(b) && i == len(a) { // global termination condition
			return r
		}
	}
}
