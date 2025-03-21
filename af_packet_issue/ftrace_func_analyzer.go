package main

import (
	"bufio"
	"fmt"
  "log"
  "os"
	"regexp"
	"strconv"
//	"strings"
)

func main() {
  /* Testing data
	logData := `   1)               |                        packet_rcv() {
   1)   0.276 us    |                          consume_skb();
   1)   0.723 us    |                        }
   1)               |                        packet_rcv() {
   1)   0.221 us    |                          consume_skb();
   1)   0.627 us    |                        }
   1)               |                        vlan_do_receive() {
   1)   0.218 us    |                          skb_push();
   1)   0.246 us    |                          skb_push();
   1)   0.198 us    |                          skb_pull();
   1)   1.651 us    |                        }`
  */

	// Regex to match function execution time lines including function blocks
	rgx := regexp.MustCompile(`\s+(\d+)\)\s+([0-9\.]+) us\s+\|\s+(\w+)\(\);`)
	openingBraceRgx := regexp.MustCompile(`\s+\|\s+(\w+)\(\)\s+\{`)
	closingBraceRgx := regexp.MustCompile(`\s+(\d+)\)\s+([0-9\.]+) us\s+\|\s+}`)
	timeMap := make(map[string]float64)
	var functionStack []string

	//scanner := bufio.NewScanner(strings.NewReader(logData))

  // ftrace data file must bhave the HEADERS removed
  file, err := os.Open("./edited-ftrace-cpu1")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
  scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// Match function execution times
		matches := rgx.FindStringSubmatch(line)
		if matches != nil {
			timeVal, _ := strconv.ParseFloat(matches[2], 64)
			funcName := matches[3]
			timeMap[funcName] += timeVal
			continue
		}

		// Match function opening braces to track function names
		matches = openingBraceRgx.FindStringSubmatch(line)
		if matches != nil {
			funcName := matches[1]
			functionStack = append(functionStack, funcName) // Push function to stack
			continue
		}

		// Match closing brace for functions
		matches = closingBraceRgx.FindStringSubmatch(line)
		if matches != nil && len(functionStack) > 0 {
			timeVal, _ := strconv.ParseFloat(matches[2], 64)
			currentFunction := functionStack[len(functionStack)-1]
			timeMap[currentFunction] += timeVal
			functionStack = functionStack[:len(functionStack)-1] // Pop function from stack
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading input:", err)
		return
	}

	// Print total time spent in each function type
	for funcName, totalTime := range timeMap {
		fmt.Printf("%s: %.3f us\n", funcName, totalTime)
	}

  // Print total time for all function types
  totalFunctionTime := 0.0
  for _, funcTime := range timeMap {
    totalFunctionTime += funcTime
  }
  fmt.Printf("Total Func Time: %.3f us\n", totalFunctionTime)
}
