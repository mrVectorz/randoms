package main

import (
	"fmt"
)

// Only primes less than or equal to N will be generated
const N = 100_000

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

func main() {
  totalSum := 0.0
  largestPrime := ""
	ch := make(chan int) // Create a new channel.
	go Generate(ch)      // Launch Generate goroutine.
	for primeCount := 1; primeCount < N; primeCount++ {
		prime := <-ch
		//fmt.Println(prime)
		    totalSum = totalSum+float64(prime*prime)

    fSum := totalSum/float64(primeCount)
    if fSum == float64(int64(fSum)) {
      //fmt.Println(primeCount, "\t", prime, "\t", totalSum, "\t", fSum)
      largestPrime = fmt.Sprint("N: ", primeCount, "\tPrime: ", prime, "\t", totalSum, "\t", fSum)
    }

		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
	fmt.Println(largestPrime)
}
