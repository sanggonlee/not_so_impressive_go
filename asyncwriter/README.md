# asyncwriter

**THIS PACKAGE WAS NOT INTENDED TO BE USED IN PRODUCTION CODE.**

This package shouldn't be used or used only for very rare use cases. To give a good rule of thumb, here is the benchmark result:

```
BenchmarkAsyncMultiWriter_Write_3_writers-8        51506             21142 ns/op
BenchmarkMultiWriter_Write_3_writers-8            121980             13687 ns/op
BenchmarkAsyncMultiWriter_Write_18_writers-8       14233             82911 ns/op
BenchmarkMultiWriter_Write_18_writers-8            20636             68919 ns/op
BenchmarkAsyncMultiWriter_Write_20_writers-8       17862             74831 ns/op
BenchmarkMultiWriter_Write_20_writers-8            18820             78882 ns/op
```

i.e. It might be worth using if you're writing to **20 or more** writers.