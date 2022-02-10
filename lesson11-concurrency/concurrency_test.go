package main

import (
	"reflect"
	"testing"
)

func TestConcurrentFillArray(t *testing.T) {
	got := make([]int, 1000000)
	ConcurrentFillArray(got)

	want := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		want[i] = i
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Array has been filled wrong")
	}
}
