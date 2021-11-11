# asyncwriter

**THIS PACKAGE WAS NOT INTENDED TO BE USED IN PRODUCTION CODE.**

This package shouldn't be used or used only for some rare use cases. To give a good rule of thumb, here is the benchmark result:

```
$ go test -bench=. -run=XXX -benchtime=10s

BenchmarkAsyncMultiWriter_Write_3_writers-8       596774             21188 ns/op
BenchmarkMultiWriter_Write_3_writers-8            595705             18871 ns/op
BenchmarkAsyncMultiWriter_Write_8_writers-8       150171             66789 ns/op
BenchmarkMultiWriter_Write_8_writers-8            267482             63881 ns/op
BenchmarkAsyncMultiWriter_Write_20_writers-8       76065            227821 ns/op
BenchmarkMultiWriter_Write_20_writers-8            72504            175907 ns/op
```

i.e. It might be worth using if you're writing to around **8 or more** writers. Also it depends on the amount of data each writer writes.