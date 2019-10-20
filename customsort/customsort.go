// Copyright 2019 NTU CE2001 SEP1 (AY 19/20) Group 3.
// See LICENSE file for more details.

package customsort

// Comparator is a function type for a function that compares two integers. It returns a value smaller than zero
// if the first integer  is to be considered smaller than the second, zero if the two integers are to be considered
// equal, and a value greater than zero if the first integer is to be considered larger than the second.
type Comparator func(int, int) int

// AscendingIntComparator is a function matching the signature of Comparator that can be passed to
// the sort functions in this package to allow sorting of integers in ascending order.
func AscendingIntComparator(a, b int) int {
	switch {
	case a < b:
		return -1
	case a == b:
		return 0
	case a > b:
		return 1
	default:
		panic("unreachable")
	}
}

// HybridInsertionMergeSort uses both merge sort and insertion sort to sort a slice of integers in data
// for elements in the range [first, last] using comparator cmp. Elements are sorted in ascending order.
// min specifies the maximum count of data elements to be sorted before merge sort is used. Any number of elements
// below that minimum will be sorted using insertion sort.
func HybridInsertionMergeSort(data []int, first, last, min int, cmp Comparator) {
	numElems := (last - first) + 1

	if numElems <= 1 {
		return
	}

	if numElems >= min {
		mid := (first + last) / 2
		MergeSort(data, first, mid, cmp)
		MergeSort(data, mid+1, last, cmp)
		merge(data, first, mid, last, cmp)
	} else {
		InsertionSort(data, first, last, cmp)
	}
}

// InsertionSort performs an insertion sort on integers stored in data at positions in range [first, last]
// using comparator cmp. Elements are sorted in ascending order.
// Invalid first, last indexes (last < first, etc,) are not considered to be errors. The InsertionSort function simply
// becomes  a no-op.
func InsertionSort(data []int, first, last int, cmp Comparator) {
	// Early return on invalid or zero size.
	// Could panic or provide an error on invalid size, but we just do nothing silently.
	if (last - first) <= 0 {
		return
	}

	for i := first + 1; i <= last; i++ {

		for j := i; j > first; j-- {
			c := cmp(data[j], data[j - 1])

			switch {
			case c < 0:
				t := data[j - 1]
				data[j - 1] = data[j]
				data[j] = t
			case c >= 0:
				break
			}

		}
	}
}

// MergeSort sorts the integer slice in data for elements in the range [first, last], using comparator function cmp.
// Elements are sorted in ascending order.
func MergeSort(data []int, first, last int, cmp Comparator) {
	// Handle the case when we are sorting 0 elements or sorting only one element:
	// Early return.
	numElems := (last - first) + 1
	if numElems <= 1 {
		return
	}
	mid := (first + last) / 2
	MergeSort(data, first, mid, cmp)
	MergeSort(data, mid + 1, last, cmp)
	merge(data, first, mid, last, cmp)
}

// merge merges two sorted (in ascending order) sublists stored in data with the first sublist stored in
// [first, mid] and the second sublist stored in [mid + 1, last], using comparator function cmp.
// Elements are merged in ascending order.
func merge(data []int, first, mid, last int, cmp Comparator) {
	if (last - first) <= 0 {
		return
	}

	f1 := first		// Index of first element in first sublist
	e1 := mid		// Index of last element in first sublist
	f2 := mid + 1	// Index of first element in second sublist

	for ;; {
		// Exit if any of the sublists are empty
		if ((e1 - f1) < 0) || ((last - f2) < 0) {
			return
		}

		// Compare elements @f1, f2
		c := cmp(data[f1], data[f2])
		switch {
		case c < 0:
			// element @f1 is smaller than that @f2, we take element @f1 directly
			// account for shrunken size of first sublist
			f1++
		case c == 0:
			// element @f1 equal to element @f2
			// save element @f2
			t := data[f2]
			// shift data from [f1 + 1, e1] forward by one index to overwrite element @f2
			// destination is [f1 + 2, e1 + 1]
			// remember slices are exclusive.
			copy(data[f1 + 2:e1 + 2], data[f1 + 1:e1 + 1])
			// store element @f2
			data[f1 + 1] = t
			// do accounting - account for the
			// shrunken size of both the second sublist and the first sublist
			f1 += 2
			e1++
			f2++
		case c > 0:
			// element @f1 > element @ f2
			// save element @f2
			t := data[f2]
			// shift data from [f1, e1] forward by one index to overwrite element @f2
			// destination is [f1 + 2, e1 + 1]
			// remember slices are exclusive.
			copy(data[f1 + 1:e1 + 2], data[f1:e1 + 1])
			// store element @f2
			data[f1] = t
			// do accounting - account for the
			// shrunken size of second sublist and the first sublist moving by one to the right (higher indexes)
			f1++
			e1++
			f2++
		}
	}

}