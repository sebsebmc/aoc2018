package main

import "testing"


func BenchmarkSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slow()
	}
}
//5762701700

func BenchmarkFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fast()
	}
}
//26090978