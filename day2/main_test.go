package main

import (
	"fmt"
	"testing"
)

func Benchmark(b *testing.B) {
	for i := 0; i < 1000; i++ {
		day2()
	}

	fmt.Print(b.Elapsed())
}
