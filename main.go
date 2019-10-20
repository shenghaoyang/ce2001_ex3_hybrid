// Copyright 2019 NTU CE2001 SEP1 (AY 19/20) Group 3.
// See LICENSE file for more details.

package main

import (
	"bufio"
	"ce2001_ex3_hybrid/customsort"
	"ce2001_ex3_hybrid/sortdata"
	"flag"
	"fmt"
	"gonum.org/v1/gonum/stat"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile)

	var mode = flag.String("mode", "none", "mode of operation [benchmark or generate]")
	var size = flag.Int("size", 1024, "size of input data to use / output data to generate")
	var dataFile = flag.String("data", "./testData",
		"path to file with input integers to sort or path to destination to write out generated data")
	var loops = flag.Int("loops", 10, "number of loops to use in benchmarking")
	var threshold = flag.Int("threshold", 0, "minumum number of elements required to use mergesort")
	var err error

	flag.Parse()

	if *size < 0 {
		log.Fatalf("input / output size must not be negative (received %v)", *size)
	}

	if *loops <= 0 {
		log.Fatalf("loop count must not be non-positive (received %v)", *loops)
	}

	if *threshold < 0 {
		log.Fatalf("threshold count must not be negative (received %v)", *threshold)
	}

	switch *mode {
	case "benchmark":
		err = benchmark(*dataFile, *size, *loops, *threshold)
	case "generate":
		err = generate(*dataFile, *size, time.Now().UnixNano())
	case "none":
		log.Fatalf("a mode of operation must be specified")
	default:
		log.Fatalf("invalid mode of operation: %s", *mode)
	}

	if err != nil {
		log.Fatalf("error in mode %v: %v", *mode, err)
	}
}

func benchmark(dataFile string, sz, loops, threshold int) error {
	in, err := os.Open(dataFile)
	if err != nil {
		return fmt.Errorf("benchmark: cannot open input file: %v", err)
	}
	defer in.Close()

	r := bufio.NewReader(in)
	data, err := sortdata.LoadData(r, sz)
	if err != nil {
		return fmt.Errorf("benchmark: cannot load data from file: %v", err)
	}

	// Second buffer to sort on, because if we don't reset it we'll sort on the same buffer all the time.
	// That means that from the 2nd loop onwards we'll be sorting sorted data....
	// This prevents that from happening.
	workon := make([]int, len(data))
	copy(workon, data)

	fmt.Printf("Bencharking with input size %v, and threshold %v (%v loops)\n",
		sz, threshold, loops)
	log.Printf("Make sure to ensure a consistent CPU frequency, and bind tasks to cores!")

	// Compare count
	compares := 0
	// Anonymous function for tracking compare count. Used only in the first loop.
	countCmp := func(a, b int) int {
		compares++
		return customsort.AscendingIntComparator(a, b)
	}

	// Separate out the run for finding number of key comparisons since it uses a different compare function
	// than the rest and will hence use a different amount CPU time.
	customsort.HybridInsertionMergeSort(workon, 0, len(workon) - 1, threshold, countCmp)
	log.Printf("Key comparisons: %v", compares)

	// Runtimes per loop in seconds
	times := make([]float64, loops)
	for i := 0; i < loops; i++ {
		copy(workon, data)
		start := time.Now()
		customsort.HybridInsertionMergeSort(workon, 0, len(workon) - 1, threshold,
			customsort.AscendingIntComparator)
		end := time.Now()

		times[i] = end.Sub(start).Seconds()

		log.Printf("Loop iteration %v: %v seconds", i, times[i])
	}

	meanTime, stdDevTime := stat.MeanStdDev(times, nil)
	fmt.Printf("Key comparisons: %v, Average time: %g s, Standard deviation: %g s\n",
		compares, meanTime, stdDevTime)

	return nil
}

// generate generates sz random int32s and saves them to file dataFile, truncating any file with the same name.
// seed is used to initialize the random source.
// The generated int32s are saved in plain-text format, encoded as UTF-8, with each int32 separated from the other
// via newlines.
//
// For example:
//
// 0
// 1
// <EOF>
func generate(dataFile string, sz int, seed int64) error {
	out, err := os.Create(dataFile)
	if err != nil {
		return fmt.Errorf("generate: cannot open output file: %v", err)
	}
	defer out.Close()

	src := rand.NewSource(seed)
	data := make([]int, sz)
	sortdata.Random(data, src)

	if err = sortdata.SaveData(out, data); err != nil {
		return fmt.Errorf("generate: cannot save data to file: %v", err)
	}

	fmt.Printf("wrote %v integers to file %v\n", sz, dataFile)

	return nil
}