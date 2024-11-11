package main

import (
	"testing"
)

func BenchmarkEchoV3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echoV3()
	}
}

func BenchmarkEchoV4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echoV4()
	}
}
