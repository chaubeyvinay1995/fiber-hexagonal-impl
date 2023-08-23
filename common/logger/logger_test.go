package logger

import "testing"

func BenchmarkLogrusPackage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logger.Info("hello")
	}
}

func BenchmarkZeroLogPackage(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}
