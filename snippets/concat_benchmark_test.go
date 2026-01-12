package snippets

import (
	"fmt"
	"strings"
	"testing"
)

func concatWithPlus() string {
	parte1 := "parte1"
	parte2 := "parte2"
	return parte1 + "_" + parte2
}

func concatWithJoin() string {
	parte1 := "parte1"
	parte2 := "parte2"
	return strings.Join([]string{parte1, parte2}, "_")
}

func concatWithSprintf() string {
	parte1 := "parte1"
	parte2 := "parte2"
	return fmt.Sprintf("%s_%s", parte1, parte2)
}

func BenchmarkConcatWithPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatWithPlus()
	}
}

func BenchmarkConcatWithJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatWithJoin()
	}
}

func BenchmarkConcatWithSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatWithSprintf()
	}
}
