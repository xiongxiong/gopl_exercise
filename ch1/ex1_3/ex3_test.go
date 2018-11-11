package ex3_test

import (
	"testing"

	"gopl.io/ch1/ex1_3"
)

func BenchmarkJoinByFor(b *testing.B) {
	args := []string{"hello", "world", "today"}

	res := ex3.JoinByFor(args)
	if res != "hello world today" {
		b.Fatalf("expected -- hello world today, got -- %v", res)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ex3.JoinByFor(args)
	}
}

func BenchmarkJoinByJoin(b *testing.B) {
	args := []string{"hello", "world", "today"}

	res := ex3.JoinByJoin(args)
	if res != "hello world today" {
		b.Fatalf("expected -- hello world today, got -- %v", res)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ex3.JoinByJoin(args)
	}
}
