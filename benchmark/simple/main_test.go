package main

import (
	"testing"
)

func Benchmark_RunUsingGo(b *testing.B) {
	for i := 1; i <= b.N; i++ {
		RunUsingGo(20)
	}
}

func Benchmark_RunUsingCGoOneTime(b *testing.B) {
	for i := 1; i <= b.N; i++ {
		RunUsingCGoOneTime(20)
	}
}

func Benchmark_RunUsingCGoLooping(b *testing.B) {
	for i := 1; i <= b.N; i++ {
		RunUsingCGoLooping(20)
	}
}
