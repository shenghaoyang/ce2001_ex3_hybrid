# Benchmark results

The benchmark results are stored in folders named `<host configuration>[_aux]`, where:

- `host_configuration` represents the configuration of the system used to run the
  benchmarks, and

- `[aux]` indicates whether the results were generated for the hybrid sort function using
  auxiliary storage.

## Result file

### Naming

In each folder, there are numerous result files in the form:

- `<input_size>_<threshold_value>_<loops>.txt`, where:

    - `input_size` represents the size of the input given to the sort function.
    - `threshold_value` represents the input size threshold value, above which the sort function uses mergesort.
    - `loops` represents the amount of timing loops used to produce the average sort function runtime.
    
### Content

The content of each result file should be self-explanatory.