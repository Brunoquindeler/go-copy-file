package main

import "testing"

func BenchmarkCopy1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		copy1()
	}
}

func BenchmarkCopy2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		copy2()
	}
}

func BenchmarkCopy3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		copy3()
	}
}

func BenchmarkCopy4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		copy4()
	}
}
