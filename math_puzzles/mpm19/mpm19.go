package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	//file, err := os.Open("./file.txt")
	file, err := os.Open("./primes.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	primeCount := 0
	totalSum := 0.0
	for scanner.Scan() {
		//fmt.Println(primeCount, " ", scanner.Text())

		prime, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatal(err)
		}
		totalSum = totalSum+(prime*prime)
		//fmt.Println(totalSum)
		primeCount++
		//fmt.Println(primeCount, "\t", scanner.Text(), "\t", totalSum, "\t", totalSum/float64(primeCount))
		fSum := totalSum/float64(primeCount)
		if fSum == float64(int64(fSum)) {
			fmt.Println(primeCount, "\t", scanner.Text(), "\t", totalSum, "\t", fSum)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
