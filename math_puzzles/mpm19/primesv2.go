package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"io"
	"bufio"
//	"encoding/binary"
//	"math"

	"github.com/jbarham/primegen"
)

var low, high uint64 = 2, 100_000_000

func usage() {
	fmt.Fprintf(os.Stderr, "usage: primes [[low=%d] high=%d]\n", low, high)
	os.Exit(2)
}

/*
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
*/

func float64ToByte(f float64) []byte {
	var i int
	xx := int64(f)
	const buf = 32
	//var i int
	digits := make([]byte, 32)
	digits[ndigits-1] = ' '
  //binary.BigEndian.PutUint64(buf[:], math.Float64bits(f))
  //return buf[:]
	for i = buf - 2; i > 0; i-- {
		digits[i] = '0' + byte(xx%10)
		if xx < 10 {
			break
		}
		xx /= 10
	}
	return digits[i:]
}

const ndigits = 32
func Write(w io.Writer, low, high uint64) (err error) {
	//var i int
	digits := make([]byte, 32)
	digits[ndigits-1] = '\n'

	totalSum := 0.0
	primeCount := 0
	//largestPrime := 0.0
	largestPrime := ""

	sieve := primegen.New()
	sieve.SkipTo(low)
	x := sieve.Next()
	b := bufio.NewWriter(w)
	for x < high {
		primeCount++
		//fmt.Println(x, " ", primeCount)
		/*
		xx := x
		for i = ndigits - 2; i > 0; i-- {
			digits[i] = '0' + byte(xx%10)
			if xx < 10 {
				break
			}
			xx /= 10
		}
		*/
		totalSum = totalSum+float64(x*x)
		fSum := totalSum/float64(primeCount)
		if fSum == float64(int64(fSum)) {
			//largestPrime = fSum
			largestPrime = fmt.Sprint("N: ", primeCount, "\tPrime: ", x, "\t", totalSum, "\t", fSum, "\n")
		}
		/*
		if _, err = b.Write(digits[i:]); err != nil {
			return err
		}
		*/
		x = sieve.Next()
	}

	//b.Write(float64ToByte(largestPrime))
	//if _, err = b.Write(digits[i:]); err != nil {
	if _, err = b.Write([]byte(largestPrime)); err != nil {
		return err
	}
	return b.Flush()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	var err error
	if len(args) == 1 {
		high, err = strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			usage()
		}
	} else if len(args) == 2 {
		low, err = strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			usage()
		}
		high, err = strconv.ParseUint(args[1], 10, 64)
		if err != nil {
			usage()
		}
	} else if len(args) > 2 {
		usage()
	}

	err = Write(os.Stdout, low, high)
	if err != nil {
		log.Fatalf("Error writing primes: %s\n", err)
	}
}
