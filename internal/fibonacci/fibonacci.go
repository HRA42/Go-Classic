package fibonacci

import (
	"flag"
	"log"
)

var FuncCalls uint16
var LoopCalls uint16

// Fibonacci returns the nth number in the Fibonacci sequence.
// Start Values are 0, 1, 1, ...
func Fibonacci(n uint16) uint16 {
	if n <= 1 {
		return n
	}
	FuncCalls++
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func FibonacciSequence(n uint16) []uint16 {
	var sequence []uint16
	for i := uint16(0); i <= n; i++ {
		sequence = append(sequence, Fibonacci(i))
	}
	return sequence
}

func FibonacciIterative(n uint16) uint16 {
	if n <= 1 {
		LoopCalls++
		return n
	}
	var a, b uint16 = 0, 1
	for i := uint16(0); i < n; i++ {
		a, b = b, a+b
		LoopCalls++
	}
	return a
}

// recursion with memoization
var fibMemo = make(map[int]int)

func FibonacciMemo(n int) int {
    if n < 2 {
        return n
    }
    // Wenn das Ergebnis im Map ist, geben Sie es zurück
    if result, ok := fibMemo[n]; ok {
        return result
    }
    // Andernfalls berechnen Sie das Ergebnis und speichern Sie es im Map
    fibMemo[n] = FibonacciMemo(n-1) + FibonacciMemo(n-2)
    return fibMemo[n]
}


func GetN() uint16 {
	var n uint
	flag.UintVar(&n, "n", 0, "the nth number in the Fibonacci sequence")
	flag.Parse()
	if n > 24 {
		log.Fatal("n is too big, max is 24")
	}
	return uint16(n)
}
