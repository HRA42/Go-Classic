package prime

import (
	"flag"
	"log"
)

var FuncCalls uint16

// Prime returns true if n is a prime number.
func Prime(n uint16) bool {
	if n <= 1 {
		return false
	}
	for i := uint16(2); i < n; i++ {
		FuncCalls++
		if n%i == 0 {
			return false
		}
	}
	return true
}

// PrimeSequence returns a slice of prime numbers up to n.
func PrimeSequence(n uint16) []uint16 {
	var sequence []uint16
	for i := uint16(0); i <= n; i++ {
		if Prime(i) {
			sequence = append(sequence, i)
		}
	}
	return sequence
}

// GetN returns the n number from the command line flag.
func GetN() uint16 {
	var n uint
	flag.UintVar(&n, "n", 0, "the nth number in the Prime sequence")
	flag.Parse()
	if n > 24 {
		log.Fatal("n is too big, max is 24")
	}
	return uint16(n)
}
