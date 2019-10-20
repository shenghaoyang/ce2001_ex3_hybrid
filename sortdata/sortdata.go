// Copyright 2019 NTU CE2001 SEP1 (AY 19/20) Group 3.
// See LICENSE file for more details.

package sortdata

import "math/rand"

// The function Random fills a slice of integers with randomly generated integers using a source.
// The slice is never resliced to access unused capacity.
func Random(out []int, src rand.Source) {
	gen := rand.New(src)

	for i := 0; i < len(out); i++ {
		out[i] = gen.Int()
	}
}
