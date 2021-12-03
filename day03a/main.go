package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("data.txt")

	if err != nil {
		log.Fatalf("failed to open data.txt")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	file.Close()

	countOne := make([]int, 12)
	countZero := make([]int, 12)
	for _, line := range text {
		splits := strings.Split(line, "")
		for i := 0; i <= 11; i++ {
			if splits[i] == "0" {
				countZero[i]++
			} else {
				countOne[i]++
			}
		}
	}

	gammaBin := ""
	epsBin := ""
	for i := 0; i <= 11; i++ {
		if countZero[i] > countOne[i] {
			gammaBin += "0"
			epsBin += "1"
		} else {
			gammaBin += "1"
			epsBin += "0"
		}
	}

	gammaNum, err := strconv.ParseInt(gammaBin, 2, 64)
	if err != nil {
		fmt.Printf("gamma failed: %s\n", err.Error())
	}
	epsNum, err := strconv.ParseInt(epsBin, 2, 64)
	if err != nil {
		fmt.Printf("epsilon failed: %s\n", err.Error())
	}
	fmt.Printf("Gamma:   bin %s ... int %d\n", gammaBin, gammaNum)
	fmt.Printf("Epsilon: bin %s ... int %d\n", epsBin, epsNum)
	fmt.Printf("Result: %d\n", gammaNum*epsNum)
}
