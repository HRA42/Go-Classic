package fibonacci_test

import (
	"testing"

	"github.com/hra42/go-classic/internal/fibonacci"
)

func TestFibonacci(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		n    uint16
		want uint16
	}{
		{
			name: "0",
			n:    0,
			want: 0,
		},
		{
			name: "1",
			n:    1,
			want: 1,
		},
		{
			name: "2",
			n:    2,
			want: 1,
		},
		{
			name: "3",
			n:    3,
			want: 2,
		},
		{
			name: "4",
			n:    4,
			want: 3,
		},
		{
			name: "5",
			n:    5,
			want: 5,
		},
		{
			name: "6",
			n:    6,
			want: 8,
		},
		{
			name: "7",
			n:    7,
			want: 13,
		},
		{
			name: "8",
			n:    8,
			want: 21,
		},
		{
			name: "9",
			n:    9,
			want: 34,
		},
		{
			name: "10",
			n:    10,
			want: 55,
		},
		{
			name: "11",
			n:    11,
			want: 89,
		},
		{
			name: "12",
			n:    12,
			want: 144,
		},
		{
			name: "13",
			n:    13,
			want: 233,
		},
		{
			name: "14",
			n:    14,
			want: 377,
		},
		{
			name: "15",
			n:    15,
			want: 610,
		},
		{
			name: "16",
			n:    16,
			want: 987,
		},
		{
			name: "17",
			n:    17,
			want: 1597,
		},
		{
			name: "18",
			n:    18,
			want: 2584,
		},
		{
			name: "19",
			n:    19,
			want: 4181,
		},
		{
			name: "20",
			n:    20,
			want: 6765,
		},
		{
			name: "21",
			n:    21,
			want: 10946,
		},
		{
			name: "22",
			n:    22,
			want: 17711,
		},
		{
			name: "23",
			n:    23,
			want: 28657,
		},
		{
			name: "24",
			n:    24,
			want: 46368,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonacci.Fibonacci(tt.n); got != tt.want {
				t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Run BenchmarkFibonacci with loop and recursion
func BenchmarkFibonacci(b *testing.B) {
	b.Run("Fibonacci", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			fibonacci.Fibonacci(24)
		}
	})
	b.Run("FibonacciIterative", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			fibonacci.FibonacciIterative(24)
		}
	})
	b.Run("FibonacciMemo", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			fibonacci.FibonacciMemo(24)
		}
	})
}
