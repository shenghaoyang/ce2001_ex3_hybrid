# ce2002_ex3_hybrid

Hybrid implementation of merge sort combining insertion sort and merge sort.

This implementation only sorts integers, but could potentially be extended to sort other types,
likely those implementing the interface sort.Interface. (But the implementation remains as is for now
since the only requirement is that we sort integers.)

Implemented in order to tackle requirements for Example class 3 in NTU's CE2002 course (AY19/20 Semester 1).

## What's working

- Hybrid merge sort 
    - insertion sort
    - merge sort
        - merge function
        
- Random tests validating these sort functions against the results obtained through go's builtin
  sort functionality.
  
## LICENSE

See `LICENSE` for more details.

