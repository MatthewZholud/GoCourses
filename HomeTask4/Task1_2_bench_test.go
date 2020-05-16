package main

import (
	"testing"
)

func BenchmarkPerimeter(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Perimeter(99)
	}
}

func BenchmarkArea(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Area(99)
	}
}

func BenchmarkEnd(b *testing.B) {
	for n := 0; n < b.N; n++ {
		End(10, 10, 10)
	}
}

