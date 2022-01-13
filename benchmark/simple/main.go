package main

import (
	"fmt"
	"time"
	"unsafe"
)

/*
#include "main.cpp"
*/
import "C"

func Process3(prev int, curr int) int {
	return prev + curr
}

func RunUsingCGoLooping(size int) (x []int) {
	x = make([]int, size)
	var prev int
	for i := 0; i < size; i++ {
		x[i] = int(C.Process1(C.int(prev), C.int(i)))
		prev = x[i]
	}
	return
}

func RunUsingCGoOneTime(size int) (x []int) {
	x = make([]int, size)
	ptr := unsafe.Pointer(C.Process2(C.int(size)))
	defer C.free(ptr)
	gSlice := (*[1 << 30]C.int)(ptr)[:size:size]
	for i := 0; i < size; i++ {
		x[i] = int(gSlice[i])
	}

	return
}

func RunUsingGo(size int) (x []int) {
	x = make([]int, size)
	var prev int
	for i := 0; i < size; i++ {
		x[i] = Process3(prev, i)
		prev = x[i]
	}
	return
}

func main() {
	size := 1000

	var x []int
	startTime := time.Now()
	x = RunUsingCGoLooping(size)
	fmt.Println("EndTime Process Using CGo Looping", time.Since(startTime), x[10])

	x = nil
	startTime = time.Now()
	x = RunUsingCGoOneTime(size)
	fmt.Println("EndTime Process Using CGo One Time Run", time.Since(startTime), x[10])

	x = nil
	startTime = time.Now()
	x = RunUsingGo(size)
	fmt.Println("EndTime Process Using Go", time.Since(startTime), x[10])
}
