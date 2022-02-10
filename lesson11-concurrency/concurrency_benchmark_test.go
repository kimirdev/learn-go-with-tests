package main

import "testing"

func BenchmarkConcurrentFillArray1000000(b *testing.B) {
	arr := make([]int, 1000000)

	ConcurrentFillArray(arr)
}

func BenchmarkFillArray1000000(b *testing.B) {
	arr := make([]int, 1000000)

	FillArray(arr)
}

func BenchmarkConcurrentFillArray10000000(b *testing.B) {
	arr := make([]int, 10000000)

	ConcurrentFillArray(arr)
}

func BenchmarkFillArray10000000(b *testing.B) {
	arr := make([]int, 10000000)

	FillArray(arr)
}
