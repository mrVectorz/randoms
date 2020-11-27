package main

import (
	"fmt"
	"math"
)

// Only primes less than or equal to N will be generated
const N = 100_000

func main() {
	var x, y, n int
	nsqrt := math.Sqrt(N)

	is_prime := [N]bool{}

	for x = 1; float64(x) <= nsqrt; x++ {
		for y = 1; float64(y) <= nsqrt; y++ {
			n = 4*(x*x) + y*y
				if n <= N && (n%12 == 1 || n%12 == 5) {
					is_prime[n] = !is_prime[n]
				}
				n = 3*(x*x) + y*y
				if n <= N && n%12 == 7 {
					is_prime[n] = !is_prime[n]
				}
				n = 3*(x*x) - y*y
				if x > y && n <= N && n%12 == 11 {
					is_prime[n] = !is_prime[n]
				}
			}
		}

	for n = 5; float64(n) <= nsqrt; n++ {
		if is_prime[n] {
			for y = n * n; y < N; y += n * n {
				is_prime[y] = false
			}
    }
  }

    is_prime[2] = true
    is_prime[3] = true

  primes := make([]int, 0, 1270606)
  for x = 0; x < len(is_prime)-1; x++ {
      if is_prime[x] {
          primes = append(primes, x)
      }
  }

  //for _, x := range primes {
  //  fmt.Println(x)
  //}

	primeCount := 0
  totalSum := 0.0
	largestPrime := ""
  for _, prime := range primes {
    totalSum = totalSum+float64(prime*prime)
    primeCount++
		//fmt.Println(primeCount, " ", prime, " ", totalSum)

    fSum := totalSum/float64(primeCount)
    if fSum == float64(int64(fSum)) {
      //fmt.Println(primeCount, "\t", prime, "\t", totalSum, "\t", fSum)
			largestPrime = fmt.Sprint("N: ", primeCount, "\tPrime: ", prime, "\t", totalSum, "\t", fSum)
    }
  }
	fmt.Println(largestPrime)
}
