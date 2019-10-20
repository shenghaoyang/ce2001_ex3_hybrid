// Copyright 2019 NTU CE2001 SEP1 (AY 19/20) Group 3.
// See LICENSE file for more details.

package main

import (
	"ce2001_ex3_hybrid/customsort"
	"fmt"
)

func main() {

	// Local variable holding number of compares
	compares := 0

	// Anonymous function that is a closure over compares,
	// used to compare between integers and also track the number of key comparisons.
	cmp := func (a, b int) int {
		compares++
		switch {
		case a < b:
			return -1
		case a == b:
			return 0
		case a > b:
			return 1

		// Required for go compiler - a missing return value warning occurs otherwise.
		// It cannot conclude that control flow cannot reach this position since
		// either one of the above cases must be true.
		default:
			panic("unreachable")
		}
	}

	customsort.HybridInsertionMergeSort(arr, 0, len(arr) - 1, 0, cmp)
	fmt.Println(arr, "Compares: ", compares)
}