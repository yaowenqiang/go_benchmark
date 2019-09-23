// Package main provides ...
package main

import (
	"fmt"
	"testing"
)

func TestCalculate(t *testing.T) {
	fmt.Println("Test Calculate")
	expected := 4
	result := Calculate(2)
	if expected != result {
		t.Error("Failed")
	}
}

func BenchmarkCalculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Calculate(2)
	}
}

func TestOther(t *testing.T) {
	fmt.Println("Testing something else")
	fmt.Println("This shouldn't run with -run=calc")
}

//go test  -bench=
//go test  -bench=.
//https://tutorialedge.net/golang/benchmarking-your-go-programs/
//go test -run=Calcute -bench=.
//go test -run=Bench -bench=.
