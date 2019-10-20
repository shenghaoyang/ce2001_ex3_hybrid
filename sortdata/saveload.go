package sortdata

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strconv"
)

// loadData reads at most n newline separated 32-bit integers (stored in plain text format, encoding UTF-8)
// from reader r. Callers should provide a buffered reader if required.
// Reading a lower amount of integers is not considered an error. Callers should check the length of the returned
// slice to make sure that the read number of integers matches their requirements.
func LoadData(r io.Reader, n int) ([]int, error) {
	in := bufio.NewScanner(r)
	out := make([]int, 0)


	for i := 0; i < n; i++ {
		if !in.Scan() {
			break
		}

		str := in.Text()
		val, err := strconv.ParseInt(str, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("loadData: cannot convert input to integer: %v", err)
		}

		out = append(out, int(val))
	}

	if e:= in.Err(); e != nil {
		return nil, fmt.Errorf("loadData: cannot scan input: %v", e)
	}

	return out, nil
}

// SaveData saves int32s in slice s to writer w in plain text with UTF-8 encoding.
// On error, some partial data may have been written to w. Callers should discard w in that case.
func SaveData(w io.Writer, s []int) error {
	out := bufio.NewWriter(w)

	for _, v := range s {
		if !((v <= math.MaxInt32) && (v >= math.MinInt32)) {
			return fmt.Errorf("SaveData: cannot save integer %v: integer cannot be represented by type int32",
				v)
		}

		if _, err := out.WriteString(fmt.Sprintln(v)); err != nil {
			return fmt.Errorf("SaveData: cannot save integer %v: %v", v, err)
		}
	}

	if err := out.Flush(); err != nil {
		return fmt.Errorf("SaveData: cannot flush data: %v", err)
	}

	return nil
}
