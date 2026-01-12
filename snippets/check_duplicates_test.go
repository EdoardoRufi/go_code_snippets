package snippets

import (
	"math/rand"
	"strconv"
	"testing"
)

func genDevelopers(n, unique int) []Developer {
	names := make([]string, unique)
	for i := 0; i < unique; i++ {
		names[i] = "dev_" + strconv.Itoa(i)
	}
	devs := make([]Developer, n)
	for i := 0; i < n; i++ {
		devs[i] = Developer{Name: names[rand.Intn(unique)]}
	}
	return devs
}

func benchmarkFilter(b *testing.B, fn func([]Developer) []string) {
	devs := genDevelopers(100_000, 10_000) // 100k totali, 10k unici (esempio)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fn(devs)
	}
}

func BenchmarkFilterUniqueV1(b *testing.B) {
	benchmarkFilter(b, FilterUniqueV1)
}

// func BenchmarkFilterUniqueV2(b *testing.B) {
// 	benchmarkFilter(b, FilterUniqueV2)
// }
