// Copyright 2019 NTU CE2001 SEP1 (AY 19/20) Group 3.
// See LICENSE file for more details.

package test

import (
	"ce2001_ex3_hybrid/customsort"
	"ce2001_ex3_hybrid/sortdata"
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

type sortfn func(data []int, first, last int, comparator customsort.Comparator)

// TestSorts tests the three sort implementations, MergeSort, InsertionSort, and HybridInsertionMergeSort
// for correctness against go's built in sort functionality.
// The test is performed on input sizes from 2^1 to 2^17 (both inclusive) in ascending powers of 2. Randomly
// generated data is sorted by both our implementation and the implementation in go's standard library.
// The test is considered to have passed if the sort results produced by both methods match.
func TestSorts(t *testing.T) {
	max := 1 << 17
	src := rand.NewSource(time.Now().UnixNano())
	refSortBuf := make([]int, max)
	sortBuf := make([]int, max)
	dataBuf := make([]int, max)
	aux := make([]int, max)

	sortdata.Random(dataBuf, src)

	tfn := func(buf, refBuf, data []int, fn sortfn) bool {
		copy(buf, data)
		copy(refBuf, data)

		sort.IntSlice(refBuf).Sort()
		fn(buf, 0, len(buf)-1, customsort.AscendingIntComparator)

		return reflect.DeepEqual(buf, refBuf)
	}

	sortfns := [...]sortfn{customsort.MergeSort, customsort.InsertionSort,
		func(data []int, f, l int, comparator customsort.Comparator) {
			customsort.HybridInsertionMergeSort(data, f, l, 16, comparator)
		},
		func(data []int, f, l int, comparator customsort.Comparator) {
			customsort.MergeSortAux(data, aux[f:l+1], f, l, comparator)
		},
		func(data []int, f, l int, comparator customsort.Comparator) {
			customsort.HybridInsertionMergeSortAux(data, aux[f:l+1], f, l, 16, comparator)
		}}
	sortfnNames := [...]string{"MergeSort", "InsertionSort", "HybridInsertionMergeSortMin16",
		"MergeSortAux", "HybridInsertionMergeSortAuxMin16"}

	for j, fn := range sortfns {
		for i := 1; (i << 1) <= max; i <<= 1 {
			t.Run(fmt.Sprintf("Sort: %v with input length %d", sortfnNames[j], i),
				func(t *testing.T) {
					if !tfn(sortBuf[0:i], refSortBuf[0:i], dataBuf[0:i], fn) {
						t.Errorf("Custom sort algorithm does not match result from that built-in to go.")
					}
				})
		}
	}
}
