package randomnames

import (
	"testing"
)

func BenchmarkName(b *testing.B) {

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = RandomName()
	}
}

func BenchmarkSafeName(b *testing.B) {

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = SafeRandomName()
	}
}
