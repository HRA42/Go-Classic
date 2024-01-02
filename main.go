package main

import (
	"fmt"

	"github.com/hra42/go-classic/internal/fibonacci"
)

func main() {
	// get n number from command line flag
	n := fibonacci.GetN()

	fmt.Printf("Fibonacci of %d: ", n)
	fmt.Println(fibonacci.Fibonacci(n))
	fmt.Println("Function Calls:", fibonacci.FuncCalls)
	seq := fibonacci.FibonacciSequence(13)
	fmt.Print("Fibonacci Sequence: ")
	for _, number := range seq {
		fmt.Printf("%d ", number)
	}
	fmt.Println("Function Calls:", fibonacci.FuncCalls)
	fmt.Println("Fibonacci Iterative:", fibonacci.FibonacciIterative(n))
	fmt.Println("Loop Calls:", fibonacci.LoopCalls)
}
