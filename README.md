# CGO Benchmark
the purpose of this repository is to benchmark CGO on Golang

## Benchmark Result

goos: darwin <br />
goarch: arm64

- github.com/ivansuteja96/cgo_benchmark/benchmark/simple
```
Function name                           Total run              Average time          Average number of bytes allocated          Number of allocations per operation 
Benchmark_RunUsingGo-8                  40524397               29.31 ns/op           160 B/op                                   1 allocs/op
Benchmark_RunUsingCGoOneTime-8           7651284               157.5 ns/op           176 B/op                                   2 allocs/op
Benchmark_RunUsingCGoLooping-8           1482529               810.2 ns/op           160 B/op                                   1 allocs/op
```
