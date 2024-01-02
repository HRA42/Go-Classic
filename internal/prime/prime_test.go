package prime_test

import (
	"testing"

	"github.com/hra42/go-classic/internal/prime"
)

func TestPrime(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		n    uint16
		want bool
	}{
		{
			name: "0",
			n:    0,
			want: false,
		},
		{
			name: "1",
			n:    1,
			want: false,
		},
		{
			name: "2",
			n:    2,
			want: true,
		},
		{
			name: "3",
			n:    3,
			want: true,
		},
		{
			name: "4",
			n:    4,
			want: false,
		},
		{
			name: "5",
			n:    5,
			want: true,
		},
		{
			name: "6",
			n:    6,
			want: false,
		},
		{
			name: "7",
			n:    7,
			want: true,
		},
		{
			name: "8",
			n:    8,
			want: false,
		},
		{
			name: "9",
			n:    9,
			want: false,
		},
		{
			name: "10",
			n:    10,
			want: false,
		},
		{
			name: "11",
			n:    11,
			want: true,
		},
		{
			name: "12",
			n:    12,
			want: false,
		},
		{
			name: "13",
			n:    13,
			want: true,
		},
		{
			name: "14",
			n:    14,
			want: false,
		},
		{
			name: "15",
			n:    15,
			want: false,
		},
		{
			name: "16",
			n:    16,
			want: false,
		},
		{
			name: "17",
			n:    17,
			want: true,
		},
		{
			name: "18",
			n:    18,
			want: false,
		},
		{
			name: "19",
			n:    19,
			want: true,
		},
		{
			name: "20",
			n:    20,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prime.Prime(tt.n); got != tt.want {
				t.Errorf("Prime() = %v, want %v", got, tt.want)
			}
		})
	}
}