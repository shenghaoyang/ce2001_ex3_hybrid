# ce2001_ex3_hybrid

Hybrid implementation of merge sort combining insertion sort and merge sort.

This implementation only sorts integers, but could potentially be extended to sort other types,
likely those implementing the interface `sort.Interface`. (But the implementation remains as is for now
since the only requirement is that we sort integers.)

Implemented in order to tackle requirements for Example class 3 in NTU's CE2002 course (AY19/20 Semester 1).

## What's working

- Hybrid merge sort 
    - insertion sort
    - merge sort
        - merge function

- Hybrid merge sort (using auxiliary storage)
    - insertion sort
    - merge sort (using auxiliary storage)
        - merge function (using auxiliary storage)

- Random tests validating these sort functions against the results obtained through go's builtin
  sort functionality.
  
## Building the benchmark tool

```
go get github.com/shenghaoyang/ce2001_ex3_hybrid
go build github.com/shenghaoyang/ce2001_ex3_hybrid
go build -o ce2001_ex3_hybrid github.com/shenghaoayang/ce2001_ex3_hybrid
```

## Generating input data for the benchmark

`./ce2001_ex3_hybrid -mode generate -size <input size> -data <output filename>`

For example:

`./ce2001_ex3_hybrid -mode generate -size 1024 -data benchdata`

Will produce a file `benchdata` in the current working directory containing 1024 integers
that can be used as input for the benchmark tool running in the benchmark mode.

## Running the benchmark

```
./ce2001_ex3_hybrid -mode benchmark -size 1024 -data benchdata \
                    -loops <number of times to run sort function for averaging run times> \
                    -threshold <maximum size of array that will be sorted using insertion sort> \
                    [-aux] (specify this option if you want to benchmark the hybrid merge sort function using
                            auxiliary storage)
```
  
## Benchmark results

Some pre-existing benchmark results for the hybrid sort function (without using auxiliary storage)
are contained in the `/results` folder.
  
## LICENSE

See `LICENSE` for more details.

